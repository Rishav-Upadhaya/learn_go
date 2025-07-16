package post

type Post struct {
	ID      int    `json:"id"`
	Username   string `json:"username" binding:"required"`
	Caption string `json:"caption"`
	ImageURL string `json:"image_url" binding:"required"`
}

func GetDummyPosts() []Post {
	return []Post{
		{ID: 1, Username: "user1", Caption: "First post", ImageURL: "http://example.com/image1.jpg"},
		{ID: 2, Username: "user2", Caption: "Second post", ImageURL: "http://example.com/image2.jpg"},
		{ID: 3, Username: "user3", Caption: "Third post", ImageURL: "http://example.com/image3.jpg"},
	}
}
