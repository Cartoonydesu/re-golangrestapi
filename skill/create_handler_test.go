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
	_ "github.com/lib/pq"
)

func TestCreateSkill(t *testing.T) {
	db, _ := sql.Open("postgres", os.Getenv("POSTGRES_URI"))
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
}
