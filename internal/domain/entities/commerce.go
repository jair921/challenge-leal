package entities

// Commerce representa un comercio en la plataforma Leal
type Commerce struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	// BranchID representa la relación entre un comercio y una sucursal
	BranchID string `json:"branch_id"`
}
