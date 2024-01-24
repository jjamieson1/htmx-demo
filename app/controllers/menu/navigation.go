package controllers

import (
	"github.com/revel/revel"
)

type Menu struct {
	*revel.Controller
}

func (c Menu) GetMenu(pageName string) revel.Result {
	revel.AppLog.Info(pageName)
	result := `
	<nav class="navbar navbar-expand-lg navbar-light">
	<div class="container-fluid">
		<a class="navbar-brand" href="/">Celestial Technologies<span style="color:red">.</span></a>
		<button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarScroll" aria-controls="navbarScroll" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
		<div class="collapse navbar-collapse" id="navbarScroll">
			<ul class="navbar-nav ms-auto mb-2 mb-lg-0">
				<li class="nav-item">
					<a class="nav-link" aria-current="Home" href="/"`
	if pageName == "home" {
		result = result +
			`
		style="color:red;"
		`
	}

	result = result +
		`
	>Home</a>
	</li>
	<li class="nav-item active">
		<a class="nav-link" aria-current="Home" href="/pages/blog"
	`

	if pageName == "blog" {
		result = result + `
	style="color:red;"
	`
	}

	result = result + `
					>Blog</a>
				</li>
			</ul>
		</div>
	</div>
</nav>
`
	return c.RenderHTML(result)
}
