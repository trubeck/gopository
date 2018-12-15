#!/bin/bash

GOOS=darwin GOARCH=amd64 go install github.com/trubeck/gopository

GOOS=windows GOARCH=amd64 go install github.com/trubeck/gopository

GOOS=linux GOARCH=amd64 go install github.com/trubeck/gopository