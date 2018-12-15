package controller

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/trubeck/simpleLogger"

	"github.com/trubeck/gopository/services"
	"github.com/trubeck/gopository/storage"
)

const VersionNotFound = "Version not Found"

// Download is the controller for downloading artifacts
func Download(w http.ResponseWriter, _ *http.Request, ps httprouter.Params) {

	var filepath string
	var filename string
	var major int
	var minor int
	var patch int

	versions := storage.Storage[ps.ByName("pkg")]
	if len(versions) == 0 {
		http.Error(w, "Package Not Found", 404)
		return
	}

	if ps.ByName("version") == "latest" {

		err := error(nil)

		major, minor, patch, err = services.GetLatestVersion(ps.ByName("pkg"))
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

	} else {

		err := error(nil)

		major, minor, patch, err = services.ParseVersion(ps.ByName("version"), "-")
		if err != nil {
			http.Error(w, VersionNotFound, 404)
			return
		}

	}

	if major > len(versions)-1 {
		http.Error(w, VersionNotFound, 404)
		return
	}

	if minor > len(versions[major])-1 {
		http.Error(w, VersionNotFound, 404)
		return
	}

	if patch > len(versions[major][minor])-1 {
		http.Error(w, VersionNotFound, 404)
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
