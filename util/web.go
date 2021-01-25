package util

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func DownloadAndSave(url, dirName, fileName string) error {

	newReq, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(newReq)

	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = CreateDirIfNotExists(dirName); err != nil {
		return err
	}

	return Write(filepath.Join(dirName, fileName), string(body))
}
