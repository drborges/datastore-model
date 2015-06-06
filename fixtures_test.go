package db_test

import (
	"time"
	"appengine/datastore"
	"github.com/drborges/datastore-model"
)

var (
	diego  = NewPerson("Diego", "Brazil")
	bruno  = NewPerson("Bruno", "Brazil")
	munjal = NewPerson("Munjal", "USA")
	people = People{diego, munjal}
)

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
	db.Model     `db:"People"`
	Name string  `db:"id"`
	Country string
}

func NewPerson(name, country string) *Person {
	person := new(Person)
	person.Name = name
	person.Country = country
	return person
}

type People []*Person

func (this People) ByCountry(country string) *db.Query {
	return db.QueryFor(new(Person)).Filter("Country=", country)
}

func CreatePeople(d db.Datastore, people ...*Person) {
	keys := make([]*datastore.Key, len(people))
	for i, person := range people {
		d.SetKeyFor(person)
		keys[i] = person.Key()
	}
	datastore.PutMulti(d.Context, keys, people)
	// Gives datastore some time to index the data
	// and make it available for queries
	time.Sleep(2 * time.Second)
}