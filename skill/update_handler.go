package skill

type UpdateSkill struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Logo string `json:"logo"`
	Tags []string `json:"tags"`
}