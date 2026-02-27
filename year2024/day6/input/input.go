package input

import "github.com/Drozd0f/reader"

func ReadExample() []string {
	return reader.ReadAllFile(Input, "example.txt")
}

func ReadInput() []string {
	return reader.ReadAllFile(Input, "input.txt")
}
