package restapi

import (
	"service/api"

	"github.com/gin-gonic/gin"
	"github.com/goextension/log"
)

// RestAPI ...
type RestAPI interface {
	api.Client
}

// Handle ...
type Handle struct {
	Method   string
	Path     string
	Function func(ctx *gin.Context)
}

type restapi struct {
	version string
	cfg     api.Config
	eng     *gin.Engine
	handles []Handle
	manager *api.Manager
}

// AddWork ...
func (r *restapi) AddWork(manager *api.Manager, work api.Work) error {
	panic("implement me")
}

// GetNode ...
func (r *restapi) GetNode(manager *api.Manager, id string) {
	panic("implement me")
}

// AddNode ...
func (r *restapi) AddNode(manager *api.Manager, node api.Node) {
	panic("implement me")
}

// GetWork ...
func (r *restapi) GetWork(manager *api.Manager, id string) (*api.Work, error) {
	panic("implement me")
}

// GetWorks ...
func (r *restapi) GetWorks(manager *api.Manager) ([]*api.Work, error) {
	return nil, nil
}

// DeleteWork ...
func (r *restapi) DeleteWork(manager *api.Manager, id string) error {
	panic("implement me")
}

// PostNode ...
func (r *restapi) PostNode(manager *api.Manager) {
	panic("implement me")
}

// DeleteNode ...
func (r *restapi) DeleteNode(manager *api.Manager) {
	panic("implement me")
}

// GetVideos ...
func (r *restapi) GetVideos(manager *api.Manager) {
	panic("implement me")
}

// Stop ...
func (r *restapi) Stop() {

}

func (r *restapi) init() {
	if r.manager == nil {
		log.Panicw("manager is null")
	}

	var restapiHandles = []Handle{
		{
			Method:   "GET",
			Path:     "/tasks",
			Function: r.getTasks(),
		},
		{
			"POST",
			"/task",
			r.postTask(),
		},
		{
			"GET",
			"/task/:id",
			r.getTask(),
		},
		{
			"DELETE",
			"/task/:id",
			r.deleteTask(),
		},
		{
			"GET",
			"/nodes",
			r.getNode(),
		},
		{
			"POST",
			"/node",
			r.postNode(),
		},
		{
			"GET",
			"/node/:id",
			r.getNode(),
		},
		{
			"DELETE",
			"/node/:id",
			r.deleteNode(),
		},
		{
			"GET",
			"videos",
			r.getVideos(),
		},
	}

	v := r.eng.Group(r.version)

	for _, h := range restapiHandles {
		v.Handle(h.Method, h.Path, h.Function)
	}
}

// Start ...
func (r *restapi) Start() error {
	err := r.eng.Run(":" + r.cfg.RestPort)
	if err != nil {
		return err
	}
	return nil
}

func (r *restapi) getTasks() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		tasks, e := r.manager.Client(api.RPCClient).GetWorks(r.manager)
		formatterResponse(ctx, e, tasks)
	}
}

func (r *restapi) postTask() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

	}
}

func (r *restapi) getTask() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

	}
}

func (r *restapi) deleteTask() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

	}
}

func (r *restapi) getNode() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

	}
}

func (r *restapi) postNode() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

	}
}

func (r *restapi) deleteNode() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

	}
}

func (r *restapi) getVideos() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

	}
}

// Options ...
type Options func(restapi *restapi)

// Manager ...
func Manager(manager *api.Manager) Options {
	return func(restapi *restapi) {
		restapi.manager = manager
	}
}

// NewRestAPI ...
func NewRestAPI(config api.Config, options ...Options) RestAPI {
	rest := &restapi{
		version: "v0",
		cfg:     config,
		eng:     gin.Default(),
	}
	for _, option := range options {
		option(rest)
	}
	rest.init()
	return rest
}
