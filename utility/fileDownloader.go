package utility

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// DownloadToFile downloads a file and returns the local filepath
func DownloadToFile(url string) string {
	file, fileErr := ioutil.TempFile("", "example")
	if fileErr != nil {
		panic(fileErr)
	}

	err := downloadFile(file.Name(), url)
	if err != nil {
		panic(err)
	}

	return file.Name()
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
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
