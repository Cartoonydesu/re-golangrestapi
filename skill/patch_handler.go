package skill

import (
	"cartoonydesu/response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) updateSkillName(c *gin.Context) {
	type name struct {
		Name string `json:"name"`
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
	response.Success(c, "success", s)
}

func (h *Handler) updateSkillDescription(c *gin.Context) {
	type desc struct  {
		Desc string `json:"description"`
	}
	p := c.Param("key")
	var d desc
	if err := c.BindJSON(&d); err != nil {
		response.BadRequest(c, "error", "Can not extract data from json")
		return
	}
	stmt, err := h.Db.Prepare("UPDATE skill SET description = $1 WHERE key = $2;")
	if err != nil {
		response.BadRequest(c, "error", "Statement error")
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(d.Desc, p); err != nil {
		response.BadRequest(c, "error", "Not be able to update skill description")
		return
	}
	s, err := h.getSkillByKey(p)
	if err != nil {
		response.InternalServerErr(c, "error", "Skill not found")
		return
	}
	response.Success(c, "success", s)	
}

func (h *Handler) updateSkillLogo(c *gin.Context) {
	p := c.Param("key")
	type logo struct {
		Logo string `json:"logo"`
	}
	var l logo
	err := c.BindJSON(&l)
	if err != nil {
		response.BadRequest(c, "error", "Can not extract data from json")
		return
	}
	stmt, err := h.Db.Prepare("UPDATE skill SET logo = $1 WHERE key = $2;")
	if err != nil {
		response.BadRequest(c, "error", "Statement error")
		return
	}
	defer stmt.Close()
	if _, err := stmt.Exec(l.Logo, p); err != nil {
		response.BadRequest(c, "error", "Not be able to update skill logo")
		return
	}
	s, err := h.getSkillByKey(p)
	if err != nil {
		response.InternalServerErr(c, "error", "Skill not found")
		return
	}
	response.Success(c, "success", s)
}
