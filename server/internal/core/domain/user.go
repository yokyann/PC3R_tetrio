package domain

// User represents a user entity.
type User struct {
	ID     string `json:"_id"`
	Pseudo string `json:"pseudo"`
	Mail   string `json:"mail"`
	// Add more fields as needed
}
