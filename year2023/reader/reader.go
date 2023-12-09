package reader

import (
	"bufio"
	"embed"
)

const (
	defaultPuzzleInput = "input.txt"
)

func ReadAllFile(fs embed.FS, path string) []string {
	if path == "" {
		path = defaultPuzzleInput
	}

	file, err := fs.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}
