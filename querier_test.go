package db_test

import (
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"testing"
)

func TestQuerierEntityAtSlicePtr(t *testing.T) {
	t.Parallel()
	entity := db.EntityAt(&people, 1)

	expect := goexpect.New(t)
	expect(entity).ToBe(munjal)
}

func TestQuerierEntityAtSlice(t *testing.T) {
	t.Parallel()
	entity := db.EntityAt(people, 0)

	expect := goexpect.New(t)
	expect(entity).ToBe(diego)
}

func TestQuerierEntityAtPanicsWhenInvalidParameterIsProvided(t *testing.T) {
	t.Parallel()
	defer func() {
		expect := goexpect.New(t)
		err := recover()
		expect(err).ToBe(db.ErrInvalidType)
	}()

	db.EntityAt(123, 0) // Should panic
	panic("Should not reach here")
}
