package main

import (
	"fmt"

	"github.com/ayntgl/discordo/config"
)

// Whether the mouse is usable or not.
var Mouse = true

// The maximum number of messages to fetch and display on the messages panel.
// Its value must not be lesser than 1 and greater than 100.
var MessagesLimit = 50

// Whether to display the timestamp of the message beside the displayed message or not.
var Timestamps = false

// The timezone of the timestamps.
// Learn more: https://pkg.go.dev/time#LoadLocation
var Timezone = "Local"

// A textual representation of the time value formatted according to the layout defined by its value.
// Learn more: https://pkg.go.dev/time#Layout
var TimeFormat = "3:04PM"

var browser = "Chrome"
var browserVersion = "104.0.5112.102"
var oss = "Linux"

// Identify properties are connection properties that are dispatched in the IDENTIFY gateway event to trigger the initial handshake with the gateway.
// Learn more: https://discord.com/developers/docs/topics/gateway#identify
var IdentifyProperties = config.IdentifyConfig{
	UserAgent: fmt.Sprintf(
		"Mozilla/5.0 (X11; %s x86_64) AppleWebKit/537.36 (KHTML, like Gecko) %s/%s Safari/537.36",
		oss,
		browser,
		browserVersion,
	),
	Browser:        browser,
	BrowserVersion: browserVersion,
	OS:             oss,
}

// Keybindings
var Keys = config.Keys{
	Application: []config.Key{
		config.NewKey("Rune[g]", "Focus the guilds tree widget.", func(core any, event any) {}),
		config.NewKey("Rune[c]", "Focus the channels tree widget.", func(core any, event any) {}),
		config.NewKey("Rune[m]", "Focus the messages panel widget.", func(core any, event any) {}),
		config.NewKey("Rune[i]", "Focus the message input widget.", func(core any, event any) {}),
	},
	// MessagesPanel: {
	//     key(
	//         "Rune[a]",
	//         "Open the message actions list widget.",
	//         function(core, event)
	//             return openMessageActionsList()
	//         end
	//     ),
	//     key(
	//         "Up",
	//         "Select the previous message.",
	//         function(core, event)
	//             return selectPreviousMessage()
	//         end
	//     ),
	//     key(
	//         "Down",
	//         "Select the next message.",
	//         function(core, event)
	//             return selectNextMessage()
	//         end
	//     ),
	//     key(
	//         "Home",
	//         "Select the first message.",
	//         function(core, event)
	//             return selectFirstMessage()
	//         end
	//     ),
	//     key(
	//         "End",
	//         "Select the last message.",
	//         function(core, event)
	//             return selectLastMessage()
	//         end
	//     )
	// },
	// MessageInput: {
	//     key(
	//         "Ctrl+E",
	//         "Open the external editor.",
	//         function()
	//             return openExternalEditor()
	//         end
	//     ),
	//     key(
	//         "Ctrl+V",
	//         "Paste the clipboard content.",
	//         function()
	//             return pasteClipboardContent()
	//         end
	//     )
	// },
}
