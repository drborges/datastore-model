package db_test

import (
	"appengine/aetest"
	"appengine/datastore"
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"reflect"
	"testing"
)

func TestHasParentExtractorExtractsFromTagWithoutKindName(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	type Tag struct {
		db.Model `db:",has_parent"`
		Name     string
	}

	tag := &Tag{}
	tag.SetParent(datastore.NewIncompleteKey(c, "Kind", nil))

	meta := &db.Metadata{}
	field := reflect.TypeOf(tag).Elem().Field(0)
	value := reflect.ValueOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(tag, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.HasParent).ToBe(true)
	expect(meta.Parent).ToBe(tag.Parent())
}

func TestHasParentExtractorExtractsFromTagWithKindName(t *testing.T) {
	t.Parallel()
	c, _ := aetest.NewContext(nil)
	defer c.Close()

	type Tag struct {
		db.Model `db:"Tags, has_parent"`
		Name     string
	}

	tag := &Tag{}
	tag.SetParent(datastore.NewIncompleteKey(c, "Kind", nil))

	meta := &db.Metadata{}
	field := reflect.TypeOf(tag).Elem().Field(0)
	value := reflect.ValueOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(tag, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.HasParent).ToBe(true)
	expect(meta.Parent).ToBe(tag.Parent())
}

func TestHasParentExtractorExtractsFromTagWithoutHasParentMetadata(t *testing.T) {
	t.Parallel()
	type Tag struct {
		db.Model `db:"Tags"`
		Name     string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	field := reflect.TypeOf(tag).Elem().Field(0)
	value := reflect.ValueOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(tag, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.HasParent).ToBe(false)
}

func TestHasParentExtractorExtractsFromTagWithoutTag(t *testing.T) {
	t.Parallel()
	type Tag struct {
		db.Model
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	field := reflect.TypeOf(tag).Elem().Field(0)
	value := reflect.ValueOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(tag, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.HasParent).ToBe(false)
}

func TestHasParentExtractorExtractsReturnsErrMissingParentKey(t *testing.T) {
	t.Parallel()
	type Tag struct {
		db.Model `db:",has_parent"`
		Name     string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	field := reflect.TypeOf(tag).Elem().Field(0)
	value := reflect.ValueOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(tag, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrMissingParentKey)
	expect(meta.HasParent).ToBe(true)
	expect(meta.Parent).ToBe((*datastore.Key)(nil))
}

func TestHasParentExtractorOnlyAcceptsFieldTypeModel(t *testing.T) {
	t.Parallel()
	type Tag struct {
		db.Model
		String string
		Int    int
		Rune   rune
		Slice  []string
		Map    map[string]string
	}

	tag := &Tag{}
	elem := reflect.TypeOf(tag).Elem()

	expect := goexpect.New(t)
	expect(db.HasParentExtractor{}.Accept(elem.Field(0))).ToBe(true)
	expect(db.HasParentExtractor{}.Accept(elem.Field(1))).ToBe(false)
	expect(db.HasParentExtractor{}.Accept(elem.Field(2))).ToBe(false)
	expect(db.HasParentExtractor{}.Accept(elem.Field(3))).ToBe(false)
	expect(db.HasParentExtractor{}.Accept(elem.Field(4))).ToBe(false)
}
