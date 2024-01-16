package main

import (
	"context"
	"database/sql"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/ent-join/ent"
	"github.com/ent-join/ent/content"
	"github.com/ent-join/ent/uploadedcontent"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	Query()
}

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func Query() {
	client := Open("postgresql://postgres:postgres@127.0.0.1/postgres")
	defer client.Close()

	// migrate
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	// INNER JOIN
	count, err := client.UploadedContent.Query().
		Where(uploadedcontent.ID("sample1.mp4")).
		Modify(func(s *entsql.Selector) {
			t := entsql.Table(uploadedcontent.ContentsTable)
			s.Join(t).
				On(
					s.C(uploadedcontent.FieldID),
					t.C(content.FieldUploadedContentFilename),
				)
		}).
		Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(count)
}
