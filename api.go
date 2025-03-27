package beegoAPI

import (
	"fmt"

	beego "github.com/beego/beego/v2/adapter"
)

// API for Client
type API struct {
	beego.Controller
}

// ResponseObject represents a standard API response with int code
type ResponseObject struct {
	Code           int         `json:"code"`
	Message        string      `json:"message"`
	ResponseObject interface{} `json:"responseObject"`
}

// ResponseObjectWithCode represents a standard API response with int64 code
type ResponseObjectWithCode struct {
	Code           int64       `json:"code"`
	Message        string      `json:"message"`
	ResponseObject interface{} `json:"responseObject"`
}

// BaseURL returns the base URL for the current request
func (api *API) BaseURL() string {
	baseUrl := api.Ctx.Input.Site() + fmt.Sprintf(":%d", api.Ctx.Input.Port())
	if api.Ctx.Input.Header("X-Forwarded-Host") != "" {
		baseUrl = api.Ctx.Input.Scheme() + "://" + api.Ctx.Input.Header("X-Forwarded-Host")
	}
	return baseUrl
}

// ResponseJSON sends a JSON response with the given results, code, and message
func (api *API) ResponseJSON(results interface{}, code int, msg string) {
	if results == nil {
		results = struct{}{}
	}
	response := &ResponseObject{
		Code:           code,
		Message:        msg,
		ResponseObject: results,
	}
	api.Data["json"] = response
	api.Ctx.Output.SetStatus(code)
	api.ServeJSON()
}

// ResponseJSONWithCode sends a JSON response with the given results, HTTP status code, 
// custom int64 code, and message
func (api *API) ResponseJSONWithCode(results interface{}, statusCode int, code int64, msg string) {
	if results == nil {
		results = struct{}{}
	}
	response := &ResponseObjectWithCode{
		Code:           code,
		Message:        msg,
		ResponseObject: results,
	}
	api.Data["json"] = response
	api.Ctx.Output.SetStatus(statusCode)
	api.ServeJSON()
}
