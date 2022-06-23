package main

import (
	"fmt"
	"os"
	"strings"
)

var resultFileName = "extensions.generated.go"

var header = `// Code generated by generate/extensions.go DO NOT EDIT.
package iterator

import (
	"iterator/extended"
	"iterator/util"
)
`

var orderedTypes = []string{
	"int", "int8", "int16", "int32", "int64",
	"uint", "uint8", "uint16", "uint32", "uint64",
	"float32", "float64",
	"string",
}

var functionsToGenerate = []string{
	`
func (i Iterator[T]) SortedBy%[1]s(keyFunc func(T) %[2]s) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIteratorAsc(i.iter, keyFunc)}
}`,
	`
func (i Iterator[T]) SortedBy%[1]sDescending(keyFunc func(T) %[2]s) Iterator[T] {
	return Iterator[T]{iter: extended.SortingIteratorDesc(i.iter, keyFunc)}
}`,
	`
func (i Iterator[T]) DistinctBy%[1]s(keyFunc func(T) %[2]s) Iterator[T] {
	return Iterator[T]{iter: extended.DistinctingIterator(i.iter, keyFunc)}
}`,
	`
func (i Iterator[T]) MaxBy%[1]s(keyFunc func(T) %[2]s) (T, bool) {
	return util.MaxBy(i.iter, keyFunc)
}
`,
	`
func (i Iterator[T]) MaxBy%[1]sOrZeroValue(keyFunc func(T) %[2]s) T {
	return util.MaxByOrZeroValue(i.iter, keyFunc)
}
`,
	`
func (i Iterator[T]) MaxBy%[1]sOrDefault(keyFunc func(T) %[2]s, def T) T {
	return util.MaxByOrDefault(i.iter, keyFunc, def)
}
`,
	`
func (i Iterator[T]) MinBy%[1]s(keyFunc func(T) %[2]s) (T, bool) {
	return util.MinBy(i.iter, keyFunc)
}
`,
	`
func (i Iterator[T]) MinBy%[1]sOrZeroValue(keyFunc func(T) %[2]s) T {
	return util.MinByOrZeroValue(i.iter, keyFunc)
}
`,
	`
func (i Iterator[T]) MinBy%[1]sOrDefault(keyFunc func(T) %[2]s, def T) T {
	return util.MinByOrDefault(i.iter, keyFunc, def)
}
`,
}

func main() {
	file, err := os.Create(resultFileName)
	if err != nil {
		fmt.Printf("Error creating file %s: %v", resultFileName, err)
		return
	}
	defer func() { _ = file.Close() }()
	_, err = file.WriteString(header)
	if err != nil {
		fmt.Printf("Error writing: %v", err)
		return
	}
	for _, template := range functionsToGenerate {
		for _, typeName := range orderedTypes {
			typeNameCap := strings.Title(typeName)
			_, err = fmt.Fprintf(file, template+"\n", typeNameCap, typeName)
		}
	}
	fmt.Println("OK")
}
