package skill

import (
	"cartoonydesu/response"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) getAllSkills(c *gin.Context) {
	rows, err := h.Db.Query("SELECT key, name, description, logo, tags FROM skill;")
	if err != nil || rows.Err() != nil {
		response.BadRequest(c, "error", "Can not get all skills")
		return
	}
	defer rows.Close()
	var skills []Skill
	for rows.Next() {
		var s Skill
		err := rows.Scan(&s.Key, &s.Name, &s.Description, &s.Logo, pq.Array(&s.Tags))
		if err != nil {
			response.BadRequest(c, "error", "Can not get all skills")
			return
		}
		skills = append(skills, s)
	}
	response.Success(c, "success", skills)
}

func (h *Handler) getSkillById(c *gin.Context) {
	p := c.Param("key")
	s, err := h.getSkillByKey(p)
	if err != nil {
		response.InternalServerErr(c, "error", "Skill not found")
		return
	}
	response.Success(c, "success", s)
}

func (h *Handler) getSkillByKey(key string) (Skill, error) {
	skill := h.Db.QueryRow(fmt.Sprintf("SELECT key, name, description, logo, tags FROM skill WHERE key = '%v';", key))
	var s Skill
	err := skill.Scan(&s.Key, &s.Name, &s.Description, &s.Logo, pq.Array(&s.Tags))
	if err != nil {
		return Skill{}, err
	}
	return s, nil
}
