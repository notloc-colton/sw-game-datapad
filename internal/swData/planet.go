package swData

// TODO: Are any or all of these fields required?!
type Planet struct {
	Name       string `json:"name,omitempty"`
	Climate    string `json:"climate,omitempty"`
	Population string `json:"population,omitempty"`
}
