package skill

import "database/sql"

type Handler struct {
	Db *sql.DB
}

type Skill struct {
	Key string `json:"key" binding:"required"`
	Name string `json:"name"`
	Description string `json:"description"`
	Logo string `json:"logo"`
	Tags []string `json:"tags"`
}