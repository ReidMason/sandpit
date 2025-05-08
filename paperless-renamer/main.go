package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const USERNAME = "your_username"
const PASSWORD = "your_password"
const HOST_URL = "paperless_url"

type ApiResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Document `json:"results"`
	All      []int      `json:"all"`
}

type Document struct {
	ID                  int           `json:"id"`
	Correspondent       int           `json:"correspondent"`
	DocumentType        int           `json:"document_type"`
	StoragePath         int           `json:"storage_path"`
	Title               string        `json:"title"`
	Content             string        `json:"content"`
	Tags                []int         `json:"tags"`
	Created             string        `json:"created"`
	CreatedDate         string        `json:"created_date"`
	Modified            string        `json:"modified"`
	Added               string        `json:"added"`
	DeletedAt           string        `json:"deleted_at"`
	ArchiveSerialNumber uint64        `json:"archive_serial_number"`
	OriginalFileName    string        `json:"original_file_name"`
	ArchivedFileName    string        `json:"archived_file_name"`
	Owner               int           `json:"owner"`
	Permissions         Permissions   `json:"permissions"`
	UserCanChange       bool          `json:"user_can_change"`
	IsSharedByRequester bool          `json:"is_shared_by_requester"`
	Notes               []Note        `json:"notes"`
	CustomFields        []CustomField `json:"custom_fields"`
	PageCount           int           `json:"page_count"`
	MimeType            string        `json:"mime_type"`
}

type Permissions struct {
	View   PermissionDetail `json:"view"`
	Change PermissionDetail `json:"change"`
}

type PermissionDetail struct {
	Users  []int `json:"users"`
	Groups []int `json:"groups"`
}

type Note struct {
	ID      int    `json:"id"`
	Note    string `json:"note"`
	Created string `json:"created"`
	User    User   `json:"user"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type CustomField struct {
	Value string `json:"value"`
	Field int    `json:"field"`
}

type DocPatch struct {
	Title string `json:"title"`
}

func main() {
	url := HOST_URL + "/api/documents/?tags__name__iexact=payslip&page_size=1000"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("accept", "application/json")
	req.SetBasicAuth(USERNAME, PASSWORD)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var apiResp ApiResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Fatalf("Error unmarshaling response: %v", err)
	}

	log.Printf("Total documents to process: %d", len(apiResp.Results))
	for _, document := range apiResp.Results {
		newName := "Payslip - " + getDocumentDate(document.Content)
		log.Printf("Renaming document ID %d to '%s'...", document.ID, newName)
		if document.Title != newName {
			RenameDocument(document.ID, newName)
		} else {
			log.Printf("Document ID %d already has the correct title. Skipping.", document.ID)
		}
	}
	log.Println("All documents processed.")
}

func getDocumentDate(content string) string {
	words := strings.Split(content, " ")
	for _, word := range words {
		// parse date
		date, err := time.Parse("02-01-2006", word)
		if err == nil {
			return date.Format("02-01-2006")
		}
	}

	panic("No date found")
}

func RenameDocument(id int, title string) {
	url := fmt.Sprintf("%s/api/documents/%d/", HOST_URL, id)

	requestBody := DocPatch{
		Title: title,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error marshaling body: %v", err)
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Set("accept", "application/json")
	req.SetBasicAuth(USERNAME, PASSWORD)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Error renaming document: %v", string(body))
	}

	log.Printf("Successfully renamed document ID %d to '%s'", id, title)
}
