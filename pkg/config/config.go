package config

import (
    "context"
    "net/http"
    "strings"
)

type ctxKey string

const tenantKey ctxKey = "tenant-id"
const TenantHeader = "X-Tenant-ID"
const TenantQueryParam = "tenant_id"

// LoadConfig placeholder
func LoadConfig() (map[string]string, error) {
    return map[string]string{"env": "dev"}, nil
}

// ExtractTenantFromRequest intenta extraer tenant desde header, query o path (/tenants/{id}/...)
func ExtractTenantFromRequest(r *http.Request) (string, bool) {
    if t := r.Header.Get(TenantHeader); t != "" {
        return t, true
    }
    if t := r.URL.Query().Get(TenantQueryParam); t != "" {
        return t, true
    }
    parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
    for i := 0; i < len(parts)-1; i++ {
        if parts[i] == "tenants" {
            return parts[i+1], true
        }
    }
    return "", false
}

// TenantMiddleware inyecta tenant en contexto si existe
func TenantMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if t, ok := ExtractTenantFromRequest(r); ok {
            ctx := WithTenantContext(r.Context(), t)
            next.ServeHTTP(w, r.WithContext(ctx))
            return
        }
        next.ServeHTTP(w, r)
    })
}

func WithTenantContext(ctx context.Context, tenant string) context.Context {
    return context.WithValue(ctx, tenantKey, tenant)
}

func TenantFromContext(ctx context.Context) (string, bool) {
    t, ok := ctx.Value(tenantKey).(string)
    return t, ok
}
