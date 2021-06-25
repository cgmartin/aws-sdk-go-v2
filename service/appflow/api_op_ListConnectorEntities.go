// Code generated by smithy-go-codegen DO NOT EDIT.

package appflow

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appflow/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the list of available connector entities supported by Amazon AppFlow.
// For example, you can query Salesforce for Account and Opportunity entities, or
// query ServiceNow for the Incident entity.
func (c *Client) ListConnectorEntities(ctx context.Context, params *ListConnectorEntitiesInput, optFns ...func(*Options)) (*ListConnectorEntitiesOutput, error) {
	if params == nil {
		params = &ListConnectorEntitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListConnectorEntities", params, optFns, c.addOperationListConnectorEntitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListConnectorEntitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListConnectorEntitiesInput struct {

	// The name of the connector profile. The name is unique for each ConnectorProfile
	// in the AWS account, and is used to query the downstream connector.
	ConnectorProfileName *string

	// The type of connector, such as Salesforce, Amplitude, and so on.
	ConnectorType types.ConnectorType

	// This optional parameter is specific to connector implementation. Some connectors
	// support multiple levels or categories of entities. You can find out the list of
	// roots for such providers by sending a request without the entitiesPath
	// parameter. If the connector supports entities at different roots, this initial
	// request returns the list of roots. Otherwise, this request returns all entities
	// supported by the provider.
	EntitiesPath *string
}

type ListConnectorEntitiesOutput struct {

	// The response of ListConnectorEntities lists entities grouped by category. This
	// map's key represents the group name, and its value contains the list of entities
	// belonging to that group.
	//
	// This member is required.
	ConnectorEntityMap map[string][]types.ConnectorEntity

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func (c *Client) addOperationListConnectorEntitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListConnectorEntities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListConnectorEntities{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListConnectorEntities(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListConnectorEntities(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appflow",
		OperationName: "ListConnectorEntities",
	}
}