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

	var provider models.TenantProvider
	provider.EdenAdapter.AdapterUrl = "http://localhost:3001/api/v1/content"
	result := `<div>`
	ctn, err := content.GetCmsFromProviderByCategoryId("", "", "55ff4dbc-ecd6-4b89-be08-7b08849fbb86", provider)
	if err != nil {
		revel.AppLog.Errorf("unable to get content with error: %s", err.Error())

		for _, a := range ctn {
			result = result + a.CmsContent.Body
		}
		result = result + `</div>`
	}
	return c.RenderHTML(result)
}
