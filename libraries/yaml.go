package libraries

import "bitbucket.org/taubyte/go-project-schema/basic"

func Yaml(name, application string, yamlData []byte) (Getter, error) {
	resource, err := basic.Yaml(yamlData)
	if err != nil {
		return nil, err
	}

	return getter{
		&library{
			Resource:    resource,
			name:        name,
			application: application,
		},
	}, nil
}
