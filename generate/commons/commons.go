package commons

import "strings"

var OrderedTypes = []string{
	"int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64",
	"float32", "float64",
	"string",
}

func Capitalize(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}
