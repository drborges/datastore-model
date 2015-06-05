package db_test

import (
	"appengine/datastore"
	"github.com/drborges/datastore-model"
)

type Persons []EntityWithStringID

func (this Persons) ByCountry(country string) *datastore.Query {
	return datastore.NewQuery("Person").Filter("Country=", country)
}

type EntityWithStringID struct {
	db.Model
	StringID string `db:"id"`
}

type EntityWithIntID struct {
	db.Model
	IntID int `db:"id"`
}

type EntityWithNoIDTag struct {
	db.Model
	StringField        string
	AnotherStringField string
}

type EntityWithMultipleIDTags struct {
	db.Model
	IntID    int64  `db:"id"`
	StringID string `db:"id"`
}

type Person struct {
	db.Model `db:"People"`
}
