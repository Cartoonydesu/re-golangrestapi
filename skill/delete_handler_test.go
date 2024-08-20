package skill

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.POST("/api/v1/skills", h.createSkill)
	r.DELETE("/api/v1/skills/:key", h.deleteSkill)
	// t.Run("delete skill", func(t *testing.T) {
	// 	s := Skill{
	// 		Key:         "testDelete",
	// 		Name:        "Test",
	// 		Description: "test",
	// 		Logo:        "test",
	// 		Tags:        []string{"test"},
	// 	}
	// 	jsonValue := json
	// })
	t.Run("delete skill by unexisted key", func(t *testing.T) {
		req, _ := http.NewRequest("DELETE", "/api/v1/skills/unexistedidforsure", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
}
