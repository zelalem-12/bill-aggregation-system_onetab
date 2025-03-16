package model

import (
	"github.com/google/uuid"
	"github.com/zelalem-12/bill-aggregation-system_onetab/provider-service/internal/domain"
)

type Provider struct {
	Base
	Name        string `gorm:"not null"`
	APIEndpoint string `gorm:"not null"`
	AuthMethod  string `gorm:"not null"`
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
	p.APIEndpoint = provider.GetAPIEndpoint()
	p.AuthMethod = provider.GetAuthMethod()

	return nil
}

func (provider *Provider) ToDomainModel() *domain.Provider {
	domainProvider := domain.Provider{}

	domainProvider.SetID(provider.ID.String())
	domainProvider.SetName(provider.Name)
	domainProvider.SetAPIEndpoint(provider.APIEndpoint)
	domainProvider.SetAuthMethod(provider.AuthMethod)

	domainProvider.SetCreatedAt(provider.CreatedAt)
	domainProvider.SetUpdatedAt(provider.UpdatedAt)
	if provider.DeletedAt.Valid {
		domainProvider.SetDeletedAt(&provider.DeletedAt.Time)
	}
	return &domainProvider
}
