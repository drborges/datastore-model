package db_test

import (
	"testing"
	"github.com/drborges/datastore-model"
	"appengine/aetest"
	"appengine"
	"appengine/datastore"
	"github.com/drborges/goexpect"
)

type Person struct {
	db.Model
	Name    string
	Country string
}

func (this *Person) Kind() string {
	return "Persons"
}

func (this *Person) NewKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, this.Kind(), this.Name, 0, nil)
}

func TestDatastoreCreate(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	d := db.Datastore{c}

	person := new(Person)
	person.Name = "Diego"
	person.Country = "Brazil"
	err := d.Create(person)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(person.Key().String()).ToBe("/Persons,Diego")
}

func TestDatastoreCreateReturnsErrEntityExists(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	d := db.Datastore{c}

	person := new(Person)
	person.Name = "Diego"
	person.Country = "Brazil"
	err := d.Create(person)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)

	err = d.Create(person)
	expect(err).ToBe(db.ErrEntityExists)
}

func TestDatastoreUpdate(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	d := db.Datastore{c}

	person := new(Person)
	person.Name = "Diego"
	person.Country = "Brazil"
	d.Create(person)

	person.Country = "US"
	err := d.Update(person)

	personFromDB := new(Person)
	personFromDB.SetKey(person.Key())
	d.Load(personFromDB)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(person.Country).ToBe(personFromDB.Country)
}

func TestDatastoreUpdateReturnsErrNoSuchEntity(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	d := db.Datastore{c}

	person := new(Person)
	person.Name = "Diego"
	person.Country = "Brazil"
	err := d.Update(person)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrNoSuchEntity)
}


func TestDatastoreDelete(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	d := db.Datastore{c}

	person := new(Person)
	person.Name = "Diego"
	person.Country = "Brazil"
	d.Create(person)
	err := d.Delete(person)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)

	personFromDB := new(Person)
	personFromDB.SetKey(person.Key())
	err = d.Load(personFromDB)
	expect(err).ToBe(db.ErrNoSuchEntity)
}

func TestDatastoreDeleteReturnsErrNoSuchEntity(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	d := db.Datastore{c}

	person := new(Person)
	person.Name = "Diego"
	person.Country = "Brazil"
	err := d.Delete(person)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrNoSuchEntity)
}
