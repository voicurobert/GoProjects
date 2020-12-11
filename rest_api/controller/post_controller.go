package controller

import (
	"encoding/json"
	"net/http"

	"github.com/voicurobert/GoProjects/rest_api/entity"
	"github.com/voicurobert/GoProjects/rest_api/errors"
	"github.com/voicurobert/GoProjects/rest_api/service"
)

var (
	postService service.PostService
)

type controller struct {
}

//NewPostController constructor
func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

//PostController interface
type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

//GetPosts handler
func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error gettings posts array"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

// AddPost handler adds a new Post
func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling request"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
