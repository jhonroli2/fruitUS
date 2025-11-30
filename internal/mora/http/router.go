package httpmora

import (
    "net/http"

    "fruitus/pkg/config"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()

    // Ruta que utiliza middleware de tenant
    mux.Handle("/v1/items", config.TenantMiddleware(http.HandlerFunc(itemsHandler)))

    // Ruta con tenant por path: /v1/tenants/{tenant_id}/items
    mux.Handle("/v1/tenants/", http.StripPrefix("/v1/tenants", config.TenantMiddleware(http.HandlerFunc(itemsByPathHandler))))

    return mux
}
