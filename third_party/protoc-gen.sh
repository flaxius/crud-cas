#!/usr/bin/env bash
protoc --proto_path=api/authentication --proto_path=third_party --go_out=plugins=grpc:pkg/pb/authentication oauth.proto
protoc --proto_path=api/authentication --proto_path=third_party --grpc-gateway_out=logtostderr=true:pkg/pb/authentication oauth.proto
protoc --proto_path=api/authentication --proto_path=third_party --swagger_out=logtostderr=true:api/authentication/swagger oauth.proto
