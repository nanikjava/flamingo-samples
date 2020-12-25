package controller

import (
	"context"
	"flamingo.me/flamingo/v3/framework/web"
	"flamingo.me/form/application"
	"nanikjava/flamingo/sessions/form"
)

type (
	// SessionController represents our first simple controller
	SessionController struct {
		responder *web.Responder
	}

	helloViewData struct {
		Name     string
		Nickname string
	}

	simpleViewData struct {
		Name    string
		Address string
	}

	FormController struct {
		formHandlerFactory application.FormHandlerFactory
		responder          *web.Responder
	}
)

func (c *FormController) Inject(resp *web.Responder, f application.FormHandlerFactory) {
	c.formHandlerFactory = f
	c.responder = resp
}

// Inject dependencies
func (controller *SessionController) Inject(responder *web.Responder) *SessionController {
	controller.responder = responder

	return controller
}

// Get is a controller action that renders the `hello.html` template
// Example:
//         curl -v -i -H "nickname: This is my nickname" http://localhost:3322/hello
func (controller *SessionController) Get(_ context.Context, r *web.Request) web.Result {
	r.Session().Store("nickname", r.Request().Header.Get("nickname"))
	data, _ := r.Session().Load("nickname")

	// Calling the Render method from the response helper and render the template "hello"
	return controller.responder.Render("hello", helloViewData{
		Name:     "World",
		Nickname: data.(string),
	})
}

// Post handles POST operation coming from the client.
// No error handling provided !!
// Example:
//         curl --data "name=Nanik" --data "address=MyAddress" http://localhost:3322/postform
func (c *FormController) Post(ctx context.Context, req *web.Request) web.Result {
	builder := c.formHandlerFactory.GetFormHandlerBuilder()

	// this is to instruct Flamingo that we need to use the DataProvider
	// called simple.provider that is configured in module.go.
	builder.SetNamedFormDataProvider("simple.provider")

	formHandler := builder.Build()
	submittedForm, _ := formHandler.HandleForm(ctx, req)

	// This conversion works because we have specified the DataProvider
	// above using the SetNamedFormDataProvider function
	formdata, _ := submittedForm.Data.(form.SimpleFormData)
	return c.responder.Render("simple", simpleViewData{
		Name:    formdata.Name,
		Address: formdata.Address,
	})
}
