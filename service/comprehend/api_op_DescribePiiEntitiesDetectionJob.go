// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the properties associated with a PII entities detection job. For example,
// you can use this operation to get the job status.
func (c *Client) DescribePiiEntitiesDetectionJob(ctx context.Context, params *DescribePiiEntitiesDetectionJobInput, optFns ...func(*Options)) (*DescribePiiEntitiesDetectionJobOutput, error) {
	if params == nil {
		params = &DescribePiiEntitiesDetectionJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePiiEntitiesDetectionJob", params, optFns, c.addOperationDescribePiiEntitiesDetectionJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePiiEntitiesDetectionJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePiiEntitiesDetectionJobInput struct {

	// The identifier that Amazon Comprehend generated for the job. The operation
	// returns this identifier in its response.
	//
	// This member is required.
	JobId *string
}

type DescribePiiEntitiesDetectionJobOutput struct {

	// Provides information about a PII entities detection job.
	PiiEntitiesDetectionJobProperties *types.PiiEntitiesDetectionJobProperties

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationDescribePiiEntitiesDetectionJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePiiEntitiesDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePiiEntitiesDetectionJob{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribePiiEntitiesDetectionJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePiiEntitiesDetectionJob(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDescribePiiEntitiesDetectionJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DescribePiiEntitiesDetectionJob",
	}
}