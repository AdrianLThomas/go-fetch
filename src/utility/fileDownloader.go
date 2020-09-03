package utility

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// HTTPClient interface
type HTTPClient interface {
	Get(url string) (resp *http.Response, err error)
}

var (
	// Client used for making HTTP calls
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

// DownloadToFile downloads a file and returns the local filepath
func DownloadToFile(url string) (string, error) {
	if len(url) <= 0 {
		return "", errors.New("URL is empty")
	}

	file, fileErr := ioutil.TempFile("", "go-spotify-img")
	if fileErr != nil {
		return "", fileErr
	}

	err := downloadFile(file.Name(), url)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func downloadFile(filepath string, url string) error {
	// Get the data
	resp, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
