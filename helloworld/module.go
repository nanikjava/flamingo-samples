package helloworld

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	"nanikjava/flamingo/helloworld/interfaces"
)

// Module is our helloWorld Module
type Module struct{}

// Configure is the default Method a Module needs to implement
func (m *Module) Configure(injector *dingo.Injector) {
	// Call Bind helper of router SessionsModule
	// It is a shortcut for: injector.BindMulti((*router.SessionsModule)(nil)).To(new(routes))
	// So what it does is register our routes struct as a router SessionsModule - so that it is "known" to the router module
	web.BindRoutes(injector, new(routes))
}

// routes struct - our internal struct that gets the interface methods for router.SessionsModule
type routes struct {
	// helloController - we will defined routes that are handled by our SessionController - so we need this as a dependency
	helloController *interfaces.HelloController
}

// Inject dependencies - this is called by Dingo and gets an initializes instance of the SessionController passed automatically
func (r *routes) Inject(controller *interfaces.HelloController) *routes {
	r.helloController = controller

	return r
}

// Routes method which defines all routes handlers in module
func (r *routes) Routes(registry *web.RouterRegistry) {
	// Bind the path /hello to a handle with the name "hello"
	registry.MustRoute("/hello", "hello")

	// Bind the controller.Action to the handle "hello":
	registry.HandleGet("hello", r.helloController.Get)

	registry.HandleGet("helloWorld.greetme", r.helloController.GreetMe)
	registry.MustRoute("/greetme", "helloWorld.greetme")
	// Bind a route with a path parameter
	registry.MustRoute("/greetme/:nickname", "helloWorld.greetme")
	// Bind a route with a default value for a param
	registry.MustRoute("/greetflamingo", `helloWorld.greetme(nickname="Flamingo")`)

	registry.HandleData("currenttime", r.helloController.CurrentTime)
}
