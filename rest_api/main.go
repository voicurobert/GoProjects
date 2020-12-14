package main

import (
	"github.com/voicurobert/GoProjects/rest_api/controller"
	router "github.com/voicurobert/GoProjects/rest_api/http"
	"github.com/voicurobert/GoProjects/rest_api/repository"
	"github.com/voicurobert/GoProjects/rest_api/service"
)

var (
	firebaseRepo   repository.PostRespository = repository.NewFirestoreRepository()
	postService    service.PostService        = service.NewPostService(firebaseRepo)
	postController controller.PostController  = controller.NewPostController(postService)
	httpRouter     router.Router              = router.NewChiRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/posts", postController.GetPosts)

	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)

}
