package controllers

import (
	"htmx-demo/app"

	contentClient "github.com/jjamieson1/celestial-sdk/clients/cms"
	"github.com/jjamieson1/celestial-sdk/models"
	"github.com/revel/revel"
)

type Content struct {
	*revel.Controller
}

func (c Content) GetContentByCategoryId(pageName string) revel.Result {
	categoryId := c.Params.Get("categoryId")
	revel.AppLog.Debugf("requesting content from categoryId: %s", categoryId)
	var provider models.TenantProvider
	provider.Adapter.AdapterUrl = "http://localhost:3001/api/v1/cms"

	contentList, err := contentClient.GetCmsByCategoryId(app.Tenant.TenantId, categoryId, provider)
	if err != nil {
		revel.AppLog.Errorf("unable to get content with error: %s", err.Error())
	}
	if len(contentList) == 0 {
		err := "No content found in this category, please select another"
		c.ViewArgs["error"] = err
		return c.RenderTemplate("templates/not-found.html")
	} else {
		c.ViewArgs["contentList"] = contentList
		return c.RenderTemplate("templates/content-list.html")
	}
}
