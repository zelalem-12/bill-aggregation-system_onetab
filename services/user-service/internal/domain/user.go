package domain

type User struct {
	Base
	firstName      string
	lastName       string
	email          string
	password       string
	isVerified     bool
	profilePicture string
	LinkedAccounts []*LinkedAccount
}

func (user *User) SetFirstName(firstName string) {
	user.firstName = firstName
}

func (user *User) GetFirstName() string {
	return user.firstName
}

func (user *User) SetLastName(lastName string) {
	user.lastName = lastName
}

func (user *User) GetLastName() string {
	return user.lastName
}

func (user *User) SetEmail(email string) {
	user.email = email
}

func (user *User) GetEmail() string {
	return user.email
}

func (user *User) SetPassword(password string) {
	user.password = password
}

func (user *User) GetPassword() string {
	return user.password
}

func (user *User) SetIsVerified(isVerified bool) {
	user.isVerified = isVerified
}

func (user *User) GetIsVerified() bool {
	return user.isVerified
}

func (user *User) SetProfilePicture(profilePicture string) {
	user.profilePicture = profilePicture
}

func (user *User) GetProfilePicture() string {
	return user.profilePicture
}

func (user *User) SetLinkedAccounts(linkedAccounts []*LinkedAccount) {
	user.LinkedAccounts = linkedAccounts
}

func (user *User) GetLinkedAccounts() []*LinkedAccount {
	return user.LinkedAccounts
}

func NewUser(firstName, lastName, email string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
	}
}

func (user *User) AddLinkedAccount(linkedAccount *LinkedAccount) {
	user.LinkedAccounts = append(user.LinkedAccounts, linkedAccount)
}
