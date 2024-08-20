package skill

import "github.com/gin-gonic/gin"

func SkillRouter(r *gin.Engine, h *Handler) {
	r.GET("/api/v1/skills", h.getAllSkills)
}
