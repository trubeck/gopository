package services

import (
	"fmt"
	"github.com/trubeck/gopository/storage"
	log "github.com/trubeck/simpleLogger"
	"strconv"
	"strings"
)

func GetAllPackageNames() (result []string) {
	result = make([]string, 0, len(storage.Storage))

	for k := range storage.Storage {
		result = append(result, k)
	}

	return
}

func GetAllVersions() (result map[string][]string) {
	packageNames := GetAllPackageNames()

	result = make(map[string][]string)

	for _, pkg := range packageNames {
		for i, major := range storage.Storage[pkg] {
			for j, minor := range major {
				for k, patch := range minor {
					if patch != "" {
						result[pkg] = append(result[pkg], fmt.Sprintf("%d.%d.%d", i, j, k))
					}
				}
			}
		}
	}

	return
}

func GetLatestVersion(packageName string) (major int, minor int, patch int, err error) {
	packages := GetAllVersions()

	packageVersions := packages[packageName]

	latestVersion := packageVersions[len(packageVersions)-1]

	split := strings.Split(latestVersion, ".")

	major, err = strconv.Atoi(split[0])
	if err != nil {
		log.Error(err)
		return 0, 0, 0, err
	}

	minor, err = strconv.Atoi(split[1])
	if err != nil {
		log.Error(err)
		return 0, 0, 0, err
	}

	patch, err = strconv.Atoi(split[2])
	if err != nil {
		log.Error(err)
		return 0, 0, 0, err
	}

	return
}
