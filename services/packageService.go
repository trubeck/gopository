package services

import (
	"fmt"

	"github.com/trubeck/gopository/storage"
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

	return ParseVersion(latestVersion, ".")
}
