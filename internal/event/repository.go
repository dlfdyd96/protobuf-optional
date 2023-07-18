package event

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/pquerna/ffjson/ffjson"
)

type Repository struct {
	db *pg.DB
}

func NewRepository() *Repository {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "qweasd123",
		Database: "sentbe",
	})

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		panic(err)
	}

	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateEvent() error {
	eventDetail1 := map[V3_LOCALE]string{}
	//model := EventModel{}
	eventDetail1[V3_LOCALE_KO] = ""
	model := Event{}

	bytes, err := ffjson.Marshal(&eventDetail1)
	if err != nil {
		return err
	}
	if err := ffjson.Unmarshal(bytes, &model.EventDetail1); err != nil {
		return err
	}

	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Close()

	if _, err := tx.Model(&model).Insert(); err != nil {
		_ = tx.Rollback()
		return err
	}

	// Commit on success.
	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return nil
}
