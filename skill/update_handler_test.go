package skill

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUpdateSkill(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.POST("/api/v1/skills", h.createSkill)
	r.PUT("/api/v1/skills/:key", h.updateSkill)

	t.Run("update skill succesfully", func(t *testing.T) {
		new := Skill{
			Key:         "testUpdate",
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
		updateS := UpdateSkill{
			Name:        "TestEdit",
			Description: "testEdit",
			Logo:        "testEdit",
			Tags:        []string{"testEdit", "somethingNew"},
		}
		jsonValue, _ = json.Marshal(updateS)
		req = httptest.NewRequest("PUT", "/api/v1/skills/testUpdate", bytes.NewBuffer(jsonValue))
		w = httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
