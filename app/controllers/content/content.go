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
	result := `<div>`
	ctn, err := content.GetCmsFromProviderByCategoryId("", "", categoryId, provider)
	if err != nil {
		revel.AppLog.Errorf("unable to get content with error: %s", err.Error())
	}
	for _, a := range ctn {
		result = result + `<div class="m-5">` + a.CmsContent.Body
		if len(a.CmsContent.Images) > 0 {
			for _, img := range a.CmsContent.Images {
				result = result + `<div class="float-right"><img src="` + img.Filename + `/></div>`
			}
		}
		result = result + `</div>`
	}
	result = result + `</div>`

	return c.RenderHTML(result)
}
