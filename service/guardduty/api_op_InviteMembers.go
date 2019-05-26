// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// InviteMembers request body.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/InviteMembersRequest
type InviteMembersInput struct {
	_ struct{} `type:"structure"`

	// A list of account IDs of the accounts that you want to invite to GuardDuty
	// as members.
	//
	// AccountIds is a required field
	AccountIds []string `locationName:"accountIds" type:"list" required:"true"`

	// DetectorId is a required field
	DetectorId *string `location:"uri" locationName:"detectorId" type:"string" required:"true"`

	// A boolean value that specifies whether you want to disable email notification
	// to the accounts that you’re inviting to GuardDuty as members.
	DisableEmailNotification *bool `locationName:"disableEmailNotification" type:"boolean"`

	// The invitation message that you want to send to the accounts that you’re
	// inviting to GuardDuty as members.
	Message *string `locationName:"message" type:"string"`
}

// String returns the string representation
func (s InviteMembersInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *InviteMembersInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "InviteMembersInput"}

	if s.AccountIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccountIds"))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InviteMembersInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if len(s.AccountIds) > 0 {
		v := s.AccountIds

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "accountIds", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	if s.DisableEmailNotification != nil {
		v := *s.DisableEmailNotification

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "disableEmailNotification", protocol.BoolValue(v), metadata)
	}
	if s.Message != nil {
		v := *s.Message

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "message", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DetectorId != nil {
		v := *s.DetectorId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "detectorId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// InviteMembers response object.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/InviteMembersResponse
type InviteMembersOutput struct {
	_ struct{} `type:"structure"`

	// A list of objects containing the unprocessed account and a result string
	// explaining why it was unprocessed.
	UnprocessedAccounts []UnprocessedAccount `locationName:"unprocessedAccounts" type:"list"`
}

// String returns the string representation
func (s InviteMembersOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s InviteMembersOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.UnprocessedAccounts) > 0 {
		v := s.UnprocessedAccounts

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "unprocessedAccounts", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opInviteMembers = "InviteMembers"

// InviteMembersRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Invites other AWS accounts (created as members of the current AWS account
// by CreateMembers) to enable GuardDuty and allow the current AWS account to
// view and manage these accounts' GuardDuty findings on their behalf as the
// master account.
//
//    // Example sending a request using InviteMembersRequest.
//    req := client.InviteMembersRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/InviteMembers
func (c *Client) InviteMembersRequest(input *InviteMembersInput) InviteMembersRequest {
	op := &aws.Operation{
		Name:       opInviteMembers,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/member/invite",
	}

	if input == nil {
		input = &InviteMembersInput{}
	}

	req := c.newRequest(op, input, &InviteMembersOutput{})
	return InviteMembersRequest{Request: req, Input: input, Copy: c.InviteMembersRequest}
}

// InviteMembersRequest is the request type for the
// InviteMembers API operation.
type InviteMembersRequest struct {
	*aws.Request
	Input *InviteMembersInput
	Copy  func(*InviteMembersInput) InviteMembersRequest
}

// Send marshals and sends the InviteMembers API request.
func (r InviteMembersRequest) Send(ctx context.Context) (*InviteMembersResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &InviteMembersResponse{
		InviteMembersOutput: r.Request.Data.(*InviteMembersOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// InviteMembersResponse is the response type for the
// InviteMembers API operation.
type InviteMembersResponse struct {
	*InviteMembersOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// InviteMembers request.
func (r *InviteMembersResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}