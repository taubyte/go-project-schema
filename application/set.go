package application

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
)

func Id(value string) basic.Op {
	return basic.Set("id", value)
}

func Description(value string) basic.Op {
	return basic.Set("description", value)
}

func Tags(value []string) basic.Op {
	return basic.Set("tags", value)
}
