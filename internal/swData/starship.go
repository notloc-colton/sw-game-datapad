package swData

// TODO: Are any or all of these fields required?!
type StarShip struct {
	Name          string `json:"name,omitempty"`
	CargoCapacity string `json:"cargoCapacity,omitempty"`
	Class         string `json:"class,omitempty"`
}
