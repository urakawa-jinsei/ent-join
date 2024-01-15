package main

import (
	"context"
	"database/sql"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/ent-join/ent"
	"github.com/ent-join/ent/contentmoviemetadata"
	"github.com/ent-join/ent/uploadedcontent"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	client := Open("postgresql://postgres:postgres@127.0.0.1/postgres")

	// migrate
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	count, err := client.Debug().UploadedContent.Query().
		Where(uploadedcontent.ID("sample1.mp4")).
		QueryContents().
		QueryContentMovieMetadata().
		Where(
			contentmoviemetadata.And(
				contentmoviemetadata.Width(1920),
				contentmoviemetadata.Height(1080),
			)).
		Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(count)
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
