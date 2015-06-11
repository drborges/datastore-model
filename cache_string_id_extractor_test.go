package db_test

import (
	"testing"
	"github.com/drborges/datastore-model"
	"reflect"
	"github.com/drborges/goexpect"
)


func TestCacheStringIdExtractorExtractsKeyFromTag(t *testing.T) {
	t.Parallel()

	type Tag struct {
		db.Model
		Name string `cache:"id"`
	}

	tag := &Tag{Name: "golang"}
	field := reflect.TypeOf(tag).Elem().Field(1)
	value := reflect.ValueOf(tag).Elem().Field(1)

	meta := &db.Metadata{}
	err := db.CacheStringIdExtractor{meta}.Extract(tag, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.CacheStringID).ToBe("golang")
}

func TestCacheStringIdExtractorExtractsReturnsErrMissingCacheStringID(t *testing.T) {
	t.Parallel()

	type Tag struct {
		db.Model
		Name string `cache:"id"`
	}

	tag := &Tag{}
	field := reflect.TypeOf(tag).Elem().Field(1)
	value := reflect.ValueOf(tag).Elem().Field(1)

	meta := &db.Metadata{}
	err := db.CacheStringIdExtractor{meta}.Extract(tag, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrMissingCacheStringId)
	expect(meta.CacheStringID).ToBe("")
}

func TestCacheStringIdExtractorAcceptsTaggedStringFields(t *testing.T) {
	t.Parallel()

	type Tag struct {
		db.Model
		Name        string `cache:"id"`
		Description string
		Popularity  int    `cache:"id"`
	}

	tag := &Tag{Name: "golang"}
	elem := reflect.TypeOf(tag).Elem()

	meta := &db.Metadata{}

	expect := goexpect.New(t)
	expect(db.CacheStringIdExtractor{meta}.Accept(elem.Field(1))).ToBe(true)
	expect(db.CacheStringIdExtractor{meta}.Accept(elem.Field(2))).ToBe(false)
	expect(db.CacheStringIdExtractor{meta}.Accept(elem.Field(3))).ToBe(false)
}
