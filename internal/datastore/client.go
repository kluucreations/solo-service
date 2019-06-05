package datastore

type datastore struct{}

type Datastore interface{}

func NewClient() Datastore {
	return datastore{}
}
