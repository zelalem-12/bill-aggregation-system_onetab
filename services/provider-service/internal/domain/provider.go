package domain

type Provider struct {
	Base
	name        string
	apiEndpoint string
	authMethod  string
}

func (provider *Provider) SetName(name string) {
	provider.name = name
}

func (provider *Provider) GetName() string {
	return provider.name
}

func (provider *Provider) SetAPIEndpoint(apiEndpoint string) {
	provider.apiEndpoint = apiEndpoint
}

func (provider *Provider) GetAPIEndpoint() string {
	return provider.apiEndpoint
}

func (provider *Provider) SetAuthMethod(authMethod string) {
	provider.authMethod = authMethod
}

func (provider *Provider) GetAuthMethod() string {
	return provider.authMethod
}

func NewProvider(name, api, authMethod string) *Provider {
	return &Provider{
		name:        name,
		apiEndpoint: api,
		authMethod:  authMethod,
	}
}
