package skill

import (
	"cartoonydesu/response"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) createSkill(c *gin.Context) {
	var s Skill
	err := c.BindJSON(&s)
	if err != nil {
		response.BadRequest(c, "error", "Can not extract data from JSON")
		return
	}
	stmt, err := h.Db.Prepare("INSERT INTO skill (key, name, description, logo, tags) VALUES($1, $2, $3, $4, $5) returning key;")
	if err != nil {
		response.BadRequest(c, "error", "Statement error")
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(s.Key, s.Name, s.Description, s.Logo, pq.Array(s.Tags)); err != nil {
		response.BadRequest(c, "error", "Skill already exists")
		return
	}
	response.Success(c, "success", s)
}
