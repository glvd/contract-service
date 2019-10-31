package restapi

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"service/api"

	"github.com/gin-gonic/gin"
	"github.com/goextension/log"
)

// RestModeServer ...
const RestModeServer = "server"

// RestModeClient ...
const RestModeClient = "client"

// RestAPI ...
type RestAPI interface {
	api.Runnable
	api.Client
}

// Handle ...
type Handle struct {
	Method   string
	Path     string
	Function func(ctx *gin.Context)
}

type restapi struct {
	mode    string
	version string
	cfg     api.Config
	eng     *gin.Engine
	handles []Handle
	manager *api.Manager
}

// AddWork ...
func (r *restapi) AddWork(manager *api.Manager, work api.Work) error {
	workBytes, e := json.Marshal(work)
	if e != nil {
		return e
	}

	resp, e := http.Post(r.URL("work"), "application/json", bytes.NewReader(workBytes))
	if e != nil {
		return e
	}
	var v interface{}
	return decodeResponse(resp, e, &v)
}

// GetNode ...
func (r *restapi) GetNode(manager *api.Manager, id string) {
	panic("implement me")
}

// AddNode ...
func (r *restapi) AddNode(manager *api.Manager, node api.Node) {
	panic("")
}

// GetWork ...
func (r *restapi) GetWork(manager *api.Manager, id string) (*api.Work, error) {
	panic("implement me")
}

// GetWorks ...
func (r *restapi) GetWorks(manager *api.Manager) ([]*api.Work, error) {
	var works []*api.Work
	resp, err := http.Get(r.URL("works"))
	err = decodeResponse(resp, err, &works)
	if err != nil {
		return nil, err
	}
	return works, nil
}

// DeleteWork ...
func (r *restapi) DeleteWork(manager *api.Manager, id string) error {
	req, e := http.NewRequest("DELETE", r.URL("work", id), nil)
	if e != nil {
		return e
	}
	resp, e := http.DefaultClient.Do(req)
	if e != nil {
		return e
	}
	var v interface{}
	return decodeResponse(resp, e, &v)
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

// URL ...
func (r *restapi) URL(prefix string, args ...string) string {
	s := []string{"http://" + r.cfg.Remote + ":" + r.cfg.RestPort, r.version, prefix}
	s = append(s, args...)
	return strings.Join(s, "/")
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
			"GET",
			"/ping",
			func(ctx *gin.Context) {
				ctx.JSON(200, gin.H{"MESSAGE": "success"})
			},
		},
		{
			Method:   "GET",
			Path:     "/works",
			Function: r.getWorks(),
		},
		{
			"POST",
			"/work",
			r.addWork(),
		},
		{
			"GET",
			"/work/:id",
			r.getWork(),
		},
		{
			"DELETE",
			"/work/:id",
			r.deleteWork(),
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
	log.Info("rest was handle on port:", r.cfg.RestPort)
	err := r.eng.Run(":" + r.cfg.RestPort)
	if err != nil {
		return err
	}
	return nil
}

func (r *restapi) getWorks() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		works, e := r.manager.Client(api.RPCClient).GetWorks(r.manager)
		formatterResponse(ctx, e, works)
	}
}

func (r *restapi) addWork() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		bytes, e := ioutil.ReadAll(ctx.Request.Body)
		if e != nil {
			formatterResponse(ctx, e, nil)
			return
		}
		var work api.Work
		if err := json.Unmarshal(bytes, &work); err != nil {
			formatterResponse(ctx, err, nil)
			return
		}

		err := r.manager.Client(api.RPCClient).AddWork(r.manager, work)
		formatterResponse(ctx, err, nil)
	}
}

func (r *restapi) getWork() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {

	}
}

func (r *restapi) deleteWork() func(ctx *gin.Context) {

	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		err := r.manager.Client(api.RPCClient).DeleteWork(r.manager, id)
		formatterResponse(ctx, err, nil)
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

// ClientMode ...
func ClientMode() Options {
	return func(restapi *restapi) {
		restapi.mode = RestModeClient
	}
}

// NewRestAPI ...
func NewRestAPI(config api.Config, options ...Options) RestAPI {
	rest := &restapi{
		mode:    RestModeServer,
		version: "v0",
		cfg:     config,
	}
	for _, option := range options {
		option(rest)
	}

	if rest.mode == RestModeServer {
		rest.eng = gin.Default()
		rest.init()
	}
	return rest
}
