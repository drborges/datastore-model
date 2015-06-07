package db

import (
	"appengine/datastore"
	"time"
)

// Model represents a datastore entity
//
// Embedding this type to a struct allows
// it to be used as an entity type in
// Datastore service
type Model struct {
	key       *datastore.Key `json:"-" datastore:"-"`
	parentKey *datastore.Key `json:"-" datastore:"-"`
	CreatedAt time.Time      `json:"-" datastore:",noindex"`
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

// SetCreatedAt sets the entity creation time
func (this *Model) SetCreatedAt(t time.Time) {
	this.CreatedAt = t
}

// StringId Returns the string representation of the datastore key
//
// An empty string is returned in case the key is invalid
func (this *Model) StringId() string {
	return this.key.Encode()
}

// SetStringId decodes the give string into a datastore key
//
// Currently the Id is the encoded datastore key
func (this *Model) SetStringId(uuid string) error {
	key, err := datastore.DecodeKey(uuid)
	this.key = key
	return err
}

// SetKey sets the entity datastore Key
func (this *Model) SetKey(k *datastore.Key) {
	this.key = k
}
