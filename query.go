package db

import "appengine/datastore"

func From(e entity) *Query {
	resolver := NewKeyResolver(nil)
	resolver.ExtractKindMetadata(e)
	return &Query{datastore.NewQuery(resolver.Metadata.Kind)}
}

type Query struct {
	*datastore.Query
}

func (this *Query) Filter(filter string, value interface{}) *Query {
	this.Query = this.Query.Filter(filter, value)
	return this
}
