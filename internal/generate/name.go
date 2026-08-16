package generate

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var titleCaser = cases.Title(language.English, cases.NoLower)

func goName(cName string) string {
	name := cName
	name = strings.TrimPrefix(name, "kImpeller")
	name = strings.TrimPrefix(name, "Impeller")

	parts := strings.Split(name, "_")
	for i, part := range parts {
		parts[i] = titleCaser.String(part)
	}
	name = strings.Join(parts, "")

	return name
}
