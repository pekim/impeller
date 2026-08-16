package generate

import (
	"fmt"
	"strconv"
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
	bulletList := false
	numberedList := false
	var itemNumber int

	// for each line
	//		- trim leading "**"
	// 		- trim leading and trailing white space
	//		- trim leading "^"
	// remove lines that start with
	//		- CAPI3REF:
	//		- KEYWORDS:
	//		- METHOD:
	// convert <ol>...</ol> in to numbered lists
	// convert <ul>...</ul> in to bullet lists
	lines := strings.Split(text, "\n")
	if strings.Contains(text, "Estimated number of rows returned") {
		fmt.Println(text)
		fmt.Println(lines)
	}
	for _, line := range lines {
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "/**")
		line = strings.TrimSuffix(line, "*/")
		line = strings.TrimPrefix(line, "**")
		line = strings.TrimSpace(line)
		line = strings.TrimPrefix(line, "^")
		line = strings.TrimSpace(line)

		// remove special lines that are of no interest
		if strings.HasPrefix(line, "CAPI3REF:") || strings.HasPrefix(line, "KEYWORDS:") || strings.HasPrefix(line, "METHOD:") {
			continue
		}

		// support ordered and unordered lists
		// https://go.dev/doc/comment#lists
		if after, ok := strings.CutPrefix(line, "<ol>"); ok {
			line = after
			numberedList = true
			itemNumber = 1
		}
		if after, ok := strings.CutPrefix(line, "</ol>"); ok {
			line = after
			numberedList = false
		}
		if after, ok := strings.CutPrefix(line, "<ul>"); ok {
			line = after
			bulletList = true
		}
		if after, ok := strings.CutPrefix(line, "</ul>"); ok {
			line = after
			bulletList = false
		}
		if after, ok := strings.CutPrefix(line, "<li>"); ok {
			line = strings.TrimSpace(after)
			if bulletList {
				line = "- " + line
			}
			if numberedList {
				line = strconv.Itoa(itemNumber) + ". " + line
				itemNumber++
			}
		}
		line = strings.Replace(line, "</li>", "", 1)
		if bulletList || numberedList {
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
