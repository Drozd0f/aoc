package day2

import (
	"regexp"
	"strings"

	"github.com/Drozd0f/tools"
)

var re = regexp.MustCompile(`(\d+)-(\d+)`)

func Parse(lines []string) [][2]int64 {
	seqs := strings.Split(lines[0], ",")

	parsed := make([][2]int64, len(seqs))
	for idx, seq := range seqs {
		s := re.FindStringSubmatch(seq)

		parsed[idx] = [2]int64{
			tools.MustInt64FromString(s[1]),
			tools.MustInt64FromString(s[2]),
		}
	}

	return parsed
}
