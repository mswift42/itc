package itc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddColors(t *testing.T) {
	assert := assert.New(t)
	fm := make(map[Themeface]string)
	somecolors := []string{"#000000", "#ffffff", "#aabbcc",
		"#aeaeae", "#122458", "#abcdef", "#fedcba",
		"#ff0000", "#ff00ff", "#00ffff"}
	for i := range somecolors {
		fm[faces[i]] = somecolors[i]
	}
	addColors(fm)
	assert.NotEmpty(fm)
	assert.Equal(fm["mainbg"], "#000000")

}
