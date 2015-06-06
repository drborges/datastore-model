package db_test

import (
	"github.com/drborges/datastore-model"
	"github.com/drborges/goexpect"
	"testing"
)

func TestExtractStringID(t *testing.T) {
	t.Parallel()
	entity := EntityWithStringID{StringID: "Diego"}
	stringID, intID, err := db.ExtractEntityKeyIDs(&entity)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(stringID).ToBe("Diego")
	expect(intID).ToBe(int64(0))
}

func TestExtractIntID(t *testing.T) {
	t.Parallel()
	entity := EntityWithIntID{IntID: 123}
	stringID, intID, err := db.ExtractEntityKeyIDs(&entity)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(stringID).ToBe("")
	expect(intID).ToBe(int64(123))
}

func TestExtractIDsChoosesFirstTaggedFieldAsID(t *testing.T) {
	t.Parallel()
	entity := EntityWithMultipleIDTags{StringID: "Diego", IntID: 123}
	stringID, intID, err := db.ExtractEntityKeyIDs(&entity)

	expect := goexpect.New(t)
	expect(err).ToBe(nil)
	expect(stringID).ToBe("")
	expect(intID).ToBe(int64(123))
}

func TestExtractIDsReturnsErrMissingStringId(t *testing.T) {
	t.Parallel()
	entity := new(EntityWithStringID)
	stringID, intID, err := db.ExtractEntityKeyIDs(entity)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrMissingStringId)
	expect(stringID).ToBe("")
	expect(intID).ToBe(int64(0))
}

func TestExtractIDsReturnsErrMissingIntId(t *testing.T) {
	t.Parallel()
	entity := new(EntityWithIntID)
	stringID, intID, err := db.ExtractEntityKeyIDs(entity)

	expect := goexpect.New(t)
	expect(err).ToBe(db.ErrMissingIntId)
	expect(stringID).ToBe("")
	expect(intID).ToBe(int64(0))
}

func TestExtractEntityKindReturnsKindFromTag(t *testing.T) {
	t.Parallel()
	kind := db.ExtractEntityKind(&Person{})

	expect := goexpect.New(t)
	expect(kind).ToBe("People")
}

func TestExtractEntityKindReturnsStructNameAsKind(t *testing.T) {
	t.Parallel()
	kind := db.ExtractEntityKind(&EntityWithStringID{})

	expect := goexpect.New(t)
	expect(kind).ToBe("EntityWithStringID")
}
