module github.com/aws/aws-sdk-go-v2/service/kinesis

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.9.2
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v0.0.0-00010101000000-000000000000
	github.com/aws/smithy-go v1.8.1-0.20211011213402-2f2385299b79
	github.com/jmespath/go-jmespath v0.4.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream => ../../aws/protocol/eventstream/
