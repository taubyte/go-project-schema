package messaging

import "bitbucket.org/taubyte/go-project-schema/pretty"

func (m *messaging) Prettify(pretty.Prettier) map[string]interface{} {
	getter := m.Get()
	return map[string]interface {
	}{
		"Id":           getter.Id(),
		"Name":         getter.Name(),
		"Description":  getter.Description(),
		"Tags":         getter.Tags(),
		"Local":        getter.Local(),
		"Regex":        getter.Regex(),
		"ChannelMatch": getter.ChannelMatch(),
		"MQTT":         getter.MQTT(),
		"WebSocket":    getter.WebSocket(),
	}
}
