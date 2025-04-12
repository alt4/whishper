package db

import (
	"context"
	"log"
	"strings"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/pluja/anysub/ent"
	"github.com/pluja/anysub/utils"
)

var client *ent.Client

// Init initializes the database client.
func Init() {
	var err error

	// Get the full DSN and type from env
	dsn := utils.Getenv("DATABASE_URI", "anysub:anysub@tcp(database:3306)/anysub?parseTime=True")
	driver := detectDriver(dsn)

	// Open the appropriate dialect
	client, err = ent.Open(driver, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", driver, err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// detectDriver attempts to detect the DB dialect from the DSN
func detectDriver(dsn string) string {
	switch {
	case strings.HasPrefix(dsn, "postgres://") || strings.HasPrefix(dsn, "postgresql://"):
		return dialect.Postgres
	case strings.Contains(dsn, "@tcp("):
		return dialect.MySQL
	default:
		log.Fatalf("unsupported or undetected database dialect in DSN: %s", dsn)
		return ""
	}
}

func Client() *ent.Client {
	return client
}
