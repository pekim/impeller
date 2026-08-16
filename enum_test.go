package impeller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnum(t *testing.T) {
	// verify that enums have been generated, and that a value is correct
	assert.Equal(t, TextDecorationType(4), TextDecorationTypeLineThrough)
}
