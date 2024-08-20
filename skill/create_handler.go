package skill

import (
	"cartoonydesu/response"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) createSkill(c *gin.Context) {
	var new Skill
	err := c.BindJSON(&new)
	if err != nil {
		response.BadRequest(c, "error", "Can not extract data from JSON")
	}
	stmt, err := h.Db.Prepare("INSERT INTO skill (key, name, description, logo, tags) VALUES($1, $2, $3, $4, $5) returning key;")
	if err != nil {
		response.BadRequest(c, "error", "Statement error")
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(new.Key, new.Name, new.Description, new.Logo, pq.Array(new.Tags)); err != nil {
		response.BadRequest(c, "error", "Skill already exists")
		return
	}
	response.Success(c, "success", new)
}
