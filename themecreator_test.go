package main

import (
	"testing"

	"github.com/mswift42/go-colorful"
	"github.com/stretchr/testify/assert"
)

var lightendedColors = []struct {
	old     string
	lighter string
	factor  float64
}{
	{
		old:     "#000000",
		lighter: "#303030",
		factor:  0.20,
	},
	{
		old:     "#202020",
		lighter: "#464646",
		factor:  0.20,
	},
	{
		old:     "#34b156",
		lighter: "#67c177",
		factor:  0.2,
	},
	{
		old:     "#23557a",
		lighter: "#547493",
		factor:  0.2,
	},
	{
		old:     "#ffffff",
		lighter: "#ffffff",
		factor:  0.1,
	},
	{
		old:     "#0e2f17",
		lighter: "#3a543f",
		factor:  0.2,
	},
	{
		old:     "#721d2e",
		lighter: "#a06066",
		factor:  0.3,
	},
	{
		old:     "#8a2436",
		lighter: "#a5525a",
		factor:  0.2,
	},
}
var darkenedColors = []struct {
	old    string
	darker string
	factor float64
}{
	{
		old:    "#ffffff",
		darker: "#e2e2e2",
		factor: 0.1,
	},
	{
		old:    "#e1e1e1",
		darker: "#d4d4d4",
		factor: 0.05,
	},
	{
		old:    "#808080",
		darker: "#727272",
		factor: 0.1,
	},
	{
		old:    "#ff40ff",
		darker: "#f13ef1",
		factor: 0.05,
	},
	{
		old:    "#b72fc1",
		darker: "#ad2eb6",
		factor: 0.05,
	},
	{
		old:    "#000000",
		darker: "#000000",
		factor: 0.05,
	},
	{
		old:    "#a5525a",
		darker: "#824248",
		factor: 0.2,
	},
	{
		old:    "#dab071",
		darker: "#93784f",
		factor: 0.3,
	},
}

var darkBg = []struct {
	bg  string
	yes bool
}{
	{
		bg:  "#000000",
		yes: true,
	},
	{
		bg:  "#ffffff",
		yes: false,
	},
	{
		bg:  "#002b36",
		yes: true,
	},
	{
		bg:  "#586e75",
		yes: true,
	},
	{
		bg:  "#fdf6e3",
		yes: false,
	},
	{
		bg:  "#d2d2d2",
		yes: false,
	},
	{
		bg:  "#daa1e6",
		yes: false,
	},
	{
		bg:  "#932ad7",
		yes: true,
	},
	{
		bg:  "#d7cf47",
		yes: false,
	},
	{
		bg:  "#87e37e",
		yes: false,
	},
	{
		bg:  "#f0c1bc",
		yes: false,
	},
	{
		bg:  "#dc6c60",
		yes: false,
	},
	{
		bg:  "#39c52b",
		yes: false,
	},
}

func TestLighten(t *testing.T) {
	assert := assert.New(t)
	for _, i := range lightendedColors {
		c, _ := colorful.Hex(i.old)
		actual := lighten(c, i.factor)
		assert.Equal(actual, i.lighter)
	}
}
func TestDarken(t *testing.T) {
	assert := assert.New(t)
	for _, i := range darkenedColors {
		c, _ := colorful.Hex(i.old)
		actual := darken(c, i.factor)
		assert.Equal(actual, i.darker)
	}
}

func TestHasDarkBg(t *testing.T) {
	assert := assert.New(t)
	for _, i := range darkBg {
		col, err := colorful.Hex(i.bg)
		if err != nil {
			t.Fatal(err)
		}
		actual := hasDarkBg(&col)
		assert.Equal(actual, i.yes)
	}
}

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
	assert.Equal(fm["keyword"], "#aabbcc")
	assert.NotNil(fm["fg2"])
	assert.Equal(fm["fg2"], "#e8e8e8")
	assert.Equal(fm["fg3"], "#d1d1d1")
	assert.Equal(fm["fg4"], "#bbbbbb")
	assert.Equal(fm["bg2"], "#181818")
	assert.Equal(fm["bg3"], "#282828")
	assert.Equal(fm["bg4"], "#393939")
	assert.Equal(fm["key2"], "#b3c2d2")
	assert.Equal(fm["key3"], "#96a4b3")

}
