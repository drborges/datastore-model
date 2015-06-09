package db_test

import (
	"testing"
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"reflect"
)

func TestHasParentExtractorExtractsFromTagWithoutKindName(t *testing.T) {
	type Tag struct {
		db.Model    `db:",has_parent"`
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	fieldModel := reflect.TypeOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(fieldModel)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.HasParent).ToBe(true)
}

func TestHasParentExtractorExtractsFromTagWithKindName(t *testing.T) {
	type Tag struct {
		db.Model    `db:"Tags, has_parent"`
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	fieldModel := reflect.TypeOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(fieldModel)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.HasParent).ToBe(true)
}

func TestHasParentExtractorExtractsFromTagWithoutHasParentMetadata(t *testing.T) {
	type Tag struct {
		db.Model    `db:"Tags"`
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	fieldModel := reflect.TypeOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(fieldModel)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.HasParent).ToBe(false)
}

func TestHasParentExtractorExtractsFromTagWithoutTag(t *testing.T) {
	type Tag struct {
		db.Model
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	fieldModel := reflect.TypeOf(tag).Elem().Field(0)

	err := db.HasParentExtractor{meta}.Extract(fieldModel)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.HasParent).ToBe(false)
}
