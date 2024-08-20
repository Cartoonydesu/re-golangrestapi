package skill

import "github.com/gin-gonic/gin"

func SkillRouter(r *gin.Engine, h *Handler) {
	r.GET("/api/v1/skills", h.getAllSkills)
	r.GET("/api/v1/skills/:key", h.getSkillById)
	r.POST("/api/v1/skills", h.createSkill)
	r.PUT("/api/v1/skills/:key", h.updateSkill)
	r.PATCH("/api/v1/skills/:key/action/name", h.updateSkillName)
	r.PATCH("/api/v1/skills/:key/action/description", h.updateSkillDescription)
	r.PATCH("/api/v1/skills/:key/action/logo", h.updateSkillLogo)
}
