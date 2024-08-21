package skill

import (
	"bytes"
	"cartoonydesu/database"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPatchSkill(t *testing.T) {
	db := database.NewPostgres()
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.POST("/api/v1/skills", h.createSkill)
	r.PUT("/api/v1/skills/:key/action/name", h.updateSkillName)
	r.PUT("/api/v1/skills/:key/action/description", h.updateSkillDescription)
	r.PUT("/api/v1/skills/:key/action/logo", h.updateSkillLogo)
	r.PUT("/api/v1/skills/:key/action/tags", h.updateSkillTags)

	new := Skill{
		Key:         "testUpdateAction",
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

	t.Run("update skill name", func(t *testing.T) {
		name := name{
			Name: "updateName",
		}
		jsonValue, _ := json.Marshal(name)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/name", new.Key), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("update unexisted skill", func(t *testing.T) {
		name := name{
			Name: "updateName",
		}
		jsonValue, _ := json.Marshal(name)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/name", "unexistedkey"), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("update skill with empty skill name", func(t *testing.T) {
		name := name{
			Name: "",
		}
		jsonValue, _ := json.Marshal(name)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/name", new.Key), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
