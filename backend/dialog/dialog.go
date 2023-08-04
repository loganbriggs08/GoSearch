package dialog

import "tawesoft.co.uk/go/dialog"

func ErrorDialog(errorMessage string) {
	dialog.Alert(errorMessage, 4)
}