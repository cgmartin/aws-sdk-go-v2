// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Submits a request to purchase an offering. If you already have an active
// reservation, you can't purchase another offering.
func (c *Client) PurchaseOffering(ctx context.Context, params *PurchaseOfferingInput, optFns ...func(*Options)) (*PurchaseOfferingOutput, error) {
	if params == nil {
		params = &PurchaseOfferingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PurchaseOffering", params, optFns, addOperationPurchaseOfferingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PurchaseOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to purchase a offering.
type PurchaseOfferingInput struct {

	// The Amazon Resource Name (ARN) of the offering.
	//
	// This member is required.
	OfferingArn *string

	// The name that you want to use for the reservation.
	//
	// This member is required.
	ReservationName *string

	// The date and time that you want the reservation to begin, in Coordinated
	// Universal Time (UTC). You can specify any date and time between 12:00am on the
	// first day of the current month to the current time on today's date, inclusive.
	// Specify the start in a 24-hour notation. Use the following format:
	// YYYY-MM-DDTHH:mm:SSZ, where T and Z are literal characters. For example, to
	// specify 11:30pm on March 5, 2020, enter 2020-03-05T23:30:00Z.
	//
	// This member is required.
	Start *string
}

type PurchaseOfferingOutput struct {

	// A pricing agreement for a discounted rate for a specific outbound bandwidth that
	// your MediaConnect account will use each month over a specific time period. The
	// discounted rate in the reservation applies to outbound bandwidth for all flows
	// from your account until your account reaches the amount of bandwidth in your
	// reservation. If you use more outbound bandwidth than the agreed upon amount in a
	// single month, the overage is charged at the on-demand rate.
	Reservation *types.Reservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPurchaseOfferingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPurchaseOffering{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPurchaseOffering{}, middleware.After)
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
	if err = addOpPurchaseOfferingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseOffering(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPurchaseOffering(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "PurchaseOffering",
	}
}