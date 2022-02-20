package khanmux

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

type Context struct {
	Request  *http.Request
	Handler  Handler
	Response http.ResponseWriter
}

func (c *Context) JSON(status int, data interface{}) error {

	c.Response.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Response.WriteHeader(status)
	err := json.NewEncoder(c.Response).Encode(data)

	if err != nil {
		return err
	}
	return nil
}

func (c *Context) XML(status int, data interface{}) error {

	c.Response.Header().Set("Content-Type", "application/xml; charset=utf-8")
	c.Response.WriteHeader(status)
	x, err := xml.MarshalIndent(data, " ", " ")
	c.Response.Write(x)
	if err != nil {
		return err
	}

	return nil
}

type Response struct {
	StatusCode int
	Data       map[string]interface{}
}

type Handler func(c Context) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := Context{
		Request:  r,
		Handler:  h,
		Response: w,
	}

	h(context)

}
