package generate

import "strings"

func goName(cName string) string {
	name := cName
	name = strings.TrimPrefix(name, "kImpeller")
	name = strings.TrimPrefix(name, "Impeller")
	name = strings.TrimSuffix(name, "_")
	return name
}
