package postgres

import (
	"github.com/Kong/kuma/app/kumactl/pkg/install/data"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"strconv"
	"strings"

	postgres_cfg "github.com/Kong/kuma/pkg/config/plugins/resources/postgres"
	core_plugins "github.com/Kong/kuma/pkg/core/plugins"
	"github.com/Kong/kuma/pkg/plugins/resources/postgres/migrations"
)

func migrateDb(cfg postgres_cfg.PostgresStoreConfig) (uint, error) {
	m, err := newMigrate(cfg)
	if err != nil {
		return 0, err
	}
	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			ver, _, err := m.Version()
			if err != nil {
				return 0, err
			}
			return ver, core_plugins.AlreadyMigrated
		}
		return 0, err
	}
	ver, _, err := m.Version()
	if err != nil {
		return 0, err
	}
	return ver, nil
}

func newMigrate(cfg postgres_cfg.PostgresStoreConfig) (*migrate.Migrate, error) {
	db, err := connectToDb(cfg)
	if err != nil {
		return nil, err
	}
	dbDriver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	sourceDriver, err := httpfs.New(migrations.Migrations, "")
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithInstance("httpfs", sourceDriver, "postgres", dbDriver)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func isDbMigrated(cfg postgres_cfg.PostgresStoreConfig) (bool, error) {
	m, err := newMigrate(cfg)
	if err != nil {
		return false, err
	}
	dbVer, _, err := m.Version()
	if err != nil {
		if err == migrate.ErrNilVersion {
			return false, nil
		}
		return false, err
	}

	fileVer, err := newestMigration()
	if err != nil {
		return false, err
	}

	return dbVer == fileVer, nil
}

func newestMigration() (uint, error) {
	files, err := data.ReadFiles(migrations.Migrations)
	if err != nil {
		return 0, err
	}
	latest := 0
	for _, file := range files {
		parts := strings.Split(file.Name, "_")
		ver, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		if ver > latest {
			latest = ver
		}
	}
	return uint(latest), nil
}