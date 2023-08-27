package repositories

import "yet_new_pet_blog/internals/models"

type PostsRepository interface {
	GetPosts() models.Posts
}

type postsRepository struct {
}

func NewPostsRepository() PostsRepository {
	return &postsRepository{}
}

func (p *postsRepository) GetPosts() models.Posts {
	return models.Posts{
		PostsList: []models.Post{},
	}
}
