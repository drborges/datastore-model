package db

import (
	"appengine/memcache"
)

type CachedDatastore struct {
	Datastore
}

func (this CachedDatastore) Load(m Entity) error {
	metadata, err := this.ResolveKey(m)
	if err != nil {
		return err
	}

	if _, err := memcache.JSON.Get(this.context, metadata.CacheStringID, m); err == memcache.ErrCacheMiss {
		if err := this.Datastore.Load(m); err != nil {
			return err
		}
	}

	return err
}

func (this CachedDatastore) Create(m Entity) error {
	metadata, err := this.ResolveKey(m)
	if err != nil {
		return err
	}

	if err := this.Datastore.Create(m); err != nil {
		return memcache.Delete(this.context, m.StringId())
	}

	return memcache.JSON.Set(this.context, &memcache.Item{
		Key:    metadata.CacheStringID,
		Object: m,
	})
}

func (this CachedDatastore) Update(m Entity) error {
	metadata, err := this.ResolveKey(m)
	if err != nil {
		return err
	}

	if err := this.Datastore.Update(m); err != nil {
		return err
	}

	return memcache.JSON.Set(this.context, &memcache.Item{
		Key:    metadata.CacheStringID,
		Object: m,
	})
}

func (this CachedDatastore) Delete(m Entity) error {
	metadata, err := this.ResolveKey(m)
	if err != nil {
		return err
	}

	if err := this.Datastore.Delete(m); err != nil {
		return err
	}

	return memcache.Delete(this.context, metadata.CacheStringID)
}
