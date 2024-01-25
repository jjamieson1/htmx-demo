package controllers

import (
	"fmt"

	content "github.com/jjamieson1/eden-sdk/clients/cms"
	"github.com/jjamieson1/eden-sdk/models"
	"github.com/revel/revel"
)

type Menu struct {
	*revel.Controller
}

func (c Menu) GetMenu(pageName string) revel.Result {

	// Change this to use the Eden Tenant service to get configuration information
	var provider models.TenantProvider
	provider.EdenAdapter.AdapterUrl = "http://localhost:3001/api/v1/content"

	contents, err := content.GetCategoryFromProvider("", "", provider)
	if err != nil {
		revel.AppLog.Errorf("unable to get content with error: %s", err.Error())
	}
	ctg := `<select name="categoryId" hx-target="#main" hx-post="/components/content" class="form-select"><option>Select Page</option>`
	for _, cat := range contents {
		ctg = ctg + `<option value="`
		ctg = ctg + cat.CmsCategoryId + `">` + cat.Name + `</option>`
	}
	ctg = ctg + `</select>`
	result := fmt.Sprintf(`
			<nav class="navbar navbar-expand-lg navbar-light">
			<div class="container-fluid">
				<a class="navbar-brand" href="/">Celestial Technologies<span style="color:red">.</span></a>
				<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarScroll" aria-controls="navbarScroll" aria-expanded="false" aria-label="Toggle navigation">
					<span class="navbar-toggler-icon"></span>
				</button>
				<div class="collapse navbar-collapse" id="navbarScroll">
					<ul class="navbar-nav ms-auto mb-2 mb-lg-0">
						<li class="nav-item">
							<a class="nav-link" style="margin: 5px" href="/">Home</a>
						</li>
						<li class="nav-item">
						%s
						</li>
					</ul>
				</div>
			</div>
		</nav>
		`, ctg)
	return c.RenderHTML(result)
}
