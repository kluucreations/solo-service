package solo

import "github.com/kluucreations/solo-service/internal/datastore"

type service struct {
	ds datastore.Datastore
}
type Service interface{}

func NewService(ds datastore.Datastore) Service {
	return service{
		ds: ds,
	}
}
