package restapi

import (
	"service/api"

	"github.com/gin-gonic/gin"
)

type restapi struct {
	cfg api.Config
	eng *gin.Engine
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
	eng := gin.Default()
	rest := &restapi{
		cfg: config,
		eng: eng,
	}
	rest.initRestHandle()
	return rest
}

func (*restapi) initRestHandle() {

}
