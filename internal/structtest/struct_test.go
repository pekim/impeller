package structtest

import (
	"fmt"
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructs(t *testing.T) {
	for _, struct_ := range structs {
		cValue := reflect.ValueOf(struct_.c)
		goValue := reflect.ValueOf(struct_.go_)

		assert.Equal(t, cValue.Type().Size(), goValue.Type().Size(),
			fmt.Sprintf("%q size", struct_.name))

		cFields := reflect.VisibleFields(cValue.Type())
		cFields = removeAnonymousFields(cFields)
		goFields := reflect.VisibleFields(goValue.Type())
		goFields = removeAnonymousFields(goFields)
		assert.Equal(t, len(cFields), len(goFields),
			fmt.Sprintf("%s fields count", struct_.name))

		for i := range cFields {
			cField := cFields[i]
			goField := goFields[i]

			assert.Equal(t, cField.Offset, goField.Offset,
				fmt.Sprintf("%s.%s offset", struct_.name, cField.Name))
			assert.Equal(t, cField.Type.Size(), goField.Type.Size(),
				fmt.Sprintf("%s.%s size", struct_.name, cField.Name))
		}
	}
}

func removeAnonymousFields(fields []reflect.StructField) []reflect.StructField {
	return slices.DeleteFunc(fields, func(field reflect.StructField) bool {
		return field.Name == "_"
	})
}
