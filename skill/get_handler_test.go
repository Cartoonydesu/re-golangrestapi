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

func TestGetSkillById(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.GET("/api/v1/skills/:key", h.getSkillById)
	t.Run("get the skill by key", func(t *testing.T) {
		s := Skill{
			Key:         "go",
			Name:        "Go",
			Description: "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software.",
			Logo:        "https://upload.wikimedia.org/wikipedia/commons/0/05/Go_Logo_Blue.svg",
			Tags:        []string{"go", "golang"},
		}
		jsonValue, _ := json.Marshal(s)
		req, _ := http.NewRequest("GET", "/api/v1/skills/go", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), string(jsonValue))
	})

	t.Run("get the skill with un exist key", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/skills/unexistedidforsure", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
}
