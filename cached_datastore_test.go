package db_test

import (
	"testing"
	"github.com/drborges/datastore-model"
	"appengine/aetest"
	"github.com/drborges/goexpect"
	"appengine/datastore"
)

type MembershipCard struct {
	db.Model
	Number int   `db:"id"`
	Owner  string
}

func TestLoadsModelFromCache(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card := &MembershipCard{Number: 123, Owner: "Borges"}

	cds := db.CachedDatastore{db.NewDatastore(c)}
	cds.Create(card)

	card.Owner = "Diego"
	cds.Datastore.Update(card)

	cardFromDB := &MembershipCard{Number: 123}
	cds.Datastore.Load(cardFromDB)

	cardFromCache := &MembershipCard{Number: 123}
	err := cds.Load(cardFromCache)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(cardFromDB.Owner).ToBe("Diego")
	expect(cardFromCache.Owner).ToBe("Borges")
}

func TestUpdatesModelInCacheAndDatastore(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card := &MembershipCard{Number: 123, Owner: "Borges"}

	cds := db.CachedDatastore{db.NewDatastore(c)}
	cds.Create(card)

	card.Owner = "Diego"
	err := cds.Update(card)

	cardFromDB := &MembershipCard{Number: 123}
	cds.Datastore.Load(cardFromDB)

	cardFromCache := &MembershipCard{Number: 123}
	cds.Load(cardFromCache)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(cardFromDB.Owner).ToBe("Diego")
	expect(cardFromCache.Owner).ToBe("Diego")
}

func TestDeletesModelFromCacheAndDatastore(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card := &MembershipCard{Number: 123, Owner: "Borges"}

	cds := db.CachedDatastore{db.NewDatastore(c)}
	cds.Create(card)

	err := cds.Delete(card)

	cardFromCache := &MembershipCard{Number: 123}
	errNoSuchEntity := cds.Load(cardFromCache)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(errNoSuchEntity).ToBe(datastore.ErrNoSuchEntity)
}
