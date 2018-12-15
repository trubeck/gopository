package controller

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	log "github.com/trubeck/simpleLogger"

	"github.com/trubeck/gopository/services"
)

func ListPackages(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	type Response struct {
		Packages []string
	}

	var resp Response

	resp.Packages = services.GetAllPackageNames()

	body, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
	}

	w.Write(body)
	return
}

func ListVersions(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	type Response struct {
		Packages []struct {
			PackageName map[string][]string
		}
	}

	body, err := json.Marshal(services.GetAllVersions())
	if err != nil {
		log.Error(err)
	}

	w.Write(body)
	return

}
