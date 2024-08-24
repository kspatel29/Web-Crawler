package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// getHTML fetches the HTML content from the given URL.
func getHTML(rawURL string) (string, error) {
        // Create an HTTP client that follows redirects
        client := &http.Client{
                CheckRedirect: func(req *http.Request, via []*http.Request) error {
                        return nil // Allow all redirects
                },
        }

        // Send GET request
        resp, err := client.Get(rawURL)
        if err != nil {
                return "", fmt.Errorf("failed to fetch the URL: %v", err)
        }
        defer resp.Body.Close()

        // Check for HTTP error status
        if resp.StatusCode >= 400 {
                return "", fmt.Errorf("HTTP error: %d", resp.StatusCode)
        }

        // Check Content-Type for HTML
        contentType := resp.Header.Get("Content-Type")
        if !strings.Contains(strings.ToLower(contentType), "text/html") {
                return "", fmt.Errorf("content type is not HTML: %s", contentType)
        }

        // Handle possible compression
        var reader io.ReadCloser
        if resp.Header.Get("Content-Encoding") == "gzip" {
                reader, err = gzip.NewReader(resp.Body)
                if err != nil {
                        return "", err
                }
                defer reader.Close()
        } else {
                reader = resp.Body
        }

        // Read the body
        bodyBytes, err := io.ReadAll(reader)
        if err != nil {
                return "", fmt.Errorf("failed to read response body: %v", err)
        }

        return string(bodyBytes), nil
}