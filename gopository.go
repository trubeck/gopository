package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	log "github.com/trubeck/simpleLogger"

	"github.com/trubeck/gopository/controller"
	"github.com/trubeck/gopository/storage"
)

var basePath string
var host string
var port string

var sslCertPath string
var sslKeyPath string

func main() {

	// Parse args
	parseArgsAndEnv()

	log.CreateLogger(true, "")
	log.Initilize()

	storage.Initialize()

	// Scan for artifacts
	scanFolders()

	// Start webserver
	router := httprouter.New()
	router.GET("/download/:pkg/:version", controller.Download)
	router.GET("/packages", controller.ListPackages)
	router.GET("/versions", controller.ListVersions)

	log.Info("Ready to accept connections")

	if sslCertPath == "" || sslKeyPath == "" {
		log.Info("You are serving gopository with HTTP.")
		log.Fatal(http.ListenAndServe(host+":"+port, router))
	} else {
		log.Info("You are serving gopository with HTTPS.")
		log.Fatal(http.ListenAndServeTLS(host+":"+port, sslCertPath, sslKeyPath, router))
	}
}

func parseArgsAndEnv() {
	// Get ENVs

	sslCertPath = os.Getenv("SSL_CERT")
	sslKeyPath = os.Getenv("SSL_KEY")

	// Parse args
	argsWithoutProg := os.Args[1:]

	for i := 0; i < len(argsWithoutProg); i++ {
		if argsWithoutProg[i] == "--path" {
			i++
			basePath = argsWithoutProg[i]
		}

		if argsWithoutProg[i] == "--host" {
			i++
			host = argsWithoutProg[i]
		}

		if argsWithoutProg[i] == "--port" {
			i++
			port = argsWithoutProg[i]
		}

		//Program arguments overwrite env values for the SSL config
		if argsWithoutProg[i] == "--sslCert" {
			i++
			sslCertPath = argsWithoutProg[i]
		}

		if argsWithoutProg[i] == "--sslKey" {
			i++
			sslKeyPath = argsWithoutProg[i]
		}
	}

	if host == "" {
		host = "localhost"
	}

	if port == "" {
		port = "8080"
	}
}
