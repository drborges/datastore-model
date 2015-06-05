package db_test

import (
	"github.com/drborges/datastore-model"
	"time"
)

type EntityWithStringID struct {
	db.Entity
	StringID string `db:"id"`
}

type EntityWithIntID struct {
	db.Entity
	IntID int `db:"id"`
}

type EntityWithNoIDTag struct {
	db.Entity
	StringField        string
	AnotherStringField string
}

type EntityWithMultipleIDTags struct {
	db.Entity
	IntID    int64  `db:"id"`
	StringID string `db:"id"`
}

type Person struct {
	db.Entity     `db:"People"`
	Name, Country string
}

type People []*Person

func (this People) ByCountry(country string) *db.Query {
	return db.QueryFor(new(Person)).Filter("Country=", country)
}

func CreatePeople(d db.Datastore, people ...*Person) {
	for _, person := range people {
		d.Create(person)
	}
	// Gives datastore some time to index the data
	// and make it available for queries
	time.Sleep(1 * time.Second)
}