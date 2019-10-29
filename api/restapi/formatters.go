package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseFailed(ctx *gin.Context, err error) {
	ctx.JSON(404, gin.H{
		"CODE":    404,
		"MESSAGE": err.Error(),
	})
}

func responseSuccess(ctx *gin.Context, v interface{}) {
	data, e := json.MarshalIndent(v, "", " ")
	if e != nil {
		ctx.JSON(200, gin.H{
			"CODE":    200,
			"MESSAGE": "SUCCESS",
			"DATA":    nil,
		})
	}

	ctx.JSON(200, gin.H{
		"CODE":    200,
		"MESSAGE": "SUCCESS",
		"DATA":    data,
	})
}

func formatterResponse(ctx *gin.Context, err error, v interface{}) {
	if err != nil {
		responseFailed(ctx, err)
		return
	}
	responseSuccess(ctx, v)
}

func decodeResponse(resp *http.Response, err error, v interface{}) error {
	if err != nil {
		return fmt.Errorf("response has some wrong:%w", err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body error:%w", err)
	}
	d := struct {
		Code    int    `json:"CODE"`
		Message string `json:"MESSAGE"`
		Data    []byte `json:"DATA"`
	}{}

	err = json.Unmarshal(bytes, &d)
	if err != nil {
		return fmt.Errorf("decode %s error:%w", string(bytes), err)
	}
	if d.Code != 200 {
		return fmt.Errorf("code status error:%d", d.Code)
	}

	if d.Data == nil {
		return nil
	}

	err = json.Unmarshal(d.Data, v)
	if err != nil {
		return fmt.Errorf("decode data error:%w", err)
	}
	return nil
}
