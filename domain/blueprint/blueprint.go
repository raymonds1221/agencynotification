package blueprint

// Blueprint is our blueprint domain model
type Blueprint struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

// NewBlueprint create new instance of Blueprint domain
func NewBlueprint(id string, name string, desc string) *Blueprint {
	return &Blueprint{
		ID:   id,
		Name: name,
		Desc: desc,
	}
}
