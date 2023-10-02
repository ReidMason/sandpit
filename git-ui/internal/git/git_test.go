package git

import (
	"strings"
	"testing"
)

func TestGetDiff(t *testing.T) {
	result := getDiff()

	expected := `diff --git a/git-ui/testfile.txt b/git-ui/testfile.txt
index 35b5809..4492ac6 100644
--- a/git-ui/testfile.txt
+++ b/git-ui/testfile.txt
@@ -1,4 +1,5 @@
 This is a test file
-
 These lines are committed now
-I have added some more content
+I have added this is a change more content
+
+This is a new thing`

	resultLines := strings.Split(result, "\n")
	for i, line := range strings.Split(expected, "\n") {
		resLine := resultLines[i]

		if line != resLine {
			t.Fatalf("Line matche failed. Expected: '%s' Got: '%s'", line, resLine)
		}
	}

	if len(result) != len(expected) {
		t.Fatalf("Length difference. Expected: %d Got: %d", len(expected), len(result))
	}
}
