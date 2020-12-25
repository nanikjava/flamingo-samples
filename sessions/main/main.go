package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"flamingo.me/flamingo/v3/core/auth"
	"flamingo.me/flamingo/v3/core/gotemplate"
	"flamingo.me/form"
	"nanikjava/flamingo/sessions/websession"
)

func main() {
	flamingo.App([]dingo.Module{
		new(form.Module),
		new(auth.WebModule), // gotemplate enables the gotemplate template engine module
		new(gotemplate.Module), // gotemplate enables the gotemplate template engine module
		new(websession.SessionsModule),
	})
}
