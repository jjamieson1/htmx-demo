package controllers

import (
	"fmt"

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
	result := ""
	image := ""
	ctn, err := content.GetCmsFromProviderByCategoryId("", "", categoryId, provider)
	if err != nil {
		revel.AppLog.Errorf("unable to get content with error: %s", err.Error())
	}
	for _, a := range ctn {
		if len(a.CmsContent.Images) > 0 {
			for _, img := range a.CmsContent.Images {
				image = img.Filename
			}
		}
		result = fmt.Sprintf(`
			<section class="py-12">
				<div class="container">
					<div class="row align-items-center justify-content-between">
						<div class="col-12 col-md-6 col-lg-5">
						<h2 class="mb-7">%s</h2> 
						%s
						</div>
						<div class="col-12 col-md-6">
							<div>
								<img src="%s" alt="..." class="img-fluid w-100">
							</div>
						</div>
					</div>
				</div>
			</section>
			`, a.CmsContent.Title, a.CmsContent.Body, image)

	}

	return c.RenderHTML(result)
}
