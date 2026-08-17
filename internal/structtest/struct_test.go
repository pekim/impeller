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
		t.Run(struct_.name, func(t *testing.T) {
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

				t.Run(cField.Name, func(t *testing.T) {
					assert.Equal(t, cField.Offset, goField.Offset, "offset")
					assert.Equal(t, cField.Type.Size(), goField.Type.Size(), "size")
				})
			}
		})
	}
}

func removeAnonymousFields(fields []reflect.StructField) []reflect.StructField {
	return slices.DeleteFunc(fields, func(field reflect.StructField) bool {
		return field.Name == "_"
	})
}
