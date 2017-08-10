package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println("error getting wd: ", err)
	}
	dir, err := ioutil.ReadDir(wd)
	if err != nil {
		log.Println("error reading wd: ", err)
	}
	for _, fileinfo := range dir {
		if fileinfo.IsDir() {
			continue
		}
		newFolder := filepath.Join(
			wd,
			getFilename(fileinfo.Name()),
		)
		err = os.MkdirAll(newFolder, 0777)
		if err != nil {
			log.Println("error creating folder: ", err)
		}
		err = os.Rename(
			filepath.Join(wd, fileinfo.Name()),
			filepath.Join(newFolder, fileinfo.Name()),
		)
		if err != nil {
			log.Println("error moving file: ", err)
		}
	}
}

func getFilename(fpath string) string {
	fname := filepath.Base(fpath)
	ext := filepath.Ext(fname)
	return fname[:len(fname)-len(ext)]
}

func pathExists(fpath string) bool {
	_, err := os.Stat(fpath)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
