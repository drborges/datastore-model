package db

import (
	"appengine/datastore"
)

// Model represents a datastore entity
//
// In order to add Datastore support to
// your models have a non pointer Entity
// embedded
type Model struct {
	key *datastore.Key `json:"-",datastore:"-"`
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
	this.key = key
	return err
}

func (this *Model) setKey(k *datastore.Key) {
	this.key = k
}
