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

type PokemonResponse struct {
	EncounterMethodRates any                 `json:"encounter_method_rates"`
	GameIndex            int                 `json:"game_index"`
	ID                   int                 `json:"id"`
	Location             any                 `json:"location"`
	Name                 string              `json:"name"`
	Names                any                 `json:"names"`
	PokemonEncounters    []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon        Pokemon `json:"pokemon"`
	VersionDetails any     `json:"version_details"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
