package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	migrationFolderPath := "internals/infra/db/migrations"

	// ask for description in the terminal
	description := ""

	fmt.Printf("> ######## provide a description for the new migration ######## \n> this will be in the name of the migration filename: ")
	fmt.Scanf("%s", &description)

	if description == "" {
		panic("description is required")
	}

	now := time.Now()

	timestamp := now.Format("20060102150405")

	upMigration := fmt.Sprintf("%s_%s.up.sql", timestamp, description)
	downMigration := fmt.Sprintf("%s_%s.down.sql", timestamp, description)

	upPath := fmt.Sprintf("%s/%s", migrationFolderPath, upMigration)
	downPath := fmt.Sprintf("%s/%s", migrationFolderPath, downMigration)

	upFile, err := os.Create(upPath)
	if err != nil {
		panic(err)
	}
	defer upFile.Close()
	downFile, err := os.Create(downPath)
	if err != nil {
		panic(err)
	}
	defer downFile.Close()
	fmt.Printf("✅ up migration file created: %s\n", upPath)
	fmt.Printf("✅ down migration file created: %s\n", downPath)

}
