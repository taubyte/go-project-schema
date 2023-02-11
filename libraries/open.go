package libraries

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
)

func Open(seer *seer.Seer, name string, application string) (Library, error) {
	library := &library{
		seer:        seer,
		name:        name,
		application: application,
	}

	var err error
	library.Resource, err = basic.New(seer, library, name)
	if err != nil {
		return nil, err
	}

	return library, nil
}
