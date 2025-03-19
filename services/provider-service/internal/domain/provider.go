package domain

type Provider struct {
	Base
	Name         string
	AuthMethod   string
	ClientID     string
	ClientSecret string
	TokenURL     string
	APIURL       string
}

func (provider *Provider) SetName(name string) {
	provider.Name = name
}

func (provider *Provider) GetName() string {
	return provider.Name
}

func (provider *Provider) SetAuthMethod(authMethod string) {
	provider.AuthMethod = authMethod
}

func (provider *Provider) GetAuthMethod() string {
	return provider.AuthMethod
}

func (provider *Provider) SetClientID(clientID string) {
	provider.ClientID = clientID
}

func (provider *Provider) GetClientID() string {
	return provider.ClientID
}

func (provider *Provider) SetClientSecret(clientSecret string) {
	provider.ClientSecret = clientSecret
}

func (provider *Provider) GetClientSecret() string {
	return provider.ClientSecret
}

func (provider *Provider) SetTokenURL(tokenURL string) {
	provider.TokenURL = tokenURL
}

func (provider *Provider) GetTokenURL() string {
	return provider.TokenURL
}

func (provider *Provider) SetAPIBaseURL(apiBaseURL string) {
	provider.APIURL = apiBaseURL
}

func (provider *Provider) GetAPIBaseURL() string {
	return provider.APIURL
}

func NewProvider(name, clientID, clientSecret, tokenURL, apiURL string) *Provider {
	return &Provider{
		Name:         name,
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     tokenURL,
		APIURL:       apiURL,
	}
}
