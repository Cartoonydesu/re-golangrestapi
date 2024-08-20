package skill

import (
	"cartoonydesu/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) deleteSkill(c *gin.Context) {
	p := c.Param("key")
	_, err := h.getSkillByKey(p)
	if err != nil {
		response.InternalServerErr(c, "error", "Skill not found")
		return
	}
	stmt, err := h.Db.Prepare("DELETE FROM skill WHERE key = $1;")
	if err != nil {
		response.BadRequest(c, "error", "Statement error")
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(p); err != nil {
		response.BadRequest(c, "error", "Not be able to delete skill")
		return
	}
	response.Success(c, "success", "Skill deleted")	
}