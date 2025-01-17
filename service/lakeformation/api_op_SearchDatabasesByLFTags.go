// Code generated by smithy-go-codegen DO NOT EDIT.

package lakeformation

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lakeformation/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation allows a search on DATABASE resources by TagCondition. This
// operation is used by admins who want to grant user permissions on certain
// TagConditions. Before making a grant, the admin can use SearchDatabasesByTags to
// find all resources where the given TagConditions are valid to verify whether the
// returned resources can be shared.
func (c *Client) SearchDatabasesByLFTags(ctx context.Context, params *SearchDatabasesByLFTagsInput, optFns ...func(*Options)) (*SearchDatabasesByLFTagsOutput, error) {
	if params == nil {
		params = &SearchDatabasesByLFTagsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchDatabasesByLFTags", params, optFns, c.addOperationSearchDatabasesByLFTagsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchDatabasesByLFTagsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchDatabasesByLFTagsInput struct {

	// A list of conditions (LFTag structures) to search for in database resources.
	//
	// This member is required.
	Expression []types.LFTag

	// The identifier for the Data Catalog. By default, the account ID. The Data
	// Catalog is the persistent metadata store. It contains database definitions,
	// table definitions, and other control information to manage your AWS Lake
	// Formation environment.
	CatalogId *string

	// The maximum number of results to return.
	MaxResults *int32

	// A continuation token, if this is not the first call to retrieve this list.
	NextToken *string

	noSmithyDocumentSerde
}

type SearchDatabasesByLFTagsOutput struct {

	// A list of databases that meet the tag conditions.
	DatabaseList []types.TaggedDatabase

	// A continuation token, present if the current list segment is not the last.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchDatabasesByLFTagsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSearchDatabasesByLFTags{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchDatabasesByLFTags{}, middleware.After)
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
	if err = addOpSearchDatabasesByLFTagsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchDatabasesByLFTags(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSearchDatabasesByLFTags(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lakeformation",
		OperationName: "SearchDatabasesByLFTags",
	}
}
