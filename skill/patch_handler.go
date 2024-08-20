package skill

import (
	"cartoonydesu/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updateSkillName(c *gin.Context) {
	type name struct {
		Name string
	}
	p := c.Param("key")
	var n name
	err := c.BindJSON(&n)
	if err != nil {
		response.BadRequest(c, "error", "Can not extract data from json")
		return
	}
	stmt, err := h.Db.Prepare("UPDATE skill SET name = $1 where key = $2;")
	if err != nil {
		response.BadRequest(c, "error", "Statement error")
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(n.Name, p); err != nil {
		response.BadRequest(c, "error", "Not be able to update skill name")
		return
	}
	s, err := h.getSkillByKey(p)
	if err != nil {
		response.BadRequest(c, "error", "Skill not found")
		return
	}
	response.Success(c, "succes", s)
}
