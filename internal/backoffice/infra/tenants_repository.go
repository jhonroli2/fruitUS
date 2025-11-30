package infra

import "fruitus/internal/backoffice/domain"

// TenantsRepository define las operaciones de persistencia para tenants
type TenantsRepository interface {
    CreateTenant(t *domain.Tenant) error
    GetTenant(id string) (*domain.Tenant, error)
    ListTenants() ([]domain.Tenant, error)
}
