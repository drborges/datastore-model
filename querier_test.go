package db_test

import (
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"testing"
)

var (
	diego  = NewPerson("Diego", "Brazil")
	bruno  = NewPerson("Bruno", "Brazil")
	munjal = NewPerson("Munjal", "USA")
	people = People{diego, munjal}
)

func NewPerson(name, country string) *Person {
	person := new(Person)
	person.Name = name
	person.Country = country
	return person
}

func TestQuerierEntityAtSlicePtr(t *testing.T) {
	entity := db.EntityAt(&people, 1)

	expect := goexpect.New(t)
	expect(entity).ToBe(munjal)
}

func TestQuerierEntityAtSlice(t *testing.T) {
	entity := db.EntityAt(people, 0)

	expect := goexpect.New(t)
	expect(entity).ToBe(diego)
}

func TestQuerierEntityAtPanicsWhenInvalidParameterIsProvided(t *testing.T) {
	defer func () {
		expect := goexpect.New(t)
		err := recover()
		expect(err).ToBe(db.ErrNotSlice)
	}()

	db.EntityAt(123, 0) // Should panic
	panic("Should not reach here")
}
