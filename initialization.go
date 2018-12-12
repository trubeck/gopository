package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
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
			split2 := strings.Split(split1[len(split1)-1], ".")

			major, err := strconv.Atoi(split2[0])
			if err != nil {
				fmt.Println(err)
			}

			minor, err := strconv.Atoi(split2[1])
			if err != nil {
				fmt.Println(err)
			}

			patch, err := strconv.Atoi(split2[2])
			if err != nil {
				fmt.Println(err)
			}

			version := storage[pkgName]

			if major > len(version)-1 {
				for i := len(version) - 1; i < major+1; i++ {
					version = append(version, make([][]string, 0))
				}
			}

			if minor > len(version[major])-1 {
				for i := len(version[major]) - 1; i < minor+1; i++ {
					version[major] = append(version[major], make([]string, 0))
				}
			}

			if patch > len(version[major][minor])-1 {
				for i := len(version[major][minor]) - 1; i < patch+1; i++ {
					version[major][minor] = append(version[major][minor], "")
				}

			}

			log.Trace(pkgName)
			log.Trace(major)
			log.Trace(minor)
			log.Trace(patch)

			version[major][minor][patch] = basePath + "/" + pkgName + "/" + f.Name()
			storage[pkgName] = version

		}

	}

}
