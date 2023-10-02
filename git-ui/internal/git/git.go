package git

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func getDiff() string {
	cmd := exec.Command("git", "diff", "../../testfile.txt")

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal("Failed to get git diff")
	}

	return strings.TrimSpace(out.String())
}
