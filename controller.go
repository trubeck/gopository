package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/trubeck/gopository/services"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Download is the controller for downloading artifacts
func Download(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	var filepath string
	var filename string
	var major int
	var minor int
	var patch int

	versions := storage[ps.ByName("pkg")]
	if len(versions) == 0 {
		http.Error(w, "Package Not Found", 404)
		return
	}

	if ps.ByName("version") == "latest" {

		err := error(nil)

		major, minor, patch, err = services.GetLatestVersion(storage, ps.ByName("pkg"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	} else {

		split := strings.Split(ps.ByName("version"), "-")

		err := error(nil)

		major, err = strconv.Atoi(split[0])
		if err != nil {
			http.Error(w, "Version not found", 404)
			return
		}

		minor, err = strconv.Atoi(split[1])
		if err != nil {
			http.Error(w, "Version not found", 404)
			return
		}

		patch, err = strconv.Atoi(split[2])
		if err != nil {
			http.Error(w, "Version not found", 404)
			return
		}

	}

	if major > len(versions)-1 {
		http.Error(w, "Version not found", 404)
		return
	}

	if minor > len(versions[major])-1 {
		http.Error(w, "Version not found", 404)
		return
	}

	if patch > len(versions[major][minor])-1 {
		http.Error(w, "Version not found", 404)
		return
	}

	filepath = versions[major][minor][patch]
	log.Trace(filepath)
	splitPath := strings.Split(filepath, "/")

	filename = splitPath[len(splitPath)-1]

	openfile, err := os.Open(filepath)
	defer openfile.Close() //Close after function return

	if err != nil {
		//File not found, send 404
		http.Error(w, "File not found", 404)
		return
	}

	fileHeader := make([]byte, 512)
	//Copy the headers into the fileHeader buffer
	_, err = openfile.Read(fileHeader)
	if err != nil && err.Error() != "EOF" {
		//File not found, send 404#
		log.Error(err)
		http.Error(w, "Corrupt fileheader", 500)
		return
	}
	//Get content type of file
	FileContentType := http.DetectContentType(fileHeader)

	//Get the file size
	FileStat, _ := openfile.Stat()                     //Get info from file
	FileSize := strconv.FormatInt(FileStat.Size(), 10) //Get file size as a string

	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", FileContentType)
	w.Header().Set("Content-Length", FileSize)

	_, err = openfile.Seek(0, 0)
	if err != nil {
		//File not found, send 404
		http.Error(w, "Internal Server Error", 500)
		return
	}

	_, err = io.Copy(w, openfile)
	if err != nil {
		//File not found, send 404
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func ListPackages(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	type Response struct {
		Packages []string
	}

	var resp Response

	resp.Packages = services.GetAllPackageNames(storage)

	body, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
	}

	fmt.Println(body)

	w.Write(body)
	return
}

func ListVersions(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	type Response struct {
		Packages []struct {
			PackageName map[string][]string
		}
	}

	body, err := json.Marshal(services.GetAllVersions(storage))
	if err != nil {
		log.Error(err)
	}

	w.Write(body)
	return

}
