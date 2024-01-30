#!/bin/bash

STAGE=${STAGE:-test}
OUTPUT_BUCKET=sam-artifacts
STACK_NAME=internal-pos-service-api-$STAGE
AWS_PROFILE=internal-pos-service-api

sam package --profile $AWS_PROFILE --template-file template.yaml --output-template-file packaged.template.yaml --s3-bucket $OUTPUT_BUCKET
sam deploy --profile $AWS_PROFILE --template-file packaged.template.yaml --stack-name $STACK_NAME --capabilities CAPABILITY_AUTO_EXPAND CAPABILITY_IAM CAPABILITY_NAMED_IAM --parameter-overrides Stage=$STAGE --s3-bucket $OUTPUT_BUCKET
