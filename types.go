package idm

import "net/http"

var idmInstalled = false
var idmInstallationChecked = false
var idmPath = ""

// Download represents the IDM parameters and data needed to verify download completion
type Download struct {
	URL string
	// /p
	Path string
	// /f
	FileName string
	// /q
	Quit bool
	// /h
	HangUp bool
	// /n
	SilentMode bool
	// Used to extract file size and file name
	URLHeader *http.Response
}
