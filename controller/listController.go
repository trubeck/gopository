package controller

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/trubeck/gopository/services"
	log "github.com/trubeck/simpleLogger"
	"net/http"
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

	body, err := json.Marshal(services.GetAllVersions())
	if err != nil {
		log.Error(err)
	}

	w.Write(body)
	return

}
