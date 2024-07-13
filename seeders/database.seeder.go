package seeders

import (
	"context"
	db "goAPI/database/sqlc"
	"log/slog"
)

type DatabaseSeeder struct {
	db *db.Queries
}

func NewDatabaseSeeder(db *db.Queries) *DatabaseSeeder {
	return &DatabaseSeeder{db}
}

func (ds DatabaseSeeder) Start() error {
	ctx := context.Background()

	s := "people"
	isExecuted, err := ds.db.CheckSeedExecuted(ctx, &s)
	if err != nil {
		return err
	}
	if !isExecuted {
		slog.Info("People seeds is not executed, running...")
		peopleSeeder := NewPeopleSeeder(ds.db)

		err := peopleSeeder.Start(20)
		if err != nil {
			return err
		}
	}

	s = "tasks"
	isExecuted, err = ds.db.CheckSeedExecuted(ctx, &s)
	if !isExecuted {
		slog.Info("Tasks seeds is not executed, running...")
		tasksSeeder := NewTasksSeeder(ds.db)

		err = tasksSeeder.Start(100)
		if err != nil {
			return err
		}
	}

	return nil
}
