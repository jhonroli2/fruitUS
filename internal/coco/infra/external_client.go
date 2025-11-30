package infra

import "context"

// ExternalClient representa un cliente para sistemas externos (acepta tenant)
type ExternalClient interface {
    SendEvent(ctx context.Context, tenantID string, payload any) error
}
