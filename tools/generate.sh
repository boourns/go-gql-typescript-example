#!/bin/bash

cd server
go run github.com/99designs/gqlgen generate
cd ..
cat server/graph/*.graphqls > js/schema/schema.graphql
cd js
yarn run graphql

