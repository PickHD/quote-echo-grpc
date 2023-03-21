package closer

import "github.com/jackc/pgx/v5/pgxpool"

type InfraCloser struct {
	DB *pgxpool.Pool
}

func NewInfraCloser(db *pgxpool.Pool) *InfraCloser {
	return &InfraCloser{DB: db}
}

func (closer *InfraCloser) Close() error {
	closer.DB.Close()

	return nil
}
