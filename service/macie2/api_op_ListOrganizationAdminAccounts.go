// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the delegated Amazon Macie administrator account for
// an AWS organization.
func (c *Client) ListOrganizationAdminAccounts(ctx context.Context, params *ListOrganizationAdminAccountsInput, optFns ...func(*Options)) (*ListOrganizationAdminAccountsOutput, error) {
	if params == nil {
		params = &ListOrganizationAdminAccountsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOrganizationAdminAccounts", params, optFns, addOperationListOrganizationAdminAccountsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOrganizationAdminAccountsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOrganizationAdminAccountsInput struct {

	// The maximum number of items to include in each page of a paginated response.
	MaxResults int32

	// The nextToken string that specifies which page of results to return in a
	// paginated response.
	NextToken *string
}

type ListOrganizationAdminAccountsOutput struct {

	// An array of objects, one for each delegated Amazon Macie administrator account
	// for the organization. Only one of these accounts can have a status of ENABLED.
	AdminAccounts []types.AdminAccount

	// The string to use in a subsequent request to get the next page of results in a
	// paginated response. This value is null if there are no additional pages.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListOrganizationAdminAccountsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListOrganizationAdminAccounts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListOrganizationAdminAccounts{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOrganizationAdminAccounts(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOrganizationAdminAccounts(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "ListOrganizationAdminAccounts",
	}
}