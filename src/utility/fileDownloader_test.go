package utility

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
)

type MockClient struct {
	GetFunc func(url string) (resp *http.Response, err error)
}

func (m *MockClient) Get(url string) (resp *http.Response, err error) {
	r := ioutil.NopCloser(bytes.NewReader([]byte("hello")))

	return &http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil
}

func TestDownloadToFile_CalledWithUrl_ReturnsAFilePath(t *testing.T) {
	// arrange
	Client = &MockClient{}

	// act
	actualPath := DownloadToFile("http://www.example.com/image.jpg")

	// assert
	if !filepath.IsAbs(actualPath) {
		t.Errorf("actualPath (%s) is not an absolute path", actualPath)
	}
}

// TODO test for nil url
// TODO test for http error (404?, 500..)
