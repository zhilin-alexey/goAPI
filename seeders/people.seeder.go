package seeders

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	db "goAPI/database/sqlc"
	"strconv"
)

type PeopleSeeder struct {
	db *db.Queries
}

func NewPeopleSeeder(db *db.Queries) *PeopleSeeder {
	return &PeopleSeeder{db}
}

func (ps *PeopleSeeder) Start(number int) error {
	var fakeData []db.InsertPeopleParams

	ctx := context.Background()

	for i := 0; i < number; i++ {
		lastName := gofakeit.LastName()
		firstName := gofakeit.FirstName()
		middleName := gofakeit.MiddleName()
		address := gofakeit.Address().Address
		fakeData = append(fakeData, db.InsertPeopleParams{
			Surname:        &lastName,
			Name:           &firstName,
			Patronymic:     &middleName,
			Address:        &address,
			PassportSerie:  gofakeit.DigitN(4),
			PassportNumber: gofakeit.DigitN(6),
		})
	}

	affectedNumber, err := ps.db.InsertPeople(ctx, fakeData)
	if affectedNumber != int64(number) {
		return errors.New(
			"after seeding people, wrong number of rows inserted, expected: " +
				string(rune(number)) +
				", got: " +
				strconv.FormatInt(affectedNumber, 10),
		)
	}
	if err != nil {
		return err
	}

	s := "people"
	err = ps.db.MarkSeedRan(ctx, &s)
	if err != nil {
		return err
	}

	return nil
}
