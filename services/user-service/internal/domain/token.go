package domain

type Token struct {
	Base
	token  string
	userId string
}

func (t *Token) GetToken() string {
	return t.token
}

func (t *Token) SetToken(tokenString string) {
	t.token = tokenString
}

func (t *Token) GetUserID() string {
	return t.userId
}

func (t *Token) SetUserID(userID string) {
	t.userId = userID
}

func NewToken(token string, userId string) *Token {
	return &Token{
		token:  token,
		userId: userId,
	}
}
