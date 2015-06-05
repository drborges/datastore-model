package db

import (
	"appengine/datastore"
	"errors"
	"time"
)

var (
	// ErrInvalidUUID returned if an invalid
	// UUID is passed to SetUUID
	ErrInvalidUUID = errors.New("Invalid UUID")
)

// Model represents a datastore entity
//
// Embedding this type to a struct allows
// it to be used as an entity type in
// Datastore service
type Entity struct {
	key       *datastore.Key `json:"-" datastore:"-"`
	parentKey *datastore.Key `json:"-" datastore:"-"`
	CreatedAt time.Time      `json:"-" datastore:",noindex"`
}

// HasKey returns true in case the
// current instance already has a
// datastore key assigned to it
//
// Returns false otherwise
func (this *Entity) HasKey() bool {
	return this.key != nil
}

// Key returns the entity datastore key
func (this *Entity) Key() *datastore.Key {
	return this.key
}

// ParentKey returns the entity's parent datastore key
func (this *Entity) Parent() *datastore.Key {
	return this.parentKey
}

// SetParent sets the entity's parent key
func (this *Entity) SetParent(parent *datastore.Key) {
	this.parentKey = parent
}

func (this *Entity) SetCreatedAt(t time.Time) {
	this.CreatedAt = t
}

// KeyAsUUID Returns the UUID representation of
// the entity's datastore key
//
// An empty string is returned in case
// the current key is invalid
func (this *Entity) KeyAsUUID() string {
	return this.key.Encode()
}

// SetKeyFromUUID assigns a datastore key to the entity
// based on the given UUID
//
// Currently the UUID is the encoded datastore key
func (this *Entity) SetKeyFromUUID(uuid string) error {
	key, err := datastore.DecodeKey(uuid)
	if err != nil {
		return ErrInvalidUUID
	}
	this.key = key
	return err
}

// SetKey sets the entity datastore Key
func (this *Entity) SetKey(k *datastore.Key) {
	this.key = k
}
