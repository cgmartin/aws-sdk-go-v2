// Code generated by smithy-go-codegen DO NOT EDIT.

package acmpca

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the resource-based policy attached to a private CA. If either the
// private CA resource or the policy cannot be found, this action returns a
// ResourceNotFoundException. The policy can be attached or updated with PutPolicy
// (https://docs.aws.amazon.com/acm-pca/latest/APIReference/API_PutPolicy.html) and
// removed with DeletePolicy. About Policies
//
// * A policy grants access on a private
// CA to an AWS customer account, to AWS Organizations, or to an AWS Organizations
// unit. Policies are under the control of a CA administrator. For more
// information, see Using a Resource Based Policy with ACM Private CA.
//
// * A policy
// permits a user of AWS Certificate Manager (ACM) to issue ACM certificates signed
// by a CA in another account.
//
// * For ACM to manage automatic renewal of these
// certificates, the ACM user must configure a Service Linked Role (SLR). The SLR
// allows the ACM service to assume the identity of the user, subject to
// confirmation against the ACM Private CA policy. For more information, see Using
// a Service Linked Role with ACM
// (https://docs.aws.amazon.com/acm/latest/userguide/acm-slr.html).
//
// * Updates made
// in AWS Resource Manager (RAM) are reflected in policies. For more information,
// see Using AWS Resource Access Manager (RAM) with ACM Private CA.
func (c *Client) GetPolicy(ctx context.Context, params *GetPolicyInput, optFns ...func(*Options)) (*GetPolicyOutput, error) {
	if params == nil {
		params = &GetPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPolicy", params, optFns, addOperationGetPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetPolicyInput struct {

	// The Amazon Resource Number (ARN) of the private CA that will have its policy
	// retrieved. You can find the CA's ARN by calling the ListCertificateAuthorities
	// action.
	//
	// This member is required.
	ResourceArn *string
}

type GetPolicyOutput struct {

	// The policy attached to the private CA as a JSON document.
	Policy *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPolicy{}, middleware.After)
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
	if err = addOpGetPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "acm-pca",
		OperationName: "GetPolicy",
	}
}