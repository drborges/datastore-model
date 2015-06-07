# datastore-model

Experimental somewhat ORM library for working with appengine `datastore`.

# Installation

```bash
$ go get -u github.com/drborges/datastore-model
```

# Usage

#### Simple Example

Embed `db.Model` to your model to add `db.Datastore` support:

```go
type Tag struct {
	db.Model
	Name  string
	Owner string
}
```

The model above, will be mapped to a datastore table named `Tag` derived from the struct's name and its datastore key is
auto generated as:

```go
datastore.NewKey(context, "Tag", "", 0, nil)
```

Struct `tags` are available to override the behavior above. 

#### Overriding Datastore Entity Kind

Just tag the embedded type `db.Model` with `db:"Tags"`, like so: 

```go
type Tag struct {
	db.Model    `db:"Tags"`
	Name  string
	Owner string
}
```

The key kind will be extracted from the tag, and the resulting key will look like:

```go
datastore.NewKey(context, "Tags", "", 0, nil)
```

#### Overriding Entity Key Generation

By tagging either a string or an integer field with `db:"id"`, `db.Datastore` will use it to create the entity's key.

```go
type Tag struct {
	db.Model     `db:"Tags"`
	Name  string `db:"id"`
	Owner string
}
```

# Datastore Operations

Consider the model below for the following examples.

```go
type Tag struct {
	db.Model     `db:"Tags"`
	Name string  `db:"id"`
	Owner string
}

// This is a convenient way to group
// queries for a given entity
type Tags []*Tag

func (this Tags) ByOwner(owner string) *db.Query {
	return db.From(new(Tag)).Filter("Owner=", owner)
}
```

#### Datastore.Create

The following code creates a new `Tag` in the `Tags` datastore entity:

```go
tag := new(Tag)
tag.Name  = "golang"
tag.Owner = "Borges"

err := db.NewDatastore(context).Create(tag)
```

Upon success the given entity has its key assigned to it, allowing you to access it by `tag.Key()`.

#### Datastore.CreateAll

The following code creates multiple entities `Tag` in a single batch:

```go
golang := new(Tag)
golang.Name = "golang"

appengine := new(Tag)
appengine.Name = "appengine"

err := db.NewDatastore(context).CreateAll(golang, appengine)
```

All created entities will have their respective keys assigned back to them upon success.

#### Datastore.Update

The following code updates a `Tag` in the `Tags` datastore entity:

```go
tag := new(Tag)
tag.Name  = "golang"
tag.Owner = "Borges"

err := db.NewDatastore(context).Update(tag)
```

Upon success the given entity has its key assigned to it, allowing you to access it by `tag.Key()`.

#### Datastore.UpdateAll

The following code updates multiple entities `Tag` in a single batch:

```go
golang := &Tag{Name: "golang"}
appengine := &Tag{Name: "appengine"}

err := db.NewDatastore(context).UpdateAll(golang, appengine)
```

All entities have their keys set to themselves upon success.

#### Datastore.Load

The following code loads data from datastore into a `Tag`:

```go
tag := new(Tag)
tag.Name = "golang"

err := db.NewDatastore(context).Load(tag)
// tag.Owner is populated with data
```

#### Datastore.LoadAll

The following code loads data from datastore into a `Tag`:

```go
golang := &Tag{Name: "golang"}
datastore := &Tag{Name: "datastore"}

err := db.NewDatastore(context).LoadAll(golang, datastore)
```

#### Datastore.Delete

The following code deletes an existent entity from datastore:

```go
tag := new(Tag)
tag.Name = "golang"

err := db.NewDatastore(context).Delete(tag)
```

#### Datastore.DeleteAll

**Warning** This API is still experimental and lots of changes might and likely will occur :)

The following code deletes multiple entities from datastore in a single batch:

```go
err := db.NewDatastore(context).DeleteAll(tag1, tag2)
```
#### Datastore.Query(q).All

The following code runs a given query and maps the matched items to a list of entities, setting their keys behind the seems.

```go
tags := Tags{}
err := db.NewDatastore(context).Query(tags.ByOwner(owner)).All(&tags)
```

For any given `tag` in the `tags` slice, one can access its key through `tag.Key()`

#### Datastore.Query(q).First

The following code runs a given query and loads the first item into the given entity instance.

```go
tag := new(Tag)
err := db.NewDatastore(context).Query(Tags{}.ByOwner(owner)).First(tag)
```

# Memcached Datastore

This is still another experiment. In order to reduce the hassle of implementing memcached operations on datastore we implemented `CachedDatastore` that abstracts that logic away.

`CachedDatastore` embeds the regular `Datastore` type only overriding some of its operations in order to add support to `memcache`.

Currently the operations supporting `memcache` are: `Load`, `Create`, `Update` and `Delete`.

Be aware if you use `CachedDatastore` as your default datastore access interface, using any operation other than the ones mentioned above there will not be caching going on.


#### Example:

```go
type MembershipCard struct {
	db.Model
	Number int   `db:"id"`
	Owner  string
}

cds := db.CachedDatastore{db.NewDatastore(c)}

card := &MembershipCard{Number: 1}
cds.Create(card)

cardFromCache := &MembershipCard{Number: 1}
err := cds.Load(cardFromCache)
```

`CachedDatastore` uses the encoded entity's key (`card.StringId()`) as the memcache key. As of now there is no mechanism to override that behavior. 

# Future Work

This is a very experimental project and there is a lot that can be done (iterator queries for instance...). This current work is essentially driven by BearchInc's use cases, though it is not restricted to them. Feel free to suggest and contribute.