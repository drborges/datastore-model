# datastore-model

Experimental somewhat ORM library for working with appengine `datastore`.

# Installation

```bash
$ go get -u github.com/drborges/datastore-model
```

# Usage: Entity

#### Simple Example

Embed `db.Entity` to your model to add `db.Datastore` support:

```go
type Tag struct {
	db.Entity
	Name string
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

Just tag the embedded type `db.Entity` with `db:"MyEntityKind"`, like so: 

```go
type Tag struct {
	db.Entity    `db:"Tags"`
	Name  string
	Owner string
}
```

#### Overriding Entity Key Generation

Just tag either a string or an integer field with `db:"id"` and `db.Datastore` will use it in the key creation

```go
type Tag struct {
	db.Entity    `db:"Tags"`
	Name  string `db:"id"`
	Owner string
}
```

# Usage: Datastore

Datastore is a essentially an extension of appengine's datastore which brings some handy features to you.

For more detailed info on the behavior of the following operations, check the godocs.

Consider the model below for the next examples.

```go
type Tag struct {
	db.Entity    `db:"Tags"`
	Name string  `json:"name" db:"id"`
	Owner string `json:"owner"`
}

type Tags []*Tag

func (this Tags) ByOwner(owner string) *db.Query {
	return db.QueryFor(new(Tag)).Filter("Owner=", owner)
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

#### Datastore.Load

The following code loads data from datastore into a `Tag`:

```go
tag := new(Tag)
tag.Name = "golang"

err := db.NewDatastore(context).Load(tag)
```

#### Datastore.Delete

The following code deletes an existent entity from datastore:

```go
tag := new(Tag)
tag.Name = "golang"

err := db.NewDatastore(context).Delete(tag)
```

#### Datastore.Query(q).All

The following code runs a given query and maps the matched items to a list of entities, setting their keys behind the seems.

```go
tags := Tags{}
err := db.NewDatastore(context).Query(tags.ByOwner(owner)).All(&tags)
```

#### Datastore.Query(q).First

The following code runs a given query and loads the first item into the given entity instance.

```go
tag := new(Tag)
err := db.NewDatastore(context).Query(Tags{}.ByOwner(owner)).First(tag)
```