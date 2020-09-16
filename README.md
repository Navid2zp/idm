# idm
Golang wrapper for Internet Download Manager (IDM) CLI.

### Install

To use in a go project:
```
go get github.com/Navid2zp/idm
```

### Usage

```go
package main

import (
	"fmt"
	"github.com/Navid2zp/idm"
	"time"
)

func main() {
	download, _ := idm.NewDownload("https://codeload.github.com/Navid2zp/idm/zip/master")
	// Silent mode and quit after finish
	download.Silent().QuitAfterFinish()
	// set download path
	download.SetFilePath("C:")
	// Start the download
	_ = download.Start()

	// Wait till file is appeared in the given path
	err := download.VerifyDownload(time.Second * 10)
	if err != nil {
		fmt.Println("couldn't verify file download:", err.Error())
	}
}
```

#### Parameters
Check out the IDM methods here: https://www.internetdownloadmanager.com/support/command_line.html

```go
// Turns on the silent mode when IDM doesn't ask any questions
download.Silent()

// IDM will hang up your connection after the successful downloading
download.HangUpAfterFinish()

// IDM will exit after the successful downloading.
// This parameter works only for the first copy
download.QuitAfterFinish()

// Defines the local path where to save the file
download.SetFilePath("C:/Users/MyPC/Downloads")

// Defines local file name to save the file
download.SetFileName("myFile.zip")

// adds specified file to download queue, but don't start downloading
download.AddToQueue()
```

**Starting the main IDM queue:**

```go
err := idm.StartMainQueue()
```

##### Verify IDM installation

```go
installed, _ := idm.VerifyIDM()

if !installed {
    fmt.Println("Couldn't find IDM")
} else {
    fmt.Println("IDM found")
}
```

This package uses [go-win64api][1]'s `InstalledSoftwareList` method to list all the installed applications and checks if IDM is present. Then it will try to find it's path. It fails to verify IDM if `IDMan.exe` won't be present in the same directory as `Uninstall.exe`

You can set the IDM path manually:

```go
idm.SetIDMPath("path/to/idm.exe")
```


##### Verify download

There is no way to be sure if a download is completed or not since we can't actually control the IDM. The closest way to check if a download is finished is by checking the download path and look for the file.

```go
err := download.VerifyDownload(time.second * 30)
```

**NOTE:** You have to specify download path using `SetFilePath` method to use this option. Also I recommend setting the file name using `SetFileName` method to make sure that we know where the file is to verify it.

If filename isn't specified, program will try to find it from URL header which might not be available or be different than what IDM chooses causing this method to fail the verification.


License
----

[MIT][2]

[1]: https://github.com/iamacarpet/go-win64api
[2]: https://github.com/Navid2zp/idm/blob/master/LICENSE
