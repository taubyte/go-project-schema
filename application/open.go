package application

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-seer"
)

// Open opens the application at root/applications/<name>, returns Application and error
func Open(seer *seer.Seer, name string) (Application, error) {
	app := &application{
		seer: seer,
		name: name,
	}

	var err error
	app.Resource, err = basic.New(seer, app, name)
	if err != nil {
		return nil, err
	}

	app.Resource.Root = app.Root
	app.Resource.Config = app.Config

	return app, nil
}
