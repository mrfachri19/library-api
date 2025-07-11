package config

import (
	"context"
	"fmt"
	"os"
	"strings"
)

func RunMigration() {
	ctx := context.Background()

	// Baca isi file SQL
	content, err := os.ReadFile("migrations/001_init.sql")
	if err != nil {
		fmt.Println("Failed to read migration file:", err)
		return
	}

	// Pisahkan per statement
	queries := strings.Split(string(content), ";")

	for _, query := range queries {
		query = strings.TrimSpace(query)
		if query == "" {
			continue
		}

		_, err := DB.Exec(ctx, query)
		if err != nil {
			fmt.Println("Migration failed:", err)
			return
		}
	}

	fmt.Println("Migration success ✅")
}
