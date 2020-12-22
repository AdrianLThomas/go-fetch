package utility

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
)

type MockClient struct {
	StatusCode int
}

func (m *MockClient) Get(url string) (resp *http.Response, err error) {
	r := ioutil.NopCloser(bytes.NewReader([]byte("hello")))

	if m.StatusCode == 404 {
		return nil, errors.New("Bad 404")
	} else {
		return &http.Response{
			StatusCode: m.StatusCode,
			Body:       r,
		}, nil
	}

}

func TestDownloadToFile_CalledWithUrl_ReturnsAFilePath(t *testing.T) {
	// arrange
	Client = &MockClient{StatusCode: 200}

	// act
	actualPath, err := DownloadToFile("http://www.example.com/image.jpg")

	// assert
	if err != nil {
		t.Error("returned unexpected error", err)
	}
	if !filepath.IsAbs(actualPath) {
		t.Errorf("actualPath (%s) is not an absolute path", actualPath)
	}
}

func TestDownloadToFile_CalledWithEmptyUrl_ReturnsError(t *testing.T) {
	// arrange
	Client = &MockClient{StatusCode: 200}

	// act
	actualPath, err := DownloadToFile("")

	// assert
	if err == nil {
		t.Error("didn't return the expected error", err)
		t.Log(actualPath)
	}
}

func TestDownloadToFile_404Url_ReturnsError(t *testing.T) {
	// arrange
	Client = &MockClient{StatusCode: 404}

	// act
	_, err := DownloadToFile("http://www.example.com/404")

	// assert
	if err == nil {
		t.Error("didn't return 404")
	}
}
