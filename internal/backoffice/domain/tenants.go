package domain

// TenantServices describe los módulos/servicios habilitados para un tenant
type TenantServices struct {
    Mora bool `json:"mora"`
    Coco bool `json:"coco"`
}

// Tenant representa un cliente/tenant de FruitUS
// Cada tenant contiene al menos los módulos `mora` (core de negocio)
// y `coco` (integraciones/middleware), además de metadata.
type Tenant struct {
    ID       string         `json:"id"`
    Name     string         `json:"name"`
    Owner    string         `json:"owner"`
    Contact  string         `json:"contact"`
    Services TenantServices `json:"services"`
    // Config puede almacenar configuración específica por tenant (opcional)
    Config map[string]any `json:"config,omitempty"`
}
