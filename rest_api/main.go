package main

import (
	"os"

	"github.com/voicurobert/GoProjects/rest_api/controller"
	router "github.com/voicurobert/GoProjects/rest_api/http"
	"github.com/voicurobert/GoProjects/rest_api/repository"
	"github.com/voicurobert/GoProjects/rest_api/service"
)

var (
	firebaseRepo   repository.PostRespository = repository.NewFirestoreRepository()
	postService    service.PostService        = service.NewPostService(firebaseRepo)
	postController controller.PostController  = controller.NewPostController(postService)
	httpRouter     router.Router              = router.NewMuxRouter()
)

func main() {

	httpRouter.GET("/posts", postController.GetPosts)

	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(os.Getenv("8000"))

}
