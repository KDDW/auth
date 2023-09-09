package migrations

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var Migrations = migrate.NewMigrations()

func Migrate(db *bun.DB) {

	if err := Migrations.DiscoverCaller(); err != nil {
		fmt.Printf("❌ failed to discover caller: %s\n", err.Error())
		return
	}

	if err := InitMigrationTable(db); err != nil {
		fmt.Printf("❌ failed to init migration table: %s\n", err.Error())
		return
	}

	if err := RunMigrations(db); err != nil {
		fmt.Printf("❌ failed to run migrations: %s\n", err.Error())
		return
	}
}

func InitMigrationTable(db *bun.DB) error {
	ctx := context.Background()
	migrator := migrate.NewMigrator(db, Migrations)
	return migrator.Init(ctx)
}

func RunMigrations(db *bun.DB) error {
	ctx := context.Background()
	migrator := migrate.NewMigrator(db, Migrations)

	if err := migrator.Lock(ctx); err != nil {
		return err
	}
	defer migrator.Unlock(ctx)

	group, err := migrator.Migrate(ctx)

	if err != nil {
		return err
	}

	if group.IsZero() {
		fmt.Printf("> ✅ there are no new migrations to run (database is up to date)\n")
		return nil
	}

	fmt.Printf("> ✅ migrated to %s\n", group)
	return nil
}
