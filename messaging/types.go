package messaging

import (
	"bitbucket.org/taubyte/go-project-schema/basic"
	"bitbucket.org/taubyte/go-project-schema/common"
	structureSpec "bitbucket.org/taubyte/go-specs/structure"
	"github.com/taubyte/go-seer"
)

type Messaging interface {
	Get() Getter
	common.Resource[*structureSpec.Messaging]
}

type messaging struct {
	*basic.Resource
	seer        *seer.Seer
	name        string
	application string
}

type Getter interface {
	basic.ResourceGetter[*structureSpec.Messaging]
	Local() bool
	Regex() bool
	ChannelMatch() string
	MQTT() bool
	WebSocket() bool
}
