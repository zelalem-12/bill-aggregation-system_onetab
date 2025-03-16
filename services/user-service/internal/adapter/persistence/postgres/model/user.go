package model

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/onetab/internal/domain"
)

type User struct {
	Base
	FirstName      string `gorm:"not null"`
	LastName       string `gorm:"not null"`
	Email          string `gorm:"not null;unique" `
	Password       string `gorm:"default:''"`
	IsVerified     bool   `gorm:"default:false" `
	ProfilePicture string `gorm:"type:text;default:''"`
	Accounts       []*LinkedAccount
}

func (user *User) FromDomainModel(domainUser *domain.User) error {

	if domainUser.GetID() != "" {
		userID, err := uuid.Parse(domainUser.GetID())
		if err != nil {
			return err
		}
		user.ID = userID
	}

	user.FirstName = domainUser.GetFirstName()
	user.LastName = domainUser.GetLastName()
	user.Email = domainUser.GetEmail()
	user.Password = domainUser.GetPassword()
	user.IsVerified = domainUser.GetIsVerified()
	user.ProfilePicture = domainUser.GetProfilePicture()

	return nil
}

func (user *User) ToDomainModel() *domain.User {
	dominUser := domain.User{}

	dominUser.SetID(user.ID.String())
	dominUser.SetFirstName(user.FirstName)
	dominUser.SetLastName(user.LastName)
	dominUser.SetEmail(user.Email)
	dominUser.SetPassword(user.Password)
	dominUser.SetIsVerified(user.IsVerified)
	dominUser.SetProfilePicture(user.ProfilePicture)

	dominUser.SetCreatedAt(user.CreatedAt)
	dominUser.SetUpdatedAt(user.UpdatedAt)
	if user.DeletedAt.Valid {
		dominUser.SetDeletedAt(&user.DeletedAt.Time)
	}
	return &dominUser

}
