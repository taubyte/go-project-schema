package project

import (
	"bitbucket.org/taubyte/go-project-schema/application"
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/databases"
	"bitbucket.org/taubyte/go-project-schema/domains"
	"bitbucket.org/taubyte/go-project-schema/functions"
	"bitbucket.org/taubyte/go-project-schema/libraries"
	"bitbucket.org/taubyte/go-project-schema/messaging"
	"bitbucket.org/taubyte/go-project-schema/pretty"
	"bitbucket.org/taubyte/go-project-schema/services"
	"bitbucket.org/taubyte/go-project-schema/smartops"
	"bitbucket.org/taubyte/go-project-schema/storages"
	"bitbucket.org/taubyte/go-project-schema/website"
	"github.com/taubyte/go-seer"
)

type Project interface {
	Get() Getter
	Set(sync bool, ops ...basic.Op) (err error)
	Delete(attributes ...string) (err error)
	Prettify(p pretty.Prettier) map[string]interface{}
	ResourceMethods() []pretty.PrettyResourceIface

	Application(name string) (application.Application, error)
	Database(name string, application string) (databases.Database, error)
	Domain(name string, application string) (domains.Domain, error)
	Function(name string, application string) (functions.Function, error)
	Library(name string, application string) (libraries.Library, error)
	Messaging(name string, application string) (messaging.Messaging, error)
	Service(name string, application string) (services.Service, error)
	SmartOps(name string, application string) (smartops.SmartOps, error)
	Storage(name string, application string) (storages.Storage, error)
	Website(name string, application string) (website.Website, error)
}

type project struct {
	*basic.Resource
	seer *seer.Seer
}

type Getter interface {
	basic.Getter
	Applications() []string
	Tags() []string
	Email() string
	Services(string) (local []string, global []string)
	Libraries(string) (local []string, global []string)
	Websites(string) (local []string, global []string)
	Messaging(string) (local []string, global []string)
	Databases(string) (local []string, global []string)
	Storages(string) (local []string, global []string)
	Domains(string) (local []string, global []string)
	SmartOps(string) (local []string, global []string)
	Functions(string) (local []string, global []string)
}
