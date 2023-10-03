package git

import (
	"testing"
)

// func TestGetRawDiff(t *testing.T) {
// 	result := GetRawDiff()
//
// 	expected := `diff --git git-ui/testfile.txt git-ui/testfile.txt
// index 35b5809..4492ac6 100644
// --- git-ui/testfile.txt
// +++ git-ui/testfile.txt
// @@ -1,4 +1,5 @@
//  This is a test file
// -
//  These lines are committed now
// -I have added some more content
// +I have added this is a change more content
// +
// +This is a new thing`
//
// 	resultLines := strings.Split(result, "\n")
// 	for i, line := range strings.Split(expected, "\n") {
// 		resLine := resultLines[i]
//
// 		if line != resLine {
// 			t.Fatalf("Line match failed. Expected: '%s' Got: '%s'", line, resLine)
// 		}
// 	}
//
// 	if len(result) != len(expected) {
// 		t.Fatalf("Length difference. Expected: %d Got: %d", len(expected), len(result))
// 	}
// }

func TestGetDiff(t *testing.T) {
	rawDiff := `diff --git a/git-ui/testfile.txt b/git-ui/testfile.txt
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

	result := GetDiff(rawDiff)

	expected := Diff{
		Diff1: []DiffLine{
			{Content: "This is a test file"},
			{Content: ""},
			{Content: "These lines are committed now"},
			{Content: "I have added some more content"},
			{Content: ""},
			{Content: ""},
		},
		Diff2: []DiffLine{
			{Content: "This is a test file"},
			{Content: ""},
			{Content: "These lines are committed now"},
			{Content: "I have added this is a change more content"},
			{Content: ""},
			{Content: "This is a new thing"},
		},
	}

	t.Log("----- Start diff1")
	for i, expectedDiffLine := range expected.Diff1 {
		resultDiffLine := result.Diff1[i]
		t.Log(resultDiffLine.Content)

		if expectedDiffLine.Content != resultDiffLine.Content {
			t.Fatalf("Diff1 failed line %d.\nExpected: '%s'\n     Got: '%s'", i+1, expectedDiffLine.Content, resultDiffLine.Content)
		}
	}

	t.Log("----- Start diff2")
	for i, expectedDiffLine := range expected.Diff2 {
		resultDiffLine := result.Diff2[i]
		t.Log(resultDiffLine.Content)

		if expectedDiffLine.Content != resultDiffLine.Content {
			t.Fatalf("Diff2 failed line %d.\nExpected: '%s'\n     Got: '%s'", i+1, expectedDiffLine.Content, resultDiffLine.Content)
		}
	}

	// if !reflect.DeepEqual(result, expected) {
	// 	t.Fatal("Expected diff length doesn't match")
	// }
}
