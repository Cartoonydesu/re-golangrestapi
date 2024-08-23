package skill

// func setupTestDB() *sql.DB {
// 	db, _ := sql.Open("sqlite3", ":memory:")
// 	// Create the table
// 	db.Exec(`CREATE TABLE skill (
//         key TEXT PRIMARY KEY,
//         name TEXT,
//         description TEXT,
//         logo TEXT,
//         tags TEXT
//     )`)
// 	return db
// }

// func TestPatchSkillName(t *testing.T) {
// 	db := setupTestDB()
// 	defer db.Close()
// 	h := Handler{Db: db}
// 	r := gin.Default()
// 	r.POST("/api/v1/skills", h.createSkill)
// 	r.PUT("/api/v1/skills/:key/action/name", h.updateSkillName)

// 	new := Skill{
// 		Key:         "testUpdateName",
// 		Name:        "Test",
// 		Description: "test",
// 		Logo:        "test",
// 		Tags:        []string{"test"},
// 	}
// 	jsonValue, _ := json.Marshal(new)
// 	req, _ := http.NewRequest("POST", "/api/v1/skills", bytes.NewBuffer(jsonValue))
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	t.Run("update skill name", func(t *testing.T) {
// 		name := map[string]string{
// 			"name": "updateName",
// 		}
// 		jsonValue, _ := json.Marshal(name)
// 		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/name", new.Key), bytes.NewBuffer(jsonValue))
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
// 		assert.Equal(t, http.StatusOK, w.Code)
// 	})

// 	t.Run("update skill name to non-existent skill", func(t *testing.T) {
// 		name := map[string]string{
// 			"name": "updateName",
// 		}
// 		jsonValue, _ := json.Marshal(name)
// 		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/name", "nonexistentkey"), bytes.NewBuffer(jsonValue))
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 	})

// 	t.Run("update skill name with empty skill name", func(t *testing.T) {
// 		name := map[string]string{
// 			"name": "",
// 		}
// 		jsonValue, _ := json.Marshal(name)
// 		req, _ := http.NewRequest("PUT", fmt.Sprintf("/api/v1/skills/%v/action/name", new.Key), bytes.NewBuffer(jsonValue))
// 		w := httptest.NewRecorder()
// 		r.ServeHTTP(w, req)
// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})
// }
