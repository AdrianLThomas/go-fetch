package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/AdrianLThomas/go-fetch/model"
)

const (
	// BaseURL contains Spotify URL to v1 of the API
	BaseURL = "https://api.spotify.com/v1/"

	// DefaultClientTimeout is time to wait before cancelling the request
	DefaultClientTimeout time.Duration = 30 * time.Second
)

// SpotifyClient is for communicating with Spotify
type SpotifyClient struct {
	client  *http.Client
	baseURL string
}

// NewSpotifyClient returns a pointer to a newly created Spotify Client
func NewSpotifyClient() *SpotifyClient {
	return &SpotifyClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
	}
}

// SetTimeout overrides the default timeout
func (c *SpotifyClient) SetTimeout(t time.Duration) {
	c.client.Timeout = t
}

// Fetch retrieves the Artist data from Spotify
func (c *SpotifyClient) Fetch(artistName string, bearerToken string) (model.Artist, error) {
	resp, err := c.client.Get(c.baseURL + c.buildURL(artistName))

	if err != nil {
		return model.Artist{}, err
	}
	defer resp.Body.Close()

	var artistResponse model.ArtistResponse
	if err := json.NewDecoder(resp.Body).Decode(&artistResponse); err != nil {
		return model.Artist{}, err
	}

	return artistResponse.Artist()
}

func (c *SpotifyClient) buildURL(artistName string) string {
	return fmt.Sprintf("search?q=%s&type=artist", url.QueryEscape(artistName))
}
