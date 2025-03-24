package domain

type Provider struct {
	Base
	name       string
	authMethod string
	apiUrl     string

	clientID     string
	clientSecret string
	tokenURL     string

	apiToken string
}

func (provider *Provider) SetName(name string) {
	provider.name = name
}

func (provider *Provider) GetName() string {
	return provider.name
}

func (provider *Provider) SetAuthMethod(authMethod string) {
	provider.authMethod = authMethod
}

func (provider *Provider) GetAuthMethod() string {
	return provider.authMethod
}

func (provider *Provider) SetClientID(clientID string) {
	provider.clientID = clientID
}

func (provider *Provider) GetClientID() string {
	return provider.clientID
}

func (provider *Provider) SetClientSecret(clientSecret string) {
	provider.clientSecret = clientSecret
}

func (provider *Provider) GetClientSecret() string {
	return provider.clientSecret
}

func (provider *Provider) SetTokenURL(tokenURL string) {
	provider.tokenURL = tokenURL
}

func (provider *Provider) GetTokenURL() string {
	return provider.tokenURL
}

func (provider *Provider) SetAPIURL(apiURL string) {
	provider.apiUrl = apiURL
}

func (provider *Provider) GetAPIURL() string {
	return provider.apiUrl
}

func (provider *Provider) SetAPIToken(apiToken string) {
	provider.apiToken = apiToken
}

func (provider *Provider) GetAPIToken() string {
	return provider.apiToken
}

func NewProvider(name, apiURL, auth_method, clientID, clientSecret, tokenURL, apiToken string) *Provider {
	return &Provider{
		name:         name,
		apiUrl:       apiURL,
		authMethod:   auth_method,
		clientID:     clientID,
		clientSecret: clientSecret,
		tokenURL:     tokenURL,
		apiToken:     apiToken,
	}
}
