package db_test

import (
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"testing"
)

func TestExtractStringID(t *testing.T) {
	entity := EntityWithStringID{StringID: "Diego"}
	stringID, intID := db.ExtractEntityKeyIDs(&entity)

	expect := goexpect.New(t)
	expect(stringID).ToBe("Diego")
	expect(intID).ToBe(int64(0))
}

func TestExtractIntID(t *testing.T) {
	entity := EntityWithIntID{IntID: 123}
	stringID, intID := db.ExtractEntityKeyIDs(&entity)

	expect := goexpect.New(t)
	expect(stringID).ToBe("")
	expect(intID).ToBe(int64(123))
}

func TestExtractIDsChoosesFirstTaggedFieldAsID(t *testing.T) {
	entity := EntityWithMultipleIDTags{StringID: "Diego", IntID: 123}
	stringID, intID := db.ExtractEntityKeyIDs(&entity)

	expect := goexpect.New(t)
	expect(stringID).ToBe("")
	expect(intID).ToBe(int64(123))
}

func TestExtractEntityKindReturnsKindFromTag(t *testing.T) {
	kind := db.ExtractEntityKind(&Person{})

	expect := goexpect.New(t)
	expect(kind).ToBe("People")
}

func TestExtractEntityKindReturnsStructNameAsKind(t *testing.T) {
	kind := db.ExtractEntityKind(&EntityWithStringID{})

	expect := goexpect.New(t)
	expect(kind).ToBe("EntityWithStringID")
}
