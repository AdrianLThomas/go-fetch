package model

// SearchResponse is the deserialised object returned from Spotify
type SearchResponse struct {
	Artists struct {
		Href     string           `json:"href"`
		Items    []ArtistResponse `json:"items"`
		Limit    int              `json:"limit"`
		Next     string           `json:"next"`
		Offset   int              `json:"offset"`
		Previous interface{}      `json:"previous"`
		Total    int              `json:"total"`
	} `json:"artists"`
}
