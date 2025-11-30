package domain

// InventoryService define operaciones de negocio para inventario (por tenant)
type InventoryService interface {
    ListItems(tenantID string) ([]Item, error)
    GetItem(tenantID, id string) (*Item, error)
}
