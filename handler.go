package main

import (
	"fmt"
	"strconv"

	"github.com/mix3/go-meiporo-sample/model"
	"github.com/mix3/meiporo"
)

type ResponseData struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func ErrorHandler(c *meiporo.Context, err interface{}) {
	c.Res.WriteJSON(&ResponseData{
		Success: false,
		Message: fmt.Sprintf("%+v", err),
		Result:  nil,
	})
}

func RootHandler(c *meiporo.Context) {
	c.Res.Render("index.tx", nil)
}

func ListHandler(c *meiporo.Context) {
	result, err := model.TodoList()
	if err != nil {
		panic(err)
	}
	c.Res.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}

func CreateHandler(c *meiporo.Context) {
	title := c.Req.FormValue("title")
	if err := model.TodoCreate(title); err != nil {
		panic(err)
	}
	result, err := model.TodoList()
	if err != nil {
		panic(err)
	}
	c.Res.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}

func SwitchHandler(c *meiporo.Context) {
	id, err := strconv.ParseInt(c.Params["id"], 10, 64)
	if err != nil {
		panic(err)
	}
	if err = model.TodoSwitch(id); err != nil {
		panic(err)
	}
	result, err := model.TodoList()
	if err != nil {
		panic(err)
	}
	c.Res.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}

func DeleteHandler(c *meiporo.Context) {
	id, err := strconv.ParseInt(c.Params["id"], 10, 64)
	if err != nil {
		panic(err)
	}
	if err = model.TodoDelete(id); err != nil {
		panic(err)
	}
	result, err := model.TodoList()
	if err != nil {
		panic(err)
	}
	c.Res.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}

func DeleteAllHandler(c *meiporo.Context) {
	if err := model.TodoDeleteAll(); err != nil {
		panic(err)
	}
	result, err := model.TodoList()
	if err != nil {
		panic(err)
	}
	c.Res.WriteJSON(&ResponseData{
		Success: true,
		Message: "",
		Result:  result,
	})
}
