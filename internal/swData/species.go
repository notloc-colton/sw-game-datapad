package swData

// TODO: Are any or all of these fields required?!
type Species struct {
	ID              string `json:"-"`
	Name            string `json:"name,omitempty"`
	AverageLifespan string `json:"averageLifespan,omitempty"`
	Language        string `json:"language,omitempty"`
}
