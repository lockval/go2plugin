#!/bin/bash

go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest

protoc --gogofaster_out=. api.proto