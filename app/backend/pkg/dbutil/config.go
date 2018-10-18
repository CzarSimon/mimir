package db

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gobuffalo/packr"
	migrate "github.com/rubenv/sql-migrate"
)

type Config struct {
	Host     string `env:"HOST"`
	Port     string `env:"PORT" envDefault:"5432"`
	Database string `env:"DATABASE"`
	Username string `env:"USERNAME"`
	Password string `env:"PASSWORD"`
	SSLMode  bool   `env:"SSL_MODE" envDefault:"false"`
}

func (c Config) ConnectPostgres() (*sql.DB, error) {
	db, err := sql.Open("postgres", c.PgDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	return db, err
}

func (c Config) PgDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%t",
		c.Host, c.Username, c.Password, c.Database, c.Port, c.SSLMode)
}

func Migrate(migrationsPath, driverName string, db *sql.DB) error {
	migrationSource := &migrate.PackrMigrationSource{
		Box: packr.NewBox(migrationsPath),
	}
	migrate.SetTable("schema_version")

	migrations, err := migrationSource.FindMigrations()
	if err != nil {
		return err
	}

	if len(migrations) == 0 {
		return errors.New("Missing database migrations")
	}

	_, err = migrate.Exec(db, driverName, migrationSource, migrate.Up)
	if err != nil {
		return fmt.Errorf("Error applying database migrations: %s", err)
	}
	return nil
}
