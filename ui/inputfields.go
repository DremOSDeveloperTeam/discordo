package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rigormorrtiss/discordo/util"
	"github.com/rivo/tview"
)

func NewMessageInputField(onMessageInputFieldDone func(key tcell.Key), theme *util.Theme) *tview.InputField {
	messageInputField := tview.NewInputField()

	messageInputField.
		SetPlaceholder("Message...").
		SetFieldWidth(0).
		SetDoneFunc(onMessageInputFieldDone).
		SetFieldBackgroundColor(tcell.GetColor(theme.InputFieldBackground)).
		SetPlaceholderTextColor(tcell.GetColor(theme.InputFieldPlaceholderForeground)).
		SetBackgroundColor(tcell.GetColor(theme.InputFieldBackground)).
		SetBorder(true).
		SetBorderPadding(0, 0, 1, 1)

	return messageInputField
}