module github.com/pevin/internal-pos-service-api/lib

go 1.21.1

replace github.com/pevin/internal-pos-service-api/domain => ../domain

require (
	github.com/aws/aws-lambda-go v1.46.0
	github.com/pevin/internal-pos-service-api/domain v0.0.0-00010101000000-000000000000
)
