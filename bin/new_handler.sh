#!/bin/bash

# This script creates a new handler in the current directory.

SVC_NAME=$1
SVC_PATH=handlers/$SVC_NAME
HANDLER_NAME=$(echo "$SVC_NAME" | sed 's/_//g' | sed 's/-//g')
GITHUB_USER=pevin

# Code to be inserted into main.go
INIT_CODE="lambda.StartWithOptions(handler.Handle, lambda.WithContext(context.Background()))"
IMPORT_CODE="\nimport (\n\t\"context\"\n\n\t\"github.com/aws/aws-lambda-go/lambda\"\n\t\"github.com/$GITHUB_USER/internal-pos-service-api/$SVC_PATH/handler\"\n)"

# duplicate template
cp -r bin/{{handler-template}} handlers/$SVC_NAME

# insert code to main.go
sed -i '' "s/lambda handler/lambda handler\n\t$INIT_CODE/g" $SVC_PATH/main.go
sed -i '' "s|package main|package main\n$IMPORT_CODE|g" $SVC_PATH/main.go

# rename handler
mv $SVC_PATH/handler/_handlername.go $SVC_PATH/handler/$HANDLER_NAME.go

# add to go.work
sed -i  '' "s|)|\t./$SVC_PATH\n)|g" go.work

# init go
cd handlers/$SVC_NAME && go mod init github.com/$GITHUB_USER/internal-pos-service-api/handlers/$SVC_NAME && go mod tidy
