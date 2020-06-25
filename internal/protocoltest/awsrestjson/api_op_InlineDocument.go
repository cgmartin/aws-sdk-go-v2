// Code generated by smithy-go-codegen DO NOT EDIT.
package awsrestjson

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This example serializes an inline document as part of the payload.
func (c *Client) InlineDocument(ctx context.Context, params *InlineDocumentInput, optFns ...func(*Options)) (*InlineDocumentOutput, error) {
	stack := middleware.NewStack("InlineDocument", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	stack.Initialize.Add(awsmiddleware.RegisterServiceMetadata{
		Region:         options.Region,
		ServiceName:    "Rest Json Protocol",
		ServiceID:      "restjsonprotocol",
		EndpointPrefix: "restjsonprotocol",
		OperationName:  "InlineDocument",
	}, middleware.Before)
	stack.Build.Add(awsmiddleware.RequestInvocationIDMiddleware{}, middleware.After)
	awsmiddleware.AddResolveServiceEndpointMiddleware(stack, options)
	stack.Serialize.Add(&awsRestjson1_serializeOpInlineDocument{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpInlineDocument{}, middleware.After)
	stack.Deserialize.Add(awsmiddleware.AttemptClockSkewMiddleware{}, middleware.After)
	stack.Finalize.Add(retry.NewAttemptMiddleware(options.Retryer, smithyhttp.RequestCloner), middleware.After)
	stack.Finalize.Add(retry.MetricsHeaderMiddleware{}, middleware.After)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "InlineDocument",
			Err:           err,
		}
	}
	out := result.(*InlineDocumentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type InlineDocumentInput struct {
	StringValue   *string
	DocumentValue smithy.Document
}

type InlineDocumentOutput struct {
	StringValue   *string
	DocumentValue smithy.Document

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}
