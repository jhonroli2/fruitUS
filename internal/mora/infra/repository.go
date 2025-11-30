package infra

import "fruitus/internal/mora/domain"

// ItemRepository es la interfaz de persistencia para items (por tenant)
type ItemRepository interface {
    FindAll(tenantID string) ([]domain.Item, error)
    FindByID(tenantID, id string) (*domain.Item, error)
    Save(tenantID string, it *domain.Item) error
}
