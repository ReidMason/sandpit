package main

import (
	"fmt"
	"log"
	"mime"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

type File struct {
	Name     string
	Path     string
	MimeType string
}

type Job struct {
	id         int
	source     string
	destPath   string
	retryCount int
}

func worker(jobs chan Job, wg *sync.WaitGroup) {
	for job := range jobs {
		err := copyFile(job.source, job.destPath)
		if err != nil {
			job.retryCount++
			if job.retryCount <= 5 {
				jobs <- job
			} else {
				log.Printf("Job %d failed after 5 retries", job.id)
				wg.Done()
			}
		} else {
			wg.Done()
		}
	}
}

// find . -type f -newermt 1970-01-01 ! -newermt 1970-01-02

// rsync -avh --partial --info=progress2

// for f in *; do   fn=$(basename "$f");   mv "$fn" "$(date -r "$f" +"%Y-%m-%d_%H-%M-%S")_$fn"; done

func main() {
	// // sourceDir := "/Volumes/Documents/Samsung/DCIM/Camera"
	// // destinationDir := "/Volumes/Documents/Media/General"
	//
	sourceDir := os.Args[1]
	destinationDir := os.Args[2]

	if sourceDir == "" {
		panic("Source directory is required")
	}

	if destinationDir == "" {
		panic("Destination directory is required")
	}

	var wg sync.WaitGroup
	files := listFiles(sourceDir)
	log.Printf("Found %d files at source", len(files))
	filesCountDestStart := len(listFiles(destinationDir))

	jobs := make(chan Job, len(files))
	for w := 1; w <= 100; w++ {
		go worker(jobs, &wg)
	}

	count := 1
	for _, file := range files {
		destination := buildDestinationPath(file, destinationDir)
		wg.Add(1)
		jobs <- Job{id: count, source: file.Path, destPath: destination}
		count++
	}

	log.Printf("Waiting for all files to be copied...")
	wg.Wait()
	close(jobs)

	filesCountDestEnd := len(listFiles(destinationDir))
	log.Printf("Destination file count changed from %d to %d", filesCountDestStart, filesCountDestEnd)
	log.Printf("Files moved %d", filesCountDestEnd-filesCountDestStart)
}

func copyFile(source string, destination string) error {
	filename := filepath.Base(destination)
	if _, err := os.Stat(destination); err == nil && !strings.HasPrefix(filename, "2019") {
		return nil
	}

	return exec.Command("rsync", "-aP", "--mkpath", source, destination).Run()
}

func buildDestinationPath(file File, destinationDir string) string {
	date, err := parseFilenameDate(file.Name)
	if err != nil {
		log.Printf("Error parsing date from filename %s: %v", file.Name, err)
		return destinationDir + "/errors/" + file.Name
	}

	return destinationDir + "/" + fmt.Sprint(date.Year()) + "/" + date.Month().String() + "/" + file.Name
}

func parseFilenameDate(filename string) (time.Time, error) {
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))

	layouts := []string{"2006-01-02", "2006_01_02", "20060102"}

	for _, layout := range layouts {
		start := 0
		layoutSpan := len(layout)

		for start < len(filename)-layoutSpan {
			section := filename[start : start+layoutSpan]
			parsedTime, err := time.Parse(layout, section)
			if err == nil {
				return parsedTime, nil
			}
			start++
		}
	}

	// Try to parse unix timestamps
	if len(filename) > 10 {
		i, err := strconv.ParseInt(filename[:10], 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(i, 0), nil
	}
	return time.Time{}, fmt.Errorf("could not parse date from filename %s", filename)
}

func listFiles(sourceDir string) []File {
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		panic(err)
	}

	files := make([]File, 0)
	for _, entry := range entries {
		if entry.IsDir() {
			nestedFiles := listFiles(sourceDir + "/" + entry.Name())
			files = append(files, nestedFiles...)
		} else {
			mimeType := mime.TypeByExtension(filepath.Ext(entry.Name()))
			if strings.HasPrefix(entry.Name(), "._") || (!strings.HasPrefix(mimeType, "image/") && !strings.HasPrefix(mimeType, "video/")) {
				continue
			}
			files = append(files, File{Name: entry.Name(), Path: sourceDir + "/" + entry.Name(), MimeType: mimeType})
		}
	}

	return files
}
