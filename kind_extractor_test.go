package db_test

import (
	"testing"
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"reflect"
)

func TestKindExtractorExtractsKindFromNonTaggedModel(t *testing.T) {
	type Tag struct {
		db.Model
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	fieldModel := reflect.TypeOf(tag).Elem().Field(0)

	err := db.KindExtractor{meta}.Extract(tag, fieldModel)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.Kind).ToBe("Tag")
}

func TestKindExtractorExtractsKindFromTag(t *testing.T) {
	type Tag struct {
		db.Model   `db:"Tags"`
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	fieldModel := reflect.TypeOf(tag).Elem().Field(0)

	err := db.KindExtractor{meta}.Extract(tag, fieldModel)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.Kind).ToBe("Tags")
}

func TestKindExtractorAccpetsModelEmbeddedField(t *testing.T) {
	type Tag struct {
		db.Model
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	fieldModel := reflect.TypeOf(tag).Elem().Field(0)

	accepts := db.KindExtractor{meta}.Accept(fieldModel)

	expect := goexpect.New(t)
	expect(accepts).ToBe(true)
}

func TestKindExtractorDoesNotAccpetNonModelEmbeddedField(t *testing.T) {
	type Tag struct {
		db.Model
		Name string
	}

	tag := &Tag{}
	meta := &db.Metadata{}
	fieldModel := reflect.TypeOf(tag).Elem().Field(1)

	accepts := db.KindExtractor{meta}.Accept(fieldModel)

	expect := goexpect.New(t)
	expect(accepts).ToBe(false)
}

