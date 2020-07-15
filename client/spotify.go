package client

import (
	"net/http"
	"time"
)

const (
	// BaseURL contains Spotify URL to Drake
	BaseURL = "https://api.spotify.com/v1/artists/3TVXtAsR1Inumwj472S9r4"

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

// TODO: Fetch()
// TODO: SaveToDisk()
// TODO: buildURL()
