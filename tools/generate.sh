#!/bin/bash

cd server
go run github.com/99designs/gqlgen generate
cd ..
