// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package organizations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListAWSServiceAccessForOrganizationRequest
type ListAWSServiceAccessForOrganizationInput struct {
	_ struct{} `type:"structure"`

	// (Optional) Use this to limit the number of results you want included per
	// page in the response. If you do not include this parameter, it defaults to
	// a value that is specific to the operation. If additional items exist beyond
	// the maximum you specify, the NextToken response element is present and has
	// a value (is not null). Include that value as the NextToken request parameter
	// in the next call to the operation to get the next part of the results. Note
	// that Organizations might return fewer results than the maximum even when
	// there are more results available. You should check NextToken after every
	// operation to ensure that you receive all of the results.
	MaxResults *int64 `min:"1" type:"integer"`

	// Use this parameter if you receive a NextToken response in a previous request
	// that indicates that there is more output available. Set it to the value of
	// the previous call's NextToken response to indicate where the output should
	// continue from.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAWSServiceAccessForOrganizationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListAWSServiceAccessForOrganizationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListAWSServiceAccessForOrganizationInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListAWSServiceAccessForOrganizationResponse
type ListAWSServiceAccessForOrganizationOutput struct {
	_ struct{} `type:"structure"`

	// A list of the service principals for the services that are enabled to integrate
	// with your organization. Each principal is a structure that includes the name
	// and the date that it was enabled for integration with AWS Organizations.
	EnabledServicePrincipals []EnabledServicePrincipal `type:"list"`

	// If present, this value indicates that there is more output available than
	// is included in the current response. Use this value in the NextToken request
	// parameter in a subsequent call to the operation to get the next part of the
	// output. You should repeat this until the NextToken response element comes
	// back as null.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListAWSServiceAccessForOrganizationOutput) String() string {
	return awsutil.Prettify(s)
}

const opListAWSServiceAccessForOrganization = "ListAWSServiceAccessForOrganization"

// ListAWSServiceAccessForOrganizationRequest returns a request value for making API operation for
// AWS Organizations.
//
// Returns a list of the AWS services that you enabled to integrate with your
// organization. After a service on this list creates the resources that it
// requires for the integration, it can perform operations on your organization
// and its accounts.
//
// For more information about integrating other services with AWS Organizations,
// including the list of services that currently work with Organizations, see
// Integrating AWS Organizations with Other AWS Services (http://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services.html)
// in the AWS Organizations User Guide.
//
// This operation can be called only from the organization's master account.
//
//    // Example sending a request using ListAWSServiceAccessForOrganizationRequest.
//    req := client.ListAWSServiceAccessForOrganizationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/organizations-2016-11-28/ListAWSServiceAccessForOrganization
func (c *Client) ListAWSServiceAccessForOrganizationRequest(input *ListAWSServiceAccessForOrganizationInput) ListAWSServiceAccessForOrganizationRequest {
	op := &aws.Operation{
		Name:       opListAWSServiceAccessForOrganization,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"NextToken"},
			OutputTokens:    []string{"NextToken"},
			LimitToken:      "MaxResults",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListAWSServiceAccessForOrganizationInput{}
	}

	req := c.newRequest(op, input, &ListAWSServiceAccessForOrganizationOutput{})
	return ListAWSServiceAccessForOrganizationRequest{Request: req, Input: input, Copy: c.ListAWSServiceAccessForOrganizationRequest}
}

// ListAWSServiceAccessForOrganizationRequest is the request type for the
// ListAWSServiceAccessForOrganization API operation.
type ListAWSServiceAccessForOrganizationRequest struct {
	*aws.Request
	Input *ListAWSServiceAccessForOrganizationInput
	Copy  func(*ListAWSServiceAccessForOrganizationInput) ListAWSServiceAccessForOrganizationRequest
}

// Send marshals and sends the ListAWSServiceAccessForOrganization API request.
func (r ListAWSServiceAccessForOrganizationRequest) Send(ctx context.Context) (*ListAWSServiceAccessForOrganizationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListAWSServiceAccessForOrganizationResponse{
		ListAWSServiceAccessForOrganizationOutput: r.Request.Data.(*ListAWSServiceAccessForOrganizationOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListAWSServiceAccessForOrganizationRequestPaginator returns a paginator for ListAWSServiceAccessForOrganization.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListAWSServiceAccessForOrganizationRequest(input)
//   p := organizations.NewListAWSServiceAccessForOrganizationRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListAWSServiceAccessForOrganizationPaginator(req ListAWSServiceAccessForOrganizationRequest) ListAWSServiceAccessForOrganizationPaginator {
	return ListAWSServiceAccessForOrganizationPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListAWSServiceAccessForOrganizationInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// ListAWSServiceAccessForOrganizationPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListAWSServiceAccessForOrganizationPaginator struct {
	aws.Pager
}

func (p *ListAWSServiceAccessForOrganizationPaginator) CurrentPage() *ListAWSServiceAccessForOrganizationOutput {
	return p.Pager.CurrentPage().(*ListAWSServiceAccessForOrganizationOutput)
}

// ListAWSServiceAccessForOrganizationResponse is the response type for the
// ListAWSServiceAccessForOrganization API operation.
type ListAWSServiceAccessForOrganizationResponse struct {
	*ListAWSServiceAccessForOrganizationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListAWSServiceAccessForOrganization request.
func (r *ListAWSServiceAccessForOrganizationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}