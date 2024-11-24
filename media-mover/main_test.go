package main

import (
	"testing"
	"time"
)

func TestParseFilenameDate(t *testing.T) {
	testCases := []struct {
		file     File
		expected time.Time
		errors   bool
	}{
		{File{Name: "20170429_104146.jpg"}, time.Date(2017, 4, 29, 0, 0, 0, 0, time.UTC), false},
		{File{Name: "20170512_025705.mp4"}, time.Date(2017, 5, 12, 0, 0, 0, 0, time.UTC), false},
		{File{Name: "Screenshot_20170512-App.mp4"}, time.Date(2017, 5, 12, 0, 0, 0, 0, time.UTC), false},
		{File{Name: "Screenshot_2016-03-18-22-26-29.png"}, time.Date(2016, 3, 18, 0, 0, 0, 0, time.UTC), false},
		{File{Name: "1560169653854.jpg"}, time.Date(2019, 6, 10, 0, 0, 0, 0, time.UTC), false},
		{File{Name: "2019-12-17_00-17-20_Snapchat-2032703187.jpg"}, time.Date(2019, 12, 17, 0, 0, 0, 0, time.UTC), false},
		{File{Name: "2017h0512_025705.mp4"}, time.Time{}, true},
	}

	t.Parallel()
	oneDay := 24 * time.Hour
	for _, tc := range testCases {
		date, err := parseFilenameDate(tc.file.Name)
		if tc.errors && err != nil {
			continue
		} else if tc.errors && err == nil {
			t.Errorf("Expected error parsing date from filename %s", tc.file.Name)
		}

		if err != nil {
			t.Errorf("Error parsing date from filename %s: %v", tc.file.Name, err)
		}

		if date.Truncate(oneDay).Sub(tc.expected.Truncate(oneDay)).Seconds() != 0 {
			t.Error(date.Truncate(oneDay).Sub(tc.expected.Truncate(oneDay)))
			t.Errorf("Expected date to be parsed from filename %s", tc.file.Name)
			t.Errorf("Expected: %v, got: %v", tc.expected, date)
		}
	}
}

func TestBuildDestinationPath(t *testing.T) {
	outputDir := "destination"
	testCases := []struct {
		file     File
		expected string
	}{
		{File{Name: "20170429_104146.jpg"}, "destination/2017/April/20170429_104146.jpg"},
		{File{Name: "20170512_025705.mp4"}, "destination/2017/May/20170512_025705.mp4"},
		{File{Name: "1560169653854.jpg"}, "destination/2019/June/1560169653854.jpg"},
	}

	t.Parallel()
	for _, tc := range testCases {
		dest := buildDestinationPath(tc.file, outputDir)

		if dest != tc.expected {
			t.Errorf("Expected date to be parsed from filename %s", tc.file.Name)
		}
	}
}
