package config

type Keys struct {
	Application   []Key
	GuildsTree    []Key
	ChannelsTree  []Key
	MessagesPanel []Key
	MessageInput  []Key
}

type Key struct {
	Name        string
	Description string
	Action      func(core any, event any)
}

func NewKey(name string, description string, action func(core any, event any)) Key {
	return Key{
		Name:        name,
		Description: description,
		Action:      action,
	}
}

type IdentifyConfig struct {
	UserAgent      string
	Browser        string
	BrowserVersion string
	OS             string
}
