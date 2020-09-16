package idm

import (
	wapi "github.com/iamacarpet/go-win64api"
	"path/filepath"
	"strings"
	"time"
)

// GetIDMPath returns the IDM installation path if idm is installed
// Returns empty string if idm isn't installed or failed to get the applications list
// Call VerifyIDM first if you're not sure if IDM is installed
func GetIDMPath() string {
	if !idmInstallationChecked {
		_, _ = VerifyIDM()
	}
	return idmPath
}

// VerifyIDM verifies the IDM installation on the machine
func VerifyIDM() (bool, error) {
	if idmInstallationChecked {
		return idmInstalled, nil
	}
	// Get the list of installed applications
	sw, err := wapi.InstalledSoftwareList()
	if err != nil {
		idmInstalled = true
		idmInstallationChecked = true
		return false, &ApplicationsListError{err: err}
	}

	for _, s := range sw {
		if s.Name() == "Internet Download Manager" {
			idmInstalled = true
			idmInstallationChecked = true
			// IDM installation path is not available as s.InstallLocation
			// So replace Uninstall.exe with IDMs' standard name "IDMan.exe" to get the exe path
			idmPath = strings.Replace(s.UninstallString, "Uninstall.exe", "IDMan.exe", 1)
			break
		}
	}
	return idmInstalled, nil
}

// StartMainQueue starts the main queue in IDM
func StartMainQueue() error {
	installed, err := VerifyIDM()
	if err != nil {
		return err
	}
	if !installed {
		return IDMNotInstalledError
	}
	return mainQueueStart()
}

// AddToQueue adds the given url to IDM queue
// File download won't be started
func AddToQueue(url string) error {
	if !isUrl(url) {
		return InValidURLError
	}
	return addToQueue(url)
}

// NewDownload creates a new download
func NewDownload(url string) (*Download, error) {
	if !idmInstallationChecked {
		exits, err := VerifyIDM()
		if err != nil {
			return nil, err
		}
		if !exits {
			return nil, IDMNotInstalledError
		}
	}
	if !isUrl(url) {
		return nil, InValidURLError
	}
	download := Download{}
	download.URL = url
	return &download, nil
}

func (d *Download) AddToQueue() error {
	return AddToQueue(d.URL)
}

// Silent sets the /n param
// turns on the silent mode when IDM doesn't ask any questions
func (d *Download) Silent() *Download {
	d.SilentMode = true
	return d
}

// QuitAfterFinish sets the /q param
// IDM will exit after the successful downloading. This parameter works only for the first copy
func (d *Download) QuitAfterFinish() *Download {
	d.Quit = true
	return d
}

// HangUpAfterFinish sets the /h param
// IDM will hang up your connection after the successful downloading
func (d *Download) HangUpAfterFinish() *Download {
	d.Quit = true
	return d
}

// SetFilePath sets the /p param
// Defines the local path where to save the file
func (d *Download) SetFilePath(path string) *Download {
	d.Path = path
	return d
}

// SetFilePath sets the /f param
// Defines local file name to save the file
func (d *Download) SetFileName(name string) *Download {
	d.FileName = name
	return d
}

// Start starts the file download in IDM
func (d *Download) Start() error {
	if !isUrl(d.URL) {
		return InValidURLError
	}
	return startDownload(d.buildArgs())
}

// WaitForFinish waits for download to be completed
// Only works if you specified the download path
// NOTE: File path must've been provided and file name must either be provided or be present in url header
func (d *Download) WaitForFinish(timeout time.Duration) error {
	if d.Path == "" {
		return DownloadFilePathNotProvidedError
	}
	err := d.setFileName()
	if err != nil {
		return err
	}

	fullPath, err := d.GetFullPath()
	if err != nil {
		return err
	}
	err = waitForFileToAppear(fullPath, timeout)
	return err
}

// GetFullPath returns the full path of the file
// File path must be provided using SetFilePath method
// If file name hasn't been set, it will try to get it from url header (returns error if not available in header)
func (d *Download) GetFullPath() (string, error) {
	if d.Path == "" {
		return "", DownloadFilePathNotProvidedError
	}
	if d.FileName == "" {
		err := d.setFileName()
		if err != nil {
			return "", FileNameDetectionError
		}
	}

	return filepath.Join(d.Path, d.FileName), nil
}
