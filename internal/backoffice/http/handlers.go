package httpbo

import (
    "encoding/json"
    "net/http"

    "fruitus/internal/backoffice/domain"
)

func listTenantsHandler(w http.ResponseWriter, r *http.Request) {
    // Ejemplo: devolvemos un tenant de ejemplo mostrando que incluye mora y coco
    tenants := []domain.Tenant{
        {
            ID:      "tenant-123",
            Name:    "Cafetería Central",
            Owner:   "Juan Perez",
            Contact: "juan@cafecentral.example",
            Services: domain.TenantServices{
                Mora: true,
                Coco: true,
            },
        },
    }
    _ = json.NewEncoder(w).Encode(tenants)
}

func getTenantHandler(w http.ResponseWriter, r *http.Request) {
    // Placeholder: devolver ejemplo mostrando servicios mora/coco
    t := domain.Tenant{
        ID:      "tenant-123",
        Name:    "Cafetería Central",
        Owner:   "Juan Perez",
        Contact: "juan@cafecentral.example",
        Services: domain.TenantServices{
            Mora: true,
            Coco: true,
        },
    }
    _ = json.NewEncoder(w).Encode(t)
}

func tenantActionHandler(w http.ResponseWriter, r *http.Request) {
    _ = json.NewEncoder(w).Encode(struct{ Msg string `json:"msg"` }{Msg: "tenant action placeholder"})
}
