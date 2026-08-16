package generate

import (
	"fmt"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

func (gen gen) commentText(cursor clang.Cursor) string {
	lineComment := gen.lineComment(cursor)
	if lineComment != "" {
		return lineComment
	}

	text := cursor.RawCommentText()
	text = strings.TrimPrefix(text, "/**")
	text = strings.TrimSuffix(text, "*/")

	var builder strings.Builder
	var codeBlock bool

	// for each line
	//		- trim leading "///"
	//		- trim leading and trailing white space
	lines := strings.Split(text, "\n")
	if strings.Contains(text, "Estimated number of rows returned") {
		fmt.Println(text)
		fmt.Println(lines)
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "///")
		line = strings.TrimSpace(line)

		// support code blocks
		if strings.HasPrefix(line, "```") {
			codeBlock = !codeBlock
			continue
		}
		if codeBlock {
			line = "  " + line
		}

		builder.WriteString(line)
		builder.WriteRune('\n')
	}

	return strings.TrimSpace(builder.String())
}

func (gen gen) lineComment(cursor clang.Cursor) string {
	range_ := cursor.Extent()
	_, endLine, endCol, _ := range_.End().FileLocation()
	line := gen.lines[endLine-1]

	comment := line[endCol-1:]
	comment = strings.TrimSpace(comment)
	comment = strings.TrimPrefix(comment, ",")
	comment = strings.TrimPrefix(comment, ";")
	comment = strings.TrimSpace(comment)

	if after, ok := strings.CutPrefix(comment, "//"); ok {
		comment = after
		comment = strings.TrimSpace(comment)
		return comment
	}

	if after, ok := strings.CutPrefix(comment, "/*"); ok {
		comment = after
		comment = strings.TrimSuffix(comment, "*/")
		comment = strings.TrimSpace(comment)
		return comment
	}

	return ""
}
