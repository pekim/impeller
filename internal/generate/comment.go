package generate

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

var commentBacktickReferenceRegexp = regexp.MustCompile(".*`(.*?)`.*")

type commentParams map[string]clang.CommentParamPassDirection

type comment struct {
	gen                       *gen
	cursor                    clang.Cursor
	rawText                   string
	text                      string
	params                    commentParams
	inMarkdownFencedCodeBlock bool
	inParam                   bool
}

func (gen *gen) newComment(cursor clang.Cursor) comment {
	c := comment{
		gen:     gen,
		cursor:  cursor,
		rawText: cursor.RawCommentText(),
		params:  make(commentParams),
	}

	var builder strings.Builder
	c.processComment(&builder, c.cursor.ParsedComment())
	c.text = builder.String()

	return c
}

func (c *comment) processComment(builder *strings.Builder, comment clang.Comment) {
	switch comment.Kind() {

	case clang.Comment_BlockCommand:
		switch comment.BlockCommandComment_getCommandName() {
		case "brief":
			c.processChildren(builder, comment)
			builder.WriteRune('\n')

		case "deprecated":
			builder.WriteString("\nDeprecated: ")
			c.processChildren(builder, comment)

		case "note":
			builder.WriteString("\nNote: ")
			c.processChildren(builder, comment)

		case "return":
			builder.WriteString("\nReturn: ")
			c.processChildren(builder, comment)

		case "see":
			reference := comment.Child(0).Child(0).TextComment_getText()
			reference = strings.TrimSpace(reference)

			// remove surrounding backticks
			if strings.HasPrefix(reference, "`") && strings.HasSuffix(reference, "`") {
				reference = reference[1 : len(reference)-1]
			}

			if entity, found := c.gen.resolveEntityReference(reference); found {
				fmt.Fprintf(builder, "\n\nsee [%s]", entity)
			} else {
				fmt.Fprintf(builder, "\n\nsee %s", reference)
			}

		case "since":
			builder.WriteString("Since: ")
			c.processChildren(builder, comment)

		case "warning":
			builder.WriteString("\nWarning: ")
			c.processChildren(builder, comment)

		default:
			fmt.Println(c.rawText)
			fatalf("unhandled block command : %s\n", comment.BlockCommandComment_getCommandName())
		}

	case clang.Comment_InlineCommand:
		switch comment.InlineCommandComment_getCommandName() {

		case "n": // newline
			builder.WriteRune('\n')
			builder.WriteString(comment.InlineCommandComment_getArgText(0))

		default:
			fmt.Println(c.rawText)
			fatalf("unhandled inline command : %s\n", comment.InlineCommandComment_getCommandName())
		}

	case clang.Comment_ParamCommand:
		name := comment.ParamCommandComment_getParamName()
		direction := comment.ParamCommandComment_getDirection()

		c.params[name] = direction
		fmt.Fprintf(builder, "  - %s - ", name)

		c.inParam = true
		c.processChildren(builder, comment)
		c.inParam = false

	case clang.Comment_FullComment:
		c.processChildren(builder, comment)

	case clang.Comment_Null:
		c.processChildren(builder, comment)

	case clang.Comment_Paragraph:
		c.processChildren(builder, comment)

		if c.inParam {
			builder.WriteString("\n")
		} else {
			builder.WriteString("\n\n")
		}

	case clang.Comment_Text:
		commentText := comment.TextComment_getText()
		commentText = strings.TrimSpace(commentText)

		if strings.Contains(commentText, "```") {
			c.inMarkdownFencedCodeBlock = !c.inMarkdownFencedCodeBlock
		} else {
			if c.inMarkdownFencedCodeBlock {
				builder.WriteString("  ")
				builder.WriteString(commentText)
				builder.WriteString("\n")
			} else {
				match := commentBacktickReferenceRegexp.FindStringSubmatch(commentText)
				if match != nil {
					reference := match[1]
					if entity, found := c.gen.resolveEntityReference(reference); found {
						commentText = strings.Replace(commentText, "`"+reference+"`", "["+entity+"]", 1)
					} else {
						commentText = strings.Replace(commentText, "`"+reference+"`", reference, 1)
					}
				}

				builder.WriteString(commentText)
			}
		}

		c.processChildren(builder, comment)

	default:
		fmt.Println(c.rawText)
		fatalf("unhandled comment type : %s\n", comment.Kind().Spelling())
	}
}

func (c *comment) processChildren(builder *strings.Builder, comment clang.Comment) {
	previousWasText := false

	for i := range comment.NumChildren() {
		child := comment.Child(i)
		if previousWasText && child.Kind() == clang.Comment_Text {
			builder.WriteRune(' ')
		}
		c.processComment(builder, child)
		previousWasText = child.Kind() == clang.Comment_Text
	}
}

func (c comment) Dump() {
	c.dumpComment(c.cursor.ParsedComment(), "")
}

func (c comment) dumpComment(comment clang.Comment, indent string) {
	fmt.Println(indent, comment.Kind().Spelling(), "=>", comment.TextComment_getText())

	switch comment.Kind() {
	case clang.Comment_BlockCommand:
		fmt.Println(indent, "  Command : Block :", comment.BlockCommandComment_getCommandName())
	case clang.Comment_InlineCommand:
		fmt.Println(indent, "  Command : Inline :", comment.InlineCommandComment_getCommandName())
	case clang.Comment_ParamCommand:
		fmt.Println(indent, "  Command : Param :", comment.ParamCommandComment_getDirection(), comment.ParamCommandComment_getParamName())

	}

	for i := range comment.NumChildren() {
		c.dumpComment(comment.Child(i), indent+"  ")
	}
}
