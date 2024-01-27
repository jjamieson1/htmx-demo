package controllers

import (
	content "github.com/jjamieson1/eden-sdk/clients/cms"
	"github.com/jjamieson1/eden-sdk/models"
	"github.com/revel/revel"
)

type Content struct {
	*revel.Controller
}

func (c Content) GetContentByCategoryId(pageName string) revel.Result {
	categoryId := c.Params.Get("categoryId")
	revel.AppLog.Debugf("requesting content from categoryId: %s", categoryId)
	var provider models.TenantProvider
	provider.EdenAdapter.AdapterUrl = "http://localhost:3001/api/v1/content"

	contentList, err := content.GetCmsFromProviderByCategoryId("", "", categoryId, provider)
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
