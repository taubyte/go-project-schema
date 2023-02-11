package services

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
)

func Open(seer *seer.Seer, name, application string) (Service, error) {
	service := &service{
		seer:        seer,
		name:        name,
		application: application,
	}

	var err error
	service.Resource, err = basic.New(seer, service, name)
	if err != nil {
		return nil, err
	}

	return service, nil
}
