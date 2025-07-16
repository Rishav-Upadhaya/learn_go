package todo

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Status 	bool `json:"status"`
}

// Date
func GetDummyTodos() []Todo {
	return []Todo{
		{ID: 1, Title: "Learn Go", Status: true},
		{ID: 2, Title: "Build a REST API", Status: false},
		{ID: 3, Title: "Deploy to Production", Status: false},
	}
}