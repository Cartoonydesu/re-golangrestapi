package skill

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSkills(t *testing.T) {
	t.Run("get all skill", func(t *testing.T) {
		router := gin.Default()
		w := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		if err != nil {
			log.Panic(err)
		}
		defer db.Close()
		rows := sqlmock.NewRows([]string{"key", "name", "description", "logo", "tags"}).
			AddRow("skill1", "Skill One", "Description for skill one", "logo1.png", pq.Array([]string{"tag1", "tag2"})).
			AddRow("skill2", "Skill Two", "Description for skill two", "logo2.png", pq.Array([]string{"tag3", "tag4"}))
		mock.ExpectQuery("SELECT key, name, description, logo, tags FROM skill").WillReturnRows(rows)
		h := &Handler{Db: db}
		path := "/api/v1/skills"
		router.GET(path, h.getAllSkills)
		req, _ := http.NewRequest("GET", path, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		expectedBody := `{"status":"success","data":[{"key":"skill1","name":"Skill One","description":"Description for skill one","logo":"logo1.png","tags":["tag1","tag2"]},{"key":"skill2","name":"Skill Two","description":"Description for skill two","logo":"logo2.png","tags":["tag3","tag4"]}]}`
		assert.JSONEq(t, expectedBody, w.Body.String())
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("get skill query error", func(t *testing.T) {
		router := gin.Default()
		w := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		if err != nil {
			log.Panic(err)
		}
		defer db.Close()
		mock.ExpectQuery("SELECT key, name, description, logo, tags FROM skill").
			WillReturnError(fmt.Errorf("query failed"))
		h := &Handler{Db: db}
		path := "/api/v1/skills"
		router.GET(path, h.getAllSkills)
		req, _ := http.NewRequest("GET", path, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedBody := `{"status":"error","message":"Can not get all skills"}`
		assert.JSONEq(t, expectedBody, w.Body.String())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
	t.Run("get all skill fail", func(t *testing.T) {
		router := gin.Default()
		w := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		if err != nil {
			log.Panic(err)
		}
		defer db.Close()
		rows := sqlmock.NewRows([]string{"name", "description", "logo", "tags"}).
			AddRow("Skill One", "Description for skill one", "logo1.png", pq.Array([]string{"tag1", "tag2"})).
			AddRow("Skill Two", "Description for skill two", "logo2.png", pq.Array([]string{"tag3", "tag4"}))
		mock.ExpectQuery("SELECT key, name, description, logo, tags FROM skill").WillReturnRows(rows)
		h := &Handler{Db: db}
		path := "/api/v1/skills"
		router.GET(path, h.getAllSkills)
		req, _ := http.NewRequest("GET", path, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		expectedBody := `{"status":"error", "message": "Can not get all skills"}`
		assert.JSONEq(t, expectedBody, w.Body.String())
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}

func TestGetSkillById(t *testing.T) {
	t.Run("get skill by id", func(t *testing.T) {
		router := gin.Default()
		w := httptest.NewRecorder()
		db, mock, err := sqlmock.New()
		if err != nil {
			log.Panic(err)
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"key", "name", "description", "logo", "tags"}).
			AddRow("skill1", "Skill One", "Description for skill one", "logo1.png", pq.Array([]string{"tag1", "tag2"}))
		mock.ExpectQuery("SELECT key, name, description, logo, tags FROM skill WHERE key = \\$1").
			WithArgs("skill1").
			WillReturnRows(rows)
		h := &Handler{Db: db}
		path := "/api/v1/skills/skill1"
		router.GET("/api/v1/skills/:key", h.getSkillById)
		req, _ := http.NewRequest("GET", path, nil)
		router.ServeHTTP(w, req)

		// Assert the response
		assert.Equal(t, http.StatusOK, w.Code)
		expectedBody := `{"status":"success","data":{"key":"skill1","name":"Skill One","description":"Description for skill one","logo":"logo1.png","tags":["tag1","tag2"]}}`
		assert.JSONEq(t, expectedBody, w.Body.String())

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Error(err)
		}
	})
}
