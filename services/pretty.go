package services

import "github.com/taubyte/go-project-schema/pretty"

func (s *service) Prettify(pretty.Prettier) map[string]interface{} {
	getter := s.Get()

	return map[string]interface {
	}{
		"Id":          getter.Id(),
		"Name":        getter.Name(),
		"Description": getter.Description(),
		"Tags":        getter.Tags(),
		"Protocol":    getter.Protocol(),
	}
}
