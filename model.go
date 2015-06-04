package db

import (
	"appengine/datastore"
	"errors"
)

var (
	ErrInvalidUUID = errors.New("Invalid UUID")
)

// Model represents a datastore entity
//
// In order to add Datastore support to
// your models have a non pointer Entity
// embedded
type Model struct {
	key *datastore.Key       `json:"-",datastore:"-"`
	parentKey *datastore.Key `json:"-",datastore:"-"`
}

// HasKey returns true in case the
// current instance already has a
// datastore key assigned to it
//
// Returns false otherwise
func (this *Model) HasKey() bool {
	return this.key != nil
}

// Key returns the entity datastore key
func (this *Model) Key() *datastore.Key {
	return this.key
}

// ParentKey returns the entity's parent datastore key
func (this *Model) Parent() *datastore.Key {
	return this.parentKey
}

// SetParent sets the entity's parent key
func (this *Model) SetParent(parent *datastore.Key) {
	this.parentKey = parent
}

// UUID Returns the UUID representation of
// the entity's datastore key
//
// An empty string is returned in case
// the current key is invalid
func (this *Model) UUID() string {
	return this.key.Encode()
}

// SetUUID assigns a datastore key to the entity
// based on the given UUID
//
// Currently the UUID is the encoded datastore key
func (this *Model) SetUUID(uuid string) error {
	key, err := datastore.DecodeKey(uuid)
	if err != nil {
		return ErrInvalidUUID
	}
	this.key = key
	return err
}

// SetKey sets the entity datastore Key
func (this *Model) SetKey(k *datastore.Key) {
	this.key = k
}
