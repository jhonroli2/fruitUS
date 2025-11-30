package httpbo

import (
    "net/http"

    "fruitus/pkg/config"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()
    // Rutas de administración. Algunas son globales para administradores
    mux.Handle("/v1/tenants", http.HandlerFunc(listTenantsHandler))
    mux.Handle("/v1/tenants/", http.HandlerFunc(getTenantHandler))
    // Rutas actuando en nombre de un tenant
    mux.Handle("/v1/tenants-actions", config.TenantMiddleware(http.HandlerFunc(tenantActionHandler)))
    return mux
}
