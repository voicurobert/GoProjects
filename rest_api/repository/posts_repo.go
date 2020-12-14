package repository

import "github.com/voicurobert/GoProjects/rest_api/entity"

//PostRespository interface
type PostRespository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
