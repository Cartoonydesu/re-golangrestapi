package skill

import (
	"bytes"
	"cartoonydesu/database"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestCreateSkill(t *testing.T) {
	db := database.NewPostgres()
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.POST("/api/v1/skills", h.createSkill)
	t.Run("create skill successfully", func(t *testing.T) {
		new := Skill{
			Key:         "testCreate",
			Name:        "Test",
			Description: "test",
			Logo:        "test",
			Tags:        []string{"test"},
		}
		jsonValue, _ := json.Marshal(new)
		req, _ := http.NewRequest("POST", "/api/v1/skills", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("create skill failed reading json", func(t *testing.T) {
		new := Skill{
			Key:         "testCreate",
			Name:        "", //Missing skill's name
			Description: "test",
			Logo:        "test",
			Tags:        []string{"test"},
		}
		jsonValue, _ := json.Marshal(new)
		req, _ := http.NewRequest("POST", "/api/v1/skills", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
