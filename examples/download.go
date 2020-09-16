package main

import (
	"fmt"
	"github.com/Navid2zp/idm"
	"time"
)

func main() {
	// Calling VerifyIDM is not necessary. NewDownload will check it.
	exits, err := idm.VerifyIDM()
	if err != nil {
		fmt.Println("error verifying IDM:", err.Error())
		return
	}
	if !exits {
		fmt.Println("IDM is not installed on this computer!")
		return
	}

	download, _ := idm.NewDownload("https://codeload.github.com/Navid2zp/idm/zip/master")
	download.Silent().QuitAfterFinish()
	// set download path
	download.SetFilePath("C:")
	// Start the download
	_ = download.Start()

	// Wait till file is appeared in the given path
	err = download.VerifyDownload(time.Second * 10)
	if err != nil {
		fmt.Println("couldn't verify file download:", err.Error())
	}
}
