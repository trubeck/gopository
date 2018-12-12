package services

import (
	"fmt"
	"github.com/trubeck/simpleLogger"
	"strconv"
	"strings"
)

var log = simpleLogger.Create(true, "")

func GetAllPackageNames(storage map[string][][][]string) (result []string) {
	result = make([]string, 0, len(storage))

	for k := range storage {
		result = append(result, k)
	}

	return
}

func GetAllVersions(storage map[string][][][]string) (result map[string][]string) {
	packageNames := GetAllPackageNames(storage)

	result = make(map[string][]string)

	for _, pkg := range packageNames {
		for i, major := range storage[pkg] {
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

func GetLatestVersion(storage map[string][][][]string, packageName string) (major int, minor int, patch int, err error) {
	packages := GetAllVersions(storage)

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
