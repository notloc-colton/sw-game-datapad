package swData

// TODO: Are any or all of these fields required?!
type Planet struct {
	ID         string `json:"-"`
	Name       string `json:"name,omitempty"`
	Climate    string `json:"climate,omitempty"`
	Population string `json:"population,omitempty"`
}
