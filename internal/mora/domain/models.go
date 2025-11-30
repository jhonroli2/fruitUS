package domain

// Item representa un recurso del inventario por tenant
type Item struct {
    ID       string
    Name     string
    Quantity int
    TenantID string
}
