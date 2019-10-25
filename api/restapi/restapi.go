package restapi

import (
	"service/api"

	"github.com/gin-gonic/gin"
)

type Handle struct {
	Method   string
	Path     string
	Function func(ctx *gin.Context)
}

type restapi struct {
	cfg     api.Config
	eng     *gin.Engine
	handles []Handle
}

func (api *restapi) Start() error {
	err := api.eng.Run(":" + api.cfg.RestPort)
	if err != nil {
		return err
	}
	return nil
}

// InitRestAPI ...
func InitRestAPI(config api.Config) Client {
	rest := &restapi{
		cfg: config,
	}
	rest.initRestHandle()
	return rest
}

func (api *restapi) initRestHandle() {
	var restapiHandles = []Handle{
		{
			Method:   "GET",
			Path:     "/tasks",
			Function: api.GetTasks(),
		},
	}

	api.eng.Handle("GET", "/tasks", func(context *gin.Context) {

	})
	api.eng.Handle("POST", "/task")
	api.eng.Handle("GET", "/task/:id")
	api.eng.Handle("DELETE", "task/:id")
	api.eng.Handle("GET", "/nodes")
	api.eng.Handle("GET", "/videos")
}

func (api *restapi) GetTasks() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}
