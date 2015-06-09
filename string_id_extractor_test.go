package db_test

import (
	"testing"
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"reflect"
)

func TestStringIdExtractorExtractsId(t *testing.T) {
	t.Parallel()

	type CreditCard struct {
		db.Model
		Owner string `db:"id"`
	}

	card := &CreditCard{Owner: "Borges"}
	field := reflect.TypeOf(card).Elem().Field(1)
	value := reflect.ValueOf(card).Elem().Field(1)

	meta := &db.Metadata{}
	err := db.StringIdExtractor{meta}.Extract(card, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.StringID).ToBe(card.Owner)
}

func TestStringIdExtractorExtractsReturnsErrMissingStringId(t *testing.T) {
	t.Parallel()

	type CreditCard struct {
		db.Model
		Owner string `db:"id"`
	}

	card := &CreditCard{}
	field := reflect.TypeOf(card).Elem().Field(1)
	value := reflect.ValueOf(card).Elem().Field(1)

	meta := &db.Metadata{}
	err := db.StringIdExtractor{meta}.Extract(card, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrMissingStringId)
	expect(meta.StringID).ToBe("")
}

func TestStringIdExtractorAcceptsTaggedStringField(t *testing.T) {
	t.Parallel()

	type MyModel struct {
		db.Model
		String string `db:"id"`
		NotTaggedString string
		Slice []string
		Map map[string]string
		Rune rune
		Int int
		Int8 int8
		Int16 int16
		Int32 int32
		Int64 int64
	}

	expect := goexpect.New(t)

	extractor := db.StringIdExtractor{}
	elem := reflect.TypeOf(&MyModel{}).Elem()

	expect(extractor.Accept(elem.Field(0))).ToBe(false)
	expect(extractor.Accept(elem.Field(1))).ToBe(true)
	expect(extractor.Accept(elem.Field(2))).ToBe(false)
	expect(extractor.Accept(elem.Field(3))).ToBe(false)
	expect(extractor.Accept(elem.Field(4))).ToBe(false)
	expect(extractor.Accept(elem.Field(5))).ToBe(false)
	expect(extractor.Accept(elem.Field(6))).ToBe(false)
	expect(extractor.Accept(elem.Field(7))).ToBe(false)
	expect(extractor.Accept(elem.Field(8))).ToBe(false)
	expect(extractor.Accept(elem.Field(9))).ToBe(false)
}
