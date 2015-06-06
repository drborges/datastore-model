package db_test

import (
	"appengine/datastore"
	"github.com/drborges/datastore-model"
	"time"
	"appengine"
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
	db.Model        `db:"People"`
	Name     string `db:"id"`
	Country  string
}

type Message struct {
	db.Model        `db:"Messages,hasparent"`
	Content string
}

func NewPerson(name, country string) *Person {
	person := new(Person)
	person.Name = name
	person.Country = country
	return person
}

func NewMessage(content string) *Message {
	message := new(Message)
	message.Content = content
	return message
}

type People []*Person

func (this People) ByCountry(country string) *db.Query {
	return db.QueryFor(new(Person)).Filter("Country=", country)
}

func CreatePeople(d db.Datastore, c appengine.Context, people ...*Person) {
	keys := make([]*datastore.Key, len(people))
	for i, person := range people {
		d.AssignEntityKey(person)
		keys[i] = person.Key()
	}
	datastore.PutMulti(c, keys, people)
	// Gives datastore some time to index the data
	// and make it available for queries
	time.Sleep(2 * time.Second)
}
