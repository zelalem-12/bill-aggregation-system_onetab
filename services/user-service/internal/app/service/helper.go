package service

type EmailContentData struct {
	Title          string
	Header         string
	FirstName      string
	LastName       string
	Link           string
	DepartmentName string
	CompanyName    string
}

func NewEmailContentData(title, header, firstName, lastName, verificationLink string) *EmailContentData {
	return &EmailContentData{
		Title:          title,
		Header:         header,
		FirstName:      firstName,
		LastName:       lastName,
		Link:           verificationLink,
		DepartmentName: "OneTab Support",
		CompanyName:    "OneTab Inc.",
	}
}
