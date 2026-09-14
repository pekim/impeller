package generate

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/go-clang/clang-v15/clang"
)

type paramDirection byte

const (
	in paramDirection = iota
	out
)

type commentParam struct {
	name        string
	description string
	direction   paramDirection
}

type commentLink struct {
	text string
	url  string
}

type comment struct {
	gen        *gen
	entityName string
	lines      []string
	params     map[string]commentParam
	links      []commentLink
	sinceMajor int
	sinceMinor int
}

type commentParams map[string]commentParam

func (gen *gen) newComment(cursor clang.Cursor) comment {
	comment := comment{
		gen:        gen,
		entityName: cursor.Spelling(),
		lines:      gen.commentLines(cursor),
	}
	comment.cleanLines()
	comment.setParams()
	comment.setSince()
	comment.formatHeadings()
	comment.formatCodeBlocks()

	return comment
}

func (comment *comment) cleanLines() {
	for i, line := range comment.lines {
		line = strings.TrimSpace(line)
		line = strings.TrimSuffix(line, "*/")
		line = strings.TrimPrefix(line, "*")
		line = strings.TrimPrefix(line, "/*!")
		line = strings.TrimPrefix(line, "/*")
		line = strings.TrimSpace(line)

		line = strings.TrimPrefix(line, "@note")
		line = strings.TrimPrefix(line, "@remarks")
		line = strings.TrimPrefix(line, "@remark")
		line = strings.TrimPrefix(line, "@warning")
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "@brief") ||
			strings.HasPrefix(line, "@callback_signature") ||
			strings.HasPrefix(line, "@defgroup") ||
			strings.HasPrefix(line, "@glfw3") ||
			strings.HasPrefix(line, "@ingroup") ||
			line == "@par" ||
			strings.HasPrefix(line, "@sa") || // TODO convert resolvable references to doc links
			strings.HasPrefix(line, "@{") ||
			strings.HasPrefix(line, "@}") {
			line = ""
		}

		comment.lines[i] = line
	}
}

func (comment *comment) setParams() {
	comment.params = make(commentParams)

	var param commentParam
	var withinDescription bool
	for i, line := range comment.lines {
		if withinDescription {
			if line == "" || strings.HasPrefix(line, "@") {
				// description was terminated by a new paragraph or another command
				comment.params[param.name] = param
				withinDescription = false
			} else {
				// description continues
				param.description += "\n" + line
				comment.lines[i] = "  " + line
			}
		}

		if strings.HasPrefix(line, "@param") {
			parts := strings.SplitN(line, " ", 3)
			param = commentParam{
				name: parts[1],
			}

			switch parts[0] {
			case "@param[in]":
				param.direction = in
			case "@param[out]":
				param.direction = out
			}

			if len(parts) > 2 {
				param.description = parts[2]
				withinDescription = true
			} else {
				comment.params[param.name] = param
			}

			var line string
			if len(comment.params) == 0 {
				// this is the first param, so create a header
				line = "\n# params\n"
			}
			outText := ""
			if param.direction == out {
				outText = " (out)"
			}
			line += fmt.Sprintf("  - %s%s - %s", param.name, outText, param.description)
			comment.lines[i] = line
		}
	}

	if withinDescription {
		comment.params[param.name] = param
	}
}

func (comment *comment) setSince() {
	re := regexp.MustCompile(`@since .* (\d)\.(\d)`)
	for _, line := range comment.lines {
		groups := re.FindStringSubmatch(line)
		if len(groups) != 3 {
			continue
		}

		major, err := strconv.Atoi(groups[1])
		if err != nil {
			fatalOnError(err)
		}
		minor, err := strconv.Atoi(groups[2])
		if err != nil {
			fatalOnError(err)
		}
		comment.sinceMajor = major
		comment.sinceMinor = minor
		break
	}
}

func (comment *comment) formatHeadings() {
	commands := []string{
		"@analysis", // non-standard, not documented at https://www.doxygen.nl/manual/commands.html
		"@deprecated",
		"@errors",
		"@macos",
		"@par",
		"@pointer_lifetime",
		"@reentrancy",
		"@return",
		"@since",
		"@thread_safety",
		"@wayland",
		"@win32",
		"@x11",
	}

	for i, line := range comment.lines {
		parts := strings.SplitN(line, " ", 2)
		if len(parts) == 2 && slices.Contains(commands, parts[0]) {
			title := parts[0][1:]
			body := parts[1]
			title = strings.ReplaceAll(title, "_", " ")

			if title == "par" {
				comment.lines[i] = "# " + body + "\n"
			} else {
				comment.lines[i] = "\n# " + title + "\n\n" + body
			}
		}
	}
}

func (comment *comment) formatCodeBlocks() {
	inBlock := false
	for i, line := range comment.lines {
		// line = strings.TrimSpace(line)
		if line == "```" {
			inBlock = !inBlock
		}
		if inBlock {
			comment.lines[i] = "    " + line
		} else {
			comment.lines[i] = line
		}
	}
}

var linkRegexp = regexp.MustCompile("`([^\\s]+)`")

func (comment *comment) resolveLinks() {
	for l, line := range comment.lines {
		parts := linkRegexp.FindStringSubmatch(line)
		if parts == nil {
			continue
		}

		if reference, ok := comment.gen.resolveEntityReference(parts[1]); ok {
			line := strings.Replace(line, parts[0], "["+reference+"]", 1)
			comment.lines[l] = line
		}
	}
}

var seeRegexp = regexp.MustCompile(`@see\s+([^\s]+)`)

func (comment *comment) resolveSees() {
	for l, line := range comment.lines {
		parts := seeRegexp.FindStringSubmatch(line)
		if parts == nil {
			continue
		}

		ref := parts[1]
		ref = strings.TrimSpace(ref)
		ref = strings.TrimPrefix(ref, "`")
		ref = strings.TrimSuffix(ref, "`")

		if strings.HasPrefix(ref, "http") {
			comment.lines[l] = fmt.Sprintf("See %s", ref)
		} else {
			if reference, ok := comment.gen.resolveEntityReference(ref); ok {
				comment.lines[l] = fmt.Sprintf("see [%s]", reference)
			}
		}
	}
}

func (comment comment) text() string {
	comment.resolveSees()
	comment.resolveLinks()

	if len(comment.links) > 0 {
		comment.lines = append(comment.lines, "")
		for _, link := range comment.links {
			comment.lines = append(comment.lines, fmt.Sprintf("[%s]: %s", link.text, link.url))
		}
	}

	return strings.Join(comment.lines, "\n")
}

func (gen gen) commentLines(cursor clang.Cursor) []string {
	_, cursorLine, _, _ := cursor.Location().FileLocation()

	var lines []string
	for l := int(cursorLine) - 2; l >= 0; l-- {
		line := strings.TrimSpace(gen.lines[l])

		if strings.HasPrefix(line, "IMPELLER_EXPORT") {
			continue
		}

		if !strings.HasPrefix(line, "//") || strings.HasPrefix(line, "//--") {
			break
		}

		line = strings.TrimPrefix(line, "///")
		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}
	slices.Reverse(lines)
	lines = append(lines, gen.lineComment(cursor))

	return lines
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
