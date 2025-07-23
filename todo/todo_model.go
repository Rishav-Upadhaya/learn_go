package todo

type Todo struct {
	ID          int    `gorm:"primary_key; auto_increment; not mull;" json:"id"`
	Title       string `gorm:"not null; type:varchar(255); default:'';" json:"title" binding:"required"`
	Status 	bool `gorm:"not null; default: true;" json:"status"`
}

// Date
func GetDummyTodos() []Todo {
	return []Todo{
		{ID: 1, Title: "Learn Go", Status: true},
		{ID: 2, Title: "Build a REST API", Status: false},
		{ID: 3, Title: "Deploy to Production", Status: false},
	}
}