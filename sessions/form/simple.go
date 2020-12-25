package form

import (
	"context"
	"flamingo.me/flamingo/v3/framework/web"
)

type (
	SimpleFormData struct {
		Name    string `form:"name"  conform:"name"`
		Address string `form:"address"  conform:"name"`
	}

	SimpleFormDataProvider struct{}
)

func (p *SimpleFormDataProvider) GetFormData(ctx context.Context, req *web.Request) (interface{}, error) {
	// define address form data with some default values
	return SimpleFormData{
		Name:    "defaultName",
		Address: "defaultAddress",
	}, nil
}
