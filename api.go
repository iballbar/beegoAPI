package beegoAPI

import (
	"fmt"

	"github.com/astaxie/beego"
)

//API for Client
type API struct {
	beego.Controller
}

type responseObject struct {
	Code           int         `json:"code"`
	Message        string      `json:"message"`
	ResponseObject interface{} `json:"responseObject"`
}

func (api *API) baseURL() string {
	var baseUrl string = api.Ctx.Input.Site() + fmt.Sprintf(":%d", api.Ctx.Input.Port())
	if api.Ctx.Input.Header("X-Forwarded-Host") != "" {
		baseUrl = api.Ctx.Input.Scheme() + "://" + api.Ctx.Input.Header("X-Forwarded-Host")
	}
	return baseUrl
}

func (api *API) ResponseJSON(results interface{}, code int, msg string) {
	response := &responseObject{
		Code:           code,
		Message:        msg,
		ResponseObject: results,
	}
	api.Data["json"] = response
	api.Ctx.Output.SetStatus(code)
	api.ServeJSON()
	return
}
