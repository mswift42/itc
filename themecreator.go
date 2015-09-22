package itc

import (
	"github.com/mswift42/go-colorful"
)

// Themeface represents the name of a color theme variable,
// for example keyword, variable, background, ... .
type Themeface string

var faces = []Themeface{"mainbg", "mainfg", "keyword",
	"comment", "type", "variable", "functionname",
	"warning", "doc", "string"}

// addColors takes a map of Themefaces and adds darker
// and lighter variants of the fg, bg and keyword faces.
func addColors(colors map[Themeface]string) map[Themeface]string {
	fg := colors["mainfg"]
	bg := colors["mainbg"]
	keyword := colors["keyword"]
	bg2 := ""
	bg3 := ""
	bg4 := ""
	fg2 := ""
	fg3 := ""
	fg4 := ""
	fgcol, _ := colorful.Hex(fg)
	bgcol, _ := colorful.Hex(bg)
	keycol, _ := colorful.Hex(keyword)
	if hasDarkBg(&bgcol) {
		fg2 = darken(fgcol, 0.08)
		fg3 = darken(fgcol, 0.16)
		fg4 = darken(fgcol, 0.24)
		bg2 = lighten(bgcol, 0.08)
		bg3 = lighten(bgcol, 0.16)
		bg4 = lighten(bgcol, 0.24)
	} else {
		fg2 = lighten(fgcol, 0.08)
		fg3 = lighten(fgcol, 0.16)
		fg4 = lighten(fgcol, 0.24)
		bg2 = darken(bgcol, 0.08)
		bg3 = darken(bgcol, 0.16)
		bg4 = darken(bgcol, 0.24)
	}
	key2 := lighten(keycol, 0.11)
	key3 := darken(keycol, 0.11)
	colors["fg2"] = fg2
	colors["fg3"] = fg3
	colors["fg4"] = fg4
	colors["bg2"] = bg2
	colors["bg3"] = bg3
	colors["bg4"] = bg4
	colors["key2"] = key2
	colors["key3"] = key3
	return colors
}

// darken - return a darker variant of a supplied color.
func darken(col colorful.Color, factor float64) string {
	black, _ := colorful.Hex("#000000")
	return col.BlendLab(black, factor).Hex()
}

// lighten - return a lighter variant of a supplied color.
func lighten(col colorful.Color, factor float64) string {
	white, _ := colorful.Hex("#ffffff")
	return col.BlendLab(white, factor).Hex()
}

// hasDarkBg - Check if supplied color is dark.
func hasDarkBg(c *colorful.Color) bool {
	l, _, _ := c.Lab()
	return l < 0.5
}
