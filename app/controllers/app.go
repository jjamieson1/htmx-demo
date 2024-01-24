package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	c.ViewArgs["pageName"] = "home"
	return c.Render()
}
