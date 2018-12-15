package services

import (
	"strconv"
	"strings"

	log "github.com/trubeck/simpleLogger"
)

func ParseVersion(version string, delimiter string) (major int, minor int, patch int, err error) {
	split := strings.Split(version, delimiter)

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

	return major, minor, patch, err
}

func ExpandStorage(major int, minor int, patch int, slice [][][]string) [][][]string {
	if major > len(slice)-1 {
		for i := len(slice) - 1; i < major+1; i++ {
			slice = append(slice, make([][]string, 0))
		}
	}

	if minor > len(slice[major])-1 {
		for i := len(slice[major]) - 1; i < minor+1; i++ {
			slice[major] = append(slice[major], make([]string, 0))
		}
	}

	if patch > len(slice[major][minor])-1 {
		for i := len(slice[major][minor]) - 1; i < patch+1; i++ {
			slice[major][minor] = append(slice[major][minor], "")
		}

	}

	return slice
}
