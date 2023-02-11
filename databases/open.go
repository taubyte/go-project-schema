package databases

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
)

func Open(seer *seer.Seer, name string, application string) (Database, error) {
	database := &database{
		seer:        seer,
		name:        name,
		application: application,
	}

	var err error
	database.Resource, err = basic.New(seer, database, name)
	if err != nil {
		return nil, err
	}

	return database, nil
}
