package skill

import (
	"cartoonydesu/response"

	"github.com/gin-gonic/gin"
)

type UpdateSkill struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Tags        []string `json:"tags"`
}

func (h *Handler) updateSkill(c *gin.Context) {
	p := c.Param("key")
	var s UpdateSkill
	err := c.BindJSON(&s)
	if err != nil {
		response.BadRequest(c, "error", "Can not extract data from JSON")
	}
	stmt, err := h.Db.Prepare("UPDATE skill SET name = $1, description = $2, logo = $3, tags = $4 where key = $5;")
	if err != nil {
		response.BadRequest(c, "error", "Statement error")
	}
	defer stmt.Close()
	if _, err := stmt.Exec(s.Name, s.Description, s.Logo, s.Tags, p); err != nil {
		response.BadRequest(c, "error", "Not be able to update skill")
	}
	sk, err := h.getSkillByKey(p)
	if err != nil {
		response.InternalServerErr(c, "error", "Skill not found")
	}
	response.Success(c, "success", sk)
}
