package main

import (
	"context"
	"database/sql"
	"log"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/ent-join/ent"
	"github.com/ent-join/ent/content"
	"github.com/ent-join/ent/contentmoviemetadata"
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

	ctx := context.Background()

	// migrate
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	// INSERT
	tx, err := client.Debug().Tx(ctx)
	if err != nil {
		log.Fatal(err)
	}

	contentEntity, err := tx.Content.
		Create().
		SetID("content4_sample1.mp4").
		SetUploadedContentFilename("sample1.mp4").
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	_, err = tx.ContentMovieMetadata.
		Create().
		SetWidth(1920).
		SetHeight(1080).
		SetContent(contentEntity). // ここがポイント 対応するedgeを追加
		Save(ctx)
	if err != nil {
		_ = tx.Rollback()
		log.Fatal(err)
	}

	tx.Commit()

	// INNER JOIN
	count, err := client.Debug().UploadedContent.
		Query().
		Where(uploadedcontent.ID("sample1.mp4")).
		Modify(func(s *entsql.Selector) {
			t := entsql.Table(uploadedcontent.ContentsTable)
			s.Join(t).
				On(
					s.C(uploadedcontent.FieldID),
					t.C(content.FieldUploadedContentFilename),
				)
			u := entsql.Table(content.ContentMovieMetadataTable)
			s.Join(u).
				On(
					t.C(content.FieldID),
					u.C(contentmoviemetadata.FieldID),
				)
			s.Where(
				entsql.And(
					entsql.EQ(u.C(contentmoviemetadata.FieldWidth), 1920),
					entsql.EQ(u.C(contentmoviemetadata.FieldHeight), 1080),
				),
			)
		}).
		Count(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(count)
}
