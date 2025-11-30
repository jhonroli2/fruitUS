package db

import "context"

// DBClient placeholder: define interfaz mínima para futuras implementaciones
type DBClient interface {
    Ping(ctx context.Context) error
}

// NewConnection placeholder
func NewConnection(dsn string) (DBClient, error) {
    // aquí inicializarías conexión real (sqlx, gorm, etc.)
    return nil, nil
}
