package messaging

import (
	"github.com/taubyte/go-project-schema/basic"
	"github.com/taubyte/go-project-schema/common"
	"github.com/taubyte/go-seer"
	structureSpec "github.com/taubyte/go-specs/structure"
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
