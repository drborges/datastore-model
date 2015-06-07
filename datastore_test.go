package db_test

import (
	"appengine/aetest"
	"appengine/datastore"
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"testing"
	"time"
)

type CreditCard struct {
	db.Model
	Number int   `db:"id"`
	Owner  string
}

type CreditCards []*CreditCard

func (this CreditCards) ByOwner(owner string) *db.Query {
	return db.QueryFor(new(CreditCard)).Filter("Owner=", owner)
}

func TestDatastoreCreate(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	createdAt := time.Now()
	clock := func() time.Time { return createdAt }

	card := &CreditCard{Number:123}
	d := db.NewDatastore(c)
	d.Clock = clock
	err := d.Create(card)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(card.CreatedAt.Format("02 Jan 06 15:04 MST")).ToBe(createdAt.Format("02 Jan 06 15:04 MST"))
	expect(card.Key().String()).ToBe("/CreditCard,123")
}

func TestDatastoreCreateDoesNotUpdateCreatedAtIfError(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card := new(CreditCard)
	err := db.NewDatastore(c).Create(card)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrMissingIntId)
	expect(card.CreatedAt).ToBe(time.Time{})
	expect(card.Key()).ToBe((*datastore.Key)(nil))
}

func TestDatastoreCreateAll(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card1 := &CreditCard{Number:1}
	card2 := &CreditCard{Number:2}

	createdAt := time.Now()
	createdAtFormatted := createdAt.Format("02 Jan 06 15:04 MST")
	clock := func() time.Time { return createdAt }

	d := db.NewDatastore(c)
	d.Clock = clock
	err := d.CreateAll(card1, card2)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(card1.CreatedAt.Format("02 Jan 06 15:04 MST")).ToBe(createdAtFormatted)
	expect(card1.Key().String()).ToBe("/CreditCard,1")

	expect(card2.CreatedAt.Format("02 Jan 06 15:04 MST")).ToBe(createdAtFormatted)
	expect(card2.Key().String()).ToBe("/CreditCard,2")
}

func TestDatastoreCreateAllRollsBackAnyChangesToEntitiesWhenReturningError(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card1 := &CreditCard{Number:1}
	card2 := &CreditCard{}

	err := db.NewDatastore(c).CreateAll(card1, card2)

	expect := goexpect.New(t)
	expect(err).ToNotBe(nil)
	expect(card1.CreatedAt).ToBe(time.Time{})
	expect(card1.Key()).ToBe((*datastore.Key)(nil))

	expect(card2.CreatedAt).ToBe(time.Time{})
	expect(card2.Key()).ToBe((*datastore.Key)(nil))
}

func TestDatastoreLoad(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	d := db.NewDatastore(c)
	d.Create(&CreditCard{Number:1, Owner: "Borges"})

	card := CreditCard{Number: 1}
	err := d.Load(&card)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(card.Number).ToBe(1)
	expect(card.Owner).ToBe("Borges")
	expect(card.Key().String()).ToBe("/CreditCard,1")
}

func TestDatastoreUpdate(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card := &CreditCard{Number: 123, Owner: "Borges"}
	err := db.NewDatastore(c).Update(card)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(card.Key().String()).ToBe("/CreditCard,123")
}

// Won't test datastore behavior on deletes
// See: https://groups.google.com/forum/#!topic/google-appengine-go/TIJEFI5gHxQ
func TestDatastoreDelete(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	err := db.NewDatastore(c).Delete(&CreditCard{Number: 123})

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
}

func TestDatastoreDeleteAll(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	err := db.NewDatastore(c).DeleteAll(&CreditCard{Number:1}, &CreditCard{Number:2})

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
}

func TestQueryAllSetKeysToMatchedItems(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card1 := &CreditCard{Number:1, Owner:"Borges"}
	card2 := &CreditCard{Number:2, Owner:"Borges"}
	card3 := &CreditCard{Number:3, Owner:"Diego"}

	d := db.NewDatastore(c)
	d.CreateAll(card1, card2, card3)

	// Gives datastore some time to index the cards before querying
	time.Sleep(1 * time.Second)

	cards := CreditCards{}
	err := d.Query(cards.ByOwner("Borges")).All(&cards)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(len(cards)).ToBe(2)
	expect(cards[0].Key()).ToDeepEqual(card1.Key())
	expect(cards[1].Key()).ToDeepEqual(card2.Key())
}

func TestQueryFirstSetKeysToMatchedItem(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	card1 := &CreditCard{Number:1, Owner:"Borges"}
	card2 := &CreditCard{Number:2, Owner:"Borges"}
	card3 := &CreditCard{Number:3, Owner:"Diego"}

	d := db.NewDatastore(c)
	d.CreateAll(card1, card2, card3)

	// Gives datastore some time to index the cards before querying
	time.Sleep(1 * time.Second)

	card := CreditCard{}
	err := d.Query(CreditCards{}.ByOwner("Borges")).First(&card)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(card.Key()).ToDeepEqual(card1.Key())
}
