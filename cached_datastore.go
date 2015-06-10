package db

import (
	"appengine/memcache"
)

type CachedDatastore struct {
	Datastore
}

func (this CachedDatastore) Load(m Entity) error {
	if err := this.ResolveKey(m); err != nil {
		return err
	}
	_, err := memcache.JSON.Get(this.context, m.StringId(), m)
	if err == memcache.ErrCacheMiss {
		if err := this.Datastore.Load(m); err != nil {
			return err
		}
	}
	return err
}

func (this CachedDatastore) Create(m Entity) error {
	if err := this.ResolveKey(m); err != nil {
		return err
	}

	item := &memcache.Item{
		Key:m.StringId(),
		Object: m,
	}

	if err := this.Datastore.Create(m); err != nil {
		return memcache.Delete(this.context, m.StringId())
	}

	return memcache.JSON.Set(this.context, item)
}

func (this CachedDatastore) Update(m Entity) error {
	if err := this.ResolveKey(m); err != nil {
		return err
	}

	item := &memcache.Item{
		Key:m.StringId(),
		Object: m,
	}

	if err := this.Datastore.Update(m); err != nil {
		return err
	}

	return memcache.JSON.Set(this.context, item)
}

func (this CachedDatastore) Delete(m Entity) error {
	if err := this.ResolveKey(m); err != nil {
		return err
	}

	if err := this.Datastore.Delete(m); err != nil {
		return err
	}

	return memcache.Delete(this.context, m.StringId())
}
