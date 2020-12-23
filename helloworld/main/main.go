package main

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3"
	"nanikjava/flamingo/helloworld/mycmd"
)

// main is our entry point
func main() {
	flamingo.App([]dingo.Module{
		new(mycmd.Module),    // gotemplate enables the gotemplate template engine module
	})
}
