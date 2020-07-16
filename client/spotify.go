package client

import (
	"encoding/json"
	"errors"
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
	client *http.Client
	token  string
}

// NewSpotifyClient returns a pointer to a newly created Spotify Client
func NewSpotifyClient(authToken string) *SpotifyClient {
	return &SpotifyClient{
		client: &http.Client{
			Timeout: DefaultClientTimeout,
		},
		token: authToken,
	}
}

// SetTimeout overrides the default timeout
func (c *SpotifyClient) SetTimeout(t time.Duration) {
	c.client.Timeout = t
}

// Fetch retrieves the Artist data from Spotify
func (c *SpotifyClient) Fetch(artistName string) (model.Artist, error) {
	url := BaseURL + c.buildURL(artistName)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))
	resp, err := c.client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return model.Artist{}, err
	}
	if resp.StatusCode == 401 {
		return model.Artist{}, errors.New("Spotify: Unauthorised")
	}

	var artistResponse model.ArtistResponse
	if err := json.NewDecoder(resp.Body).Decode(&artistResponse); err != nil {
		return model.Artist{}, err
	}

	return artistResponse.Artist(), nil
}

func (c *SpotifyClient) buildURL(artistName string) string {
	return fmt.Sprintf("search?q=%s&type=artist", url.QueryEscape(artistName))
}
