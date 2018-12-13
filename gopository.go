package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/trubeck/simpleLogger"
	"net/http"
	"os"
)

var basePath string
var host string
var port string

var sslCertPath string
var sslKeyPath string

var storage map[string][][][]string

var log = simpleLogger.Create(false, "")

func main() {

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

	storage = make(map[string][][][]string)

	// Scan for artifacts
	scanFolders()

	// Start webserver
	router := httprouter.New()
	router.GET("/download/:pkg/:version", Download)
	router.GET("/packages", ListPackages)
	router.GET("/versions", ListVersions)

	log.Info("Ready to accept connections")

	if sslCertPath == "" || sslKeyPath == "" {
		log.Info("You are serving gopository with HTTP.")
		log.Fatal(http.ListenAndServe(host+":"+port, router))
	} else {
		log.Info("You are serving gopository with HTTPS.")
		log.Fatal(http.ListenAndServeTLS(host+":"+port, sslCertPath, sslKeyPath, router))
	}
}
