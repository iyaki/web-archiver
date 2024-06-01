package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

type URL struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URL    `xml:"url"`
}

func fetchSitemap(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching sitemap: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func parseSitemap(data []byte) (*URLSet, error) {
	var urlset URLSet
	err := xml.Unmarshal(data, &urlset)
	if err != nil {
		return nil, err
	}
	return &urlset, nil
}

func saveToWebArchive(urlToSave string) error {
	// Define the URL and the data to be sent
	apiURL := "https://web.archive.org/save/"
	data := url.Values{}
	data.Set("url", urlToSave) // Replace with actual relative path
	data.Set("capture_all", "on")
	data.Set("capture_outlinks", "on")
	data.Set("capture_screenshot", "on")

	// Create a new request using http
	req, err := http.NewRequest("POST", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		return err
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("LOW %s:%s", os.Getenv("WAYBACK_S3_ACCESS_KEY"), os.Getenv("WAYBACK_S3_SECRET_KEY")))

	// Send the request using http.Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Print the response status
	bodyResponse, _ := io.ReadAll(resp.Body)
	fmt.Println("Web Archive Response Status: ", resp.Status)
	fmt.Println("Web Archive Response: ", string(bodyResponse))

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("A sitemap URL as first argument is required")
		return
	}

	date := ""
	if len(os.Args) == 3 {
		date = os.Args[2]
	}

	url := os.Args[1]

	data, err := fetchSitemap(url)
	if err != nil {
		fmt.Printf("Error fetching sitemap: %v\n", err)
		return
	}

	urlset, err := parseSitemap(data)
	if err != nil {
		fmt.Printf("Error parsing sitemap: %v\n", err)
		return
	}

	for _, url := range urlset.URLs {
		fmt.Println(url)
		if url.Lastmod == "" || (date != "" && date <= url.Lastmod) {
			saveToWebArchive(url.Loc)
		} else {
			fmt.Printf("Skipping %q\n", url.Loc)
		}
	}
}
