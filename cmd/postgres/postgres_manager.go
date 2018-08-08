package postgres

import "github.com/vulcanize/vulcanizedb/pkg/config"

type PostgresManager interface {
	EnsureSchema(databaseConfig config.Database) error
}
