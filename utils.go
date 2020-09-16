package idm

import (
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Validate a url
func isUrl(str string) bool {
	u, err := url.Parse(str)
	// IDM accepts urls without schema
	return err == nil && u.Host != ""
}

// Build the final args
func (d *Download) buildArgs() []string {
	args := []string{"/d", d.URL}
	if d.SilentMode {
		args = append(args, "/n")
	}
	if d.Quit {
		args = append(args, "/q")
	}
	if d.HangUp {
		args = append(args, "/h")
	}

	if d.Path != "" {
		args = append(args, "/p", d.Path)
	}

	if d.FileName != "" {
		args = append(args, "/f", d.FileName)
	}

	return args
}

// Get the header for url
// Used to extract file name
func (d *Download) getAndSetFileHeader() error {
	resp, err := http.Head(d.URL)
	if err != nil {
		return &URLHeaderError{err: err}
	}
	if resp.StatusCode != http.StatusOK {
		return URLStatusCodeError
	}
	d.URLHeader = resp
	return nil
}

// Get the file size from url header
func (d *Download) setFileName() error {
	var err error
	if d.URLHeader == nil {
		err = d.getAndSetFileHeader()
		if err != nil {
			return err
		}
	}

	// file name might be present in Content-Disposition header
	// e.g: attachment; filename=file.zip
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition
	cd := d.URLHeader.Header.Get("Content-Disposition")
	if cd == "" {
		return FileNameHeaderError
	}
	sp := strings.Split(cd, " ")
	for _, s := range sp {
		if strings.Contains(s, "filename=") {
			d.FileName = strings.Split(s, "=")[1]
			return nil
		}
	}
	return FileNameHeaderError
}

// Wait's until the file is appeared in the given path or it times out
func waitForFileToAppear(path string, timeout time.Duration) error {
	tick := time.Tick(timeout)
	for {
		select {
		case <-tick:
			return WaitForDownloadTimeOutError
			// Wait 1 sec before checking again
		case <-time.After(time.Second * 1):
			if _, err := os.Stat(path); err == nil {
				return nil
			}
		}
	}
}
