package website

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
)

func Open(seer *seer.Seer, name string, application string) (Website, error) {
	website := &website{
		seer:        seer,
		name:        name,
		application: application,
	}

	var err error
	website.Resource, err = basic.New(seer, website, name)
	if err != nil {
		return nil, err
	}

	return website, nil
}
