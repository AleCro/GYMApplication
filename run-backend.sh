#!/bin/bash
export $(grep -v '^#' .env | xargs)
cd api
go run main.go