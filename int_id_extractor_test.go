package db_test

import (
	"testing"
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"reflect"
)

func TestIntIdExtractorExtractsId(t *testing.T) {
	t.Parallel()

	type CreditCard struct {
		db.Model
		Number int `db:"id"`
	}

	card := &CreditCard{Number:123}
	field := reflect.TypeOf(card).Elem().Field(1)
	value := reflect.ValueOf(card).Elem().Field(1)

	meta := &db.Metadata{}
	err := db.IntIdExtractor{meta}.Extract(card, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(meta.IntID).ToBe(int64(123))
}

func TestIntIdExtractorExtractsReturnsErrMissingIntId(t *testing.T) {
	t.Parallel()

	type CreditCard struct {
		db.Model
		Number int `db:"id"`
	}

	card := &CreditCard{}
	field := reflect.TypeOf(card).Elem().Field(1)
	value := reflect.ValueOf(card).Elem().Field(1)

	meta := &db.Metadata{}
	err := db.IntIdExtractor{meta}.Extract(card, field, value)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrMissingIntId)
	expect(meta.IntID).ToBe(int64(0))
}

func TestIntIdExtractorAcceptsTaggedIntFields(t *testing.T) {
	t.Parallel()

	type MyModel struct {
		db.Model
		String string
		Byte byte
		Map map[string]string
		Slice []string
		Struct struct {
			Int int
		}
		Int int     `db:"id"`
		Int8 int8   `db:"id"`
		Int16 int16 `db:"id"`
		Int32 int32 `db:"id"`
		Int64 int64 `db:"id"`
		NotTaggedInt64 int64
	}

	expect := goexpect.New(t)

	extractor := db.IntIdExtractor{}
	elem := reflect.TypeOf(&MyModel{}).Elem()

	expect(extractor.Accept(elem.Field(0))).ToBe(false)
	expect(extractor.Accept(elem.Field(1))).ToBe(false)
	expect(extractor.Accept(elem.Field(2))).ToBe(false)
	expect(extractor.Accept(elem.Field(3))).ToBe(false)
	expect(extractor.Accept(elem.Field(4))).ToBe(false)
	expect(extractor.Accept(elem.Field(5))).ToBe(false)
	expect(extractor.Accept(elem.Field(6))).ToBe(true)
	expect(extractor.Accept(elem.Field(7))).ToBe(true)
	expect(extractor.Accept(elem.Field(8))).ToBe(true)
	expect(extractor.Accept(elem.Field(9))).ToBe(true)
	expect(extractor.Accept(elem.Field(10))).ToBe(true)
	expect(extractor.Accept(elem.Field(11))).ToBe(false)
}

