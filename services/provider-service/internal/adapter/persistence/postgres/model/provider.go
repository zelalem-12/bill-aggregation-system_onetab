package model

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/domain"
)

type Provider struct {
	Base
	Name       string `gorm:"not null;unique"`
	AuthMethod string `gorm:"not null;default:'OAuth2'"`
	APIURL     string `gorm:"not null"`

	ClientID     string `gorm:"not null"`
	ClientSecret string `gorm:"not null"`
	TokenURL     string `gorm:"not null"`

	APIToken string `gorm:"not null"`
}

func (p *Provider) FromDomainModel(provider *domain.Provider) error {
	if provider.GetID() != "" {
		providerID, err := uuid.Parse(provider.GetID())
		if err != nil {
			return err
		}
		p.ID = providerID
	}
	p.Name = provider.GetName()
	p.AuthMethod = provider.GetAuthMethod()
	p.ClientID = provider.GetClientID()
	p.ClientSecret = provider.GetClientSecret()
	p.TokenURL = provider.GetTokenURL()
	p.APIURL = provider.GetAPIURL()
	p.APIToken = provider.GetAPIToken()

	return nil
}

func (provider *Provider) ToDomainModel() *domain.Provider {
	domainProvider := domain.Provider{}

	domainProvider.SetID(provider.ID.String())
	domainProvider.SetName(provider.Name)
	domainProvider.SetAuthMethod(provider.AuthMethod)
	domainProvider.SetClientID(provider.ClientID)
	domainProvider.SetClientSecret(provider.ClientSecret)
	domainProvider.SetTokenURL(provider.TokenURL)
	domainProvider.SetAPIURL(provider.APIURL)
	domainProvider.SetAPIToken(provider.APIToken)

	domainProvider.SetCreatedAt(provider.CreatedAt)
	domainProvider.SetUpdatedAt(provider.UpdatedAt)
	if provider.DeletedAt.Valid {
		domainProvider.SetDeletedAt(&provider.DeletedAt.Time)
	}
	return &domainProvider
}
