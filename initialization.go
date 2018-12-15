package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	log "github.com/trubeck/simpleLogger"

	"github.com/trubeck/gopository/services"
	"github.com/trubeck/gopository/storage"
)

func scanFolders() {
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println("Cannot open Path: ", err)
	}

	for _, f := range files {

		if f.IsDir() && !strings.HasPrefix(f.Name(), ".") {
			getFiles(f.Name())
		}

	}
}

func getFiles(pkgName string) {

	files, err := ioutil.ReadDir(basePath + "/" + pkgName)
	if err != nil {
		fmt.Println("Cannot open Path: ", err)
	}

	for _, f := range files {

		if !f.IsDir() {
			split1 := strings.Split(f.Name(), "_")

			major, minor, patch, err := services.ParseVersion(split1[len(split1)-1], ".")
			if err != nil {
				log.Error(err)
			}

			version := storage.Storage[pkgName]

			version = services.ExpandStorage(major, minor, patch, version)

			log.Trace(pkgName)
			log.Trace(major)
			log.Trace(minor)
			log.Trace(patch)

			version[major][minor][patch] = basePath + "/" + pkgName + "/" + f.Name()
			storage.Storage[pkgName] = version

		}

	}

}
