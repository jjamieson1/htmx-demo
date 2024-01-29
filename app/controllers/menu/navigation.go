package controllers

import (
	"htmx-demo/app"

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
	provider.EdenAdapter.AdapterUrl = "http://localhost:3001/api/v1/cms"

	categories, err := content.GetCategoryFromProvider(app.Tenant.TenantId, "", provider)
	if err != nil {
		revel.AppLog.Errorf("unable to get content with error: %s", err.Error())
	}

	c.ViewArgs["categories"] = categories
	return c.RenderTemplate("templates/navbar.html")
}
