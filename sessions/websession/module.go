package websession

import (
	"flamingo.me/dingo"
	"flamingo.me/flamingo/v3/framework/web"
	formDomain "flamingo.me/form/domain"
	"nanikjava/flamingo/sessions/form"

	"nanikjava/flamingo/sessions/controller"
)

// SessionsModule is our helloWorld SessionsModule
type SessionsModule struct{}

// Configure is the default Method a SessionsModule needs to implement
func (m *SessionsModule) Configure(injector *dingo.Injector) {
	// Call Bind helper of router SessionsModule
	// It is a shortcut for: injector.BindMulti((*router.SessionsModule)(nil)).To(new(routes))
	// So what it does is register our routes struct as a router SessionsModule - so that it is "known" to the router module
	web.BindRoutes(injector, new(routes))

	injector.BindMap(new(formDomain.FormDataProvider), "simple.provider").To(form.SimpleFormDataProvider{})

}

// routes struct - our internal struct that gets the interface methods for router.SessionsModule
type routes struct {
	// sessionController - we will defined routes that are handled by our HelloController - so we need this as a dependency
	sessionController *controller.SessionController
	formController    *controller.FormController
}

// Inject dependencies - this is called by Dingo and gets an initializes instance of the HelloController passed automatically
func (r *routes) Inject(controller *controller.SessionController, formcontroller *controller.FormController) *routes {
	r.sessionController = controller
	r.formController = formcontroller

	return r
}

// Routes method which defines all routes handlers in module
func (r *routes) Routes(registry *web.RouterRegistry) {
	// Bind the path /hello to a handle with the name "hello"
	registry.MustRoute("/hello", "hello")
	registry.HandleGet("hello", r.sessionController.Get)

	registry.MustRoute("/postform", "postform")
	registry.HandlePost("postform", r.formController.Post)

}
