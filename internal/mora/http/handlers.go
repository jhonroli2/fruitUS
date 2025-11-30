package httpmora

import (
    "context"
    "encoding/json"
    "net/http"

    "fruitus/pkg/config"
)

type itemResponse struct {
    Tenant string `json:"tenant_id"`
    Msg    string `json:"msg"`
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
    tenant, _ := config.TenantFromContext(r.Context())
    resp := itemResponse{Tenant: tenant, Msg: "mora items (header/query)"}
    _ = json.NewEncoder(w).Encode(resp)
}

func itemsByPathHandler(w http.ResponseWriter, r *http.Request) {
    tenant, _ := config.TenantFromContext(r.Context())
    resp := itemResponse{Tenant: tenant, Msg: "mora items (path)"}
    _ = json.NewEncoder(w).Encode(resp)
}

// helper para posibles pruebas
func withSampleContext() context.Context {
    return config.WithTenantContext(context.Background(), "sample-tenant")
}
