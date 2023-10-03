package git

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
	"unicode/utf8"
)

type Diff struct {
	Diff1 []DiffLine
	Diff2 []DiffLine
}

type DiffLine struct {
	Content string
}

func GetDiff(diffString string) Diff {
	lines := strings.Split(diffString, "\n")

	diff := Diff{Diff1: make([]DiffLine, 0), Diff2: make([]DiffLine, 0)}

	start := false
	removals := 0
	additions := 0
	diffLineBlank := DiffLine{Content: ""}
	for _, line := range lines {
		if strings.HasPrefix(line, "@@") && strings.HasSuffix(line, "@@") {
			start = true
			continue
		}

		if !start {
			continue
		}

		letter, lineString := trimFirstRune(line)
		diffLine := DiffLine{Content: lineString}

		if letter == '-' {
			removals++
			diff.Diff1 = append(diff.Diff1, diffLine)
		} else if letter == '+' {
			diff.Diff2 = append(diff.Diff2, diffLine)
			if removals > 0 {
				removals--
			}

			if removals == 0 {
				additions++
			}
		} else {
			for i := 0; i < removals; i++ {
				diff.Diff2 = append(diff.Diff2, diffLineBlank)
			}
			removals = 0

			for i := 0; i < additions; i++ {
				diff.Diff1 = append(diff.Diff1, diffLineBlank)
			}
			additions = 0

			diff.Diff1 = append(diff.Diff1, diffLine)
			diff.Diff2 = append(diff.Diff2, diffLine)
		}
	}

	for i := 0; i < removals; i++ {
		diff.Diff2 = append(diff.Diff2, diffLineBlank)
	}

	for i := 0; i < additions; i++ {
		diff.Diff1 = append(diff.Diff1, diffLineBlank)
	}

	return diff
}

func trimFirstRune(s string) (rune, string) {
	r, i := utf8.DecodeRuneInString(s)
	return r, s[i:]
}

func GetRawDiff() string {
	cmd := exec.Command("git", "diff", "--no-prefix", "-U1000", "testfile.txt")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal("Failed to get git diff")
	}

	return out.String()
}
