package idm

import (
	"errors"
	"fmt"
)

// InValidURLError represents the error for when the provided url is not a valid url
var InValidURLError = errors.New("provided url is not a valid url")

// IDMNotInstalledError represents the error when idm is not installed
var IDMNotInstalledError = errors.New("idm is not installed")

// URLStatusCodeError represents the error when url didn't return a status code 200
var URLStatusCodeError = errors.New("url didn't return status code 200")

// FileNameHeaderError represents the error when the file name isn't present in url headers
var FileNameHeaderError = errors.New("file name is not available in url headers")

// FileNameDetectionError represents the error when program can't find file name in header and no local file name is provided
var FileNameDetectionError = errors.New("file name is not specified and it doesn't exist on url header either")

// DownloadFilePathNotProvidedError represents the error when no file path is provided but one is required
var DownloadFilePathNotProvidedError = errors.New("no file path is provided for download. (necessary for download verification)")

// WaitForDownloadTimeOutError represents the error when time out duration is reached while waiting for a file download to be complete
var WaitForDownloadTimeOutError = errors.New("file didn't appear in specified duration")

// ApplicationsListError represents the error when program can't list the installed programs
// NOTE: only windows
type ApplicationsListError struct {
	err error
}

func (e *ApplicationsListError) Error() string {
	return fmt.Sprintf("failed to get applications list: %s", e.err.Error())
}

// URLHeaderError represents the error when something goes wrong while fetching the url header
type URLHeaderError struct {
	err error
}

func (e *URLHeaderError) Error() string {
	return fmt.Sprintf("failed to get url header: %s", e.err.Error())
}
