package itc

// Themeface represents the name of a color theme variable,
// for example keyword, variable, background, ... .
type Themeface string

var faces = []Themeface{"mainbg", "mainfg", "keyword",
	"comment", "type", "variable", "functionname",
	"warning", "doc", "string"}
