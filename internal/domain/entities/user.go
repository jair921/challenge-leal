package entities

// User representa a un usuario de la plataforma Leal
type User struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Points float64 `json:"points"`
}
