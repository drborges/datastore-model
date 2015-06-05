package db_test

import (
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"testing"
)

type Tags []*EntityWithStringID

var (
	golang = NewEntity("golang")
	appengine = NewEntity("appengine")
	tags = Tags{golang, appengine}
)

func NewEntity(id string) *EntityWithStringID {
	entity := new(EntityWithStringID)
	entity.StringID = id
	return entity
}

func TestQuerierEntityAtSlicePtr(t *testing.T) {
	entity := db.EntityAt(&tags, 1)

	expect := goexpect.New(t)
	expect(entity).ToBe(appengine)
}

func TestQuerierEntityAtSlice(t *testing.T) {
	entity := db.EntityAt(tags, 0)

	expect := goexpect.New(t)
	expect(entity).ToBe(golang)
}
