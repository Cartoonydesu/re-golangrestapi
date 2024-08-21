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

func TestPatchSkillName(t *testing.T) {
	db := database.NewPostgres()
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.POST("/api/v1/skills", h.createSkill)
	r.PUT("/api/v1/skills/:key/action/name", h.updateSkillName)

	new := Skill{
		Key:         "testUpdateName",
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

	t.Run("update skill name to unexisted skill", func(t *testing.T) {
		name := name{
			Name: "updateName",
		}
		jsonValue, _ := json.Marshal(name)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/name", "unexistedkey"), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("update skill name with empty skill name", func(t *testing.T) {
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

func TestPatchSkillDescription(t *testing.T) {
	db := database.NewPostgres()
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.POST("/api/v1/skills", h.createSkill)
	r.PUT("/api/v1/skills/:key/action/description", h.updateSkillDescription)

	new := Skill{
		Key:         "testUpdateDescription",
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

	t.Run("update skill description", func(t *testing.T) {
		desc := desc{
			Desc: "updateLogo",
		}
		jsonValue, _ := json.Marshal(desc)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/description", new.Key), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("update skill description to unexisted skill", func(t *testing.T) {
		desc := desc{
			Desc: "updateLogo",
		}
		jsonValue, _ := json.Marshal(desc)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/description", "unexistedkey"), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("update skill description with empty skill description", func(t *testing.T) {
		desc := desc{
			Desc: "",
		}
		jsonValue, _ := json.Marshal(desc)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/description", new.Key), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestPatchSkillLogo(t *testing.T) {
	db := database.NewPostgres()
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.POST("/api/v1/skills", h.createSkill)
	r.PUT("/api/v1/skills/:key/action/logo", h.updateSkillLogo)

	new := Skill{
		Key:         "testUpdateLogo",
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

	t.Run("update skill logo", func(t *testing.T) {
		logo := logo{
			Logo: "updateLogo",
		}
		jsonValue, _ := json.Marshal(logo)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/logo", new.Key), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("update skill logo to unexisted skill", func(t *testing.T) {
		logo := logo{
			Logo: "updateLogo",
		}
		jsonValue, _ := json.Marshal(logo)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/logo", "unexistedkey"), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("update skill logo with empty skill logo", func(t *testing.T) {
		logo := logo{
			Logo: "",
		}
		jsonValue, _ := json.Marshal(logo)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/logo", new.Key), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestPatchSkillTags(t *testing.T) {
	db := database.NewPostgres()
	defer db.Close()
	h := Handler{Db: db}
	r := gin.Default()
	r.POST("/api/v1/skills", h.createSkill)
	r.PUT("/api/v1/skills/:key/action/tags", h.updateSkillTags)

	new := Skill{
		Key:         "testUpdateTags",
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

	t.Run("update skill tags", func(t *testing.T) {
		tags := tags{
			Tags: []string{"update"},
		}
		jsonValue, _ := json.Marshal(tags)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/tags", new.Key), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("update skill tags to unexisted skill", func(t *testing.T) {
		tags := tags{
			Tags: []string{"update"},
		}
		jsonValue, _ := json.Marshal(tags)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/tags", "unexistedkey"), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("update skill tags with empty skill tags", func(t *testing.T) {
		tags := tags{
			Tags: nil,
		}
		jsonValue, _ := json.Marshal(tags)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/tags", new.Key), bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
