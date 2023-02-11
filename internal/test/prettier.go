package internal

import (
	"bitbucket.org/taubyte/go-project-schema/pretty"
	commonSpec "bitbucket.org/taubyte/go-specs/common"
)

type PrettierFetchMethod func(path *commonSpec.TnsPath) (pretty.Object, error)
type PrettierStringMethod func() string

type Setter interface {
	Fetch(PrettierFetchMethod)
	Project(PrettierStringMethod)
	Branch(PrettierStringMethod)
	AssetCID(cid string)
}

type Prettier interface {
	Fetch(path *commonSpec.TnsPath) (pretty.Object, error)
	Project() string
	Branch() string

	Set() Setter
}

type prettier struct {
	cid     string
	fetch   PrettierFetchMethod
	project PrettierStringMethod
	branch  PrettierStringMethod
}

func (p *prettier) Fetch(path *commonSpec.TnsPath) (pretty.Object, error) {
	return p.fetch(path)
}

func (p *prettier) Project() string {
	return p.project()
}

func (p *prettier) Branch() string {
	return p.branch()
}

type setter struct {
	*prettier
}

func (s *setter) Fetch(method PrettierFetchMethod) {
	s.fetch = method
}

func (s *setter) Project(method PrettierStringMethod) {
	s.project = method
}

func (s *setter) Branch(method PrettierStringMethod) {
	s.branch = method
}

func (s *setter) AssetCID(cid string) {
	s.cid = cid
}

func (p *prettier) Set() Setter {
	return &setter{p}
}

type object struct {
	*prettier
}

func (o *object) Interface() interface{} {
	return o.cid
}

// Used for overriding calls to Prettier and testing error returns
func NewMockPrettier() Prettier {
	p := &prettier{
		project: func() string {
			return "test_project_id"
		},
		branch: func() string {
			return "main"
		},
	}

	p.fetch = func(path *commonSpec.TnsPath) (pretty.Object, error) {
		return &object{p}, nil
	}

	return p
}
