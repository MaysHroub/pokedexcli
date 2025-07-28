package pokeapi

type LocationAreaResponse struct {
	Count         int            `json:"count"`
	Next          *string        `json:"next"`
	Previous      *string        `json:"previous"`
	LocationAreas []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
