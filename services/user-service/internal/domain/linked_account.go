package domain

type LinkedAccount struct {
	Base
	userID     string
	user       *User
	providerID string
	authToken  string
}

func (la *LinkedAccount) SetUserID(userID string) {
	la.userID = userID
}

func (la *LinkedAccount) GetUserID() string {
	return la.userID
}

func (la *LinkedAccount) SetProviderID(providerID string) {
	la.providerID = providerID
}

func (la *LinkedAccount) GetProviderID() string {
	return la.providerID
}

func (account *LinkedAccount) SetAuthToken(authToken string) {
	account.authToken = authToken
}

func (la *LinkedAccount) GetAuthToken() string {
	return la.authToken
}

func (la *LinkedAccount) SetUser(user *User) {
	la.user = user
}

func (la *LinkedAccount) GetUser() *User {
	return la.user
}

func NewAccount(userID, providerID, authToken, accountID string) *LinkedAccount {
	return &LinkedAccount{
		userID:     userID,
		providerID: providerID,
		authToken:  authToken,
	}
}
