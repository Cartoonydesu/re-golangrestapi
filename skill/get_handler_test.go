package skill

import (
	"cartoonydesu/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetSkillById(t *testing.T) {
	db := database.NewPostgres()
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.GET("/api/v1/skills/:key", h.getSkillById)
	t.Run("get skill by key", func(t *testing.T) {
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

	t.Run("get skill with unexisted key", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/skills/unexistedidforsure", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
		assert.Contains(t, w.Body.String(), "error")
	})
}

func TestGetAllSkills(t *testing.T) {
	db := database.NewPostgres()
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.GET("/api/v1/skills", h.getAllSkills)
	t.Run("get all skills", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/v1/skills", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "success")
	})
}
