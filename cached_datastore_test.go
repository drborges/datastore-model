package db_test

import (
	"appengine/aetest"
	"appengine/datastore"
	"appengine/memcache"
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"testing"
)

type MembershipCard struct {
	db.Model
	Number int `db:"id"`
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

func TestCachedDatastoreUsesTaggedFieldAsCacheKey(t *testing.T) {
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	type MembershipCard struct {
		db.Model
		Number int `db:"id"`
		Owner  string `cache:"id"`
	}

	card := &MembershipCard{Number: 123, Owner: "Diego"}

	cds := db.CachedDatastore{db.NewDatastore(c)}
	cds.Create(card)

	cachedCard := &MembershipCard{}
	_, err := memcache.JSON.Get(c, "Diego", cachedCard)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(cachedCard.Number).ToBe(card.Number)
	expect(cachedCard.Owner).ToBe(card.Owner)
}
