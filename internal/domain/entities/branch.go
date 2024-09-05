package entities

// Branch representa una sucursal de un comercio
type Branch struct {
	ID         string `json:"id"`
	CommerceID string `json:"commerce_id"` // Referencia al comercio propietario
	Name       string `json:"name"`
	Address    string `json:"address"`
}
