// Package fileutil provide some utils to help with files
package fileutil

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// GetExtension return extension of file
func GetExtension(name string) string {
	i := strings.LastIndex(name, ".")
	if i < 0 {
		return ""
	}
	return name[i+1:]
}

// GetName return name of file without extension
func GetName(name string) string {
	i := strings.LastIndex(name, ".")
	if i < 0 {
		return name
	}
	return name[:i]
}

// Download url into directoy. Return os.File if method is successful
func Download(url string, dir string) (*os.File, error) {
	// File name
	_, fName := filepath.Split(url)
	fPath := filepath.Join(dir, fName)

	// Open URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Create File
	f, err := os.Create(fPath)
	if err != nil {
		return nil, err
	}

	// Copy URL body to new file
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return nil, err
	}

	return f, nil
}
