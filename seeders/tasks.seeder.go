package seeders

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	db "goAPI/database/sqlc"
	"strconv"
	"time"
)

type TasksSeeder struct {
	db *db.Queries
}

func NewTasksSeeder(db *db.Queries) *TasksSeeder {
	return &TasksSeeder{db}
}

func (ts *TasksSeeder) Start(number int) error {
	var fakeData []db.InsertTasksParams

	ctx := context.Background()

	for i := 0; i < number; i++ {
		randomPeopleId, err := ts.db.RandomPeopleId(ctx)
		if err != nil {
			return err
		}
		minDate := time.Date(1960, time.January, 1, 0, 0, 0, 0, time.UTC)

		endTime := gofakeit.DateRange(
			minDate,
			time.Now(),
		)
		fakeData = append(fakeData, db.InsertTasksParams{
			PeopleID: randomPeopleId,
			Name:     gofakeit.VerbAction(),
			StartTime: gofakeit.DateRange(
				minDate,
				endTime,
			),
			EndTime: &endTime,
		})
	}

	affectedNumber, err := ts.db.InsertTasks(ctx, fakeData)
	if affectedNumber != int64(number) {
		return errors.New(
			"after seeding tasks, wrong number of rows inserted, expected: " +
				string(rune(number)) +
				", got: " +
				strconv.FormatInt(affectedNumber, 10),
		)
	}
	if err != nil {
		return err
	}

	s := "tasks"
	err = ts.db.MarkSeedRan(ctx, &s)
	if err != nil {
		return err
	}

	return nil
}
