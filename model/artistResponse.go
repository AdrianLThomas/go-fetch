package model

// ArtistResponse represents a struct for the JSON response
type ArtistResponse struct {
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Genres     []string      `json:"genres"`
	Images     []ArtistImage `json:"images"`
	Name       string        `json:"name"`
	Popularity int           `json:"popularity"`
	Type       string        `json:"type"`
	URI        string        `json:"uri"`
}

// Artist returns an Artist given an ArtistResponse
func (ar ArtistResponse) Artist() Artist {
	return Artist{
		Name:       ar.Name,
		Popularity: ar.Popularity,
		Type:       ar.Type,
		Followers:  ar.Followers.Total,
		Genres:     ar.Genres,
		Image:      ar.Images[0],
	}
}
