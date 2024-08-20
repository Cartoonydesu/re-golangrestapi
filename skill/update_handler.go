package skill

import (
	"cartoonydesu/response"

	"github.com/gin-gonic/gin"
)

type UpdateSkill struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Logo string `json:"logo"`
	Tags []string `json:"tags"`
}

func (h *Handler) updateSkill(c *gin.Context) {
	var s UpdateSkill
	err := c.BindJSON(&s)
	// c.ShouldBindJSON()
	if err != nil {
		response.BadRequest(c, "error", "Can not extract data from JSON")
	}
	// stmt, err := h.D
}