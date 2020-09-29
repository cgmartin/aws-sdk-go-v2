// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/ptr"
)

type ForbiddenException struct {
	Message *string

	Code *string
}

func (e *ForbiddenException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ForbiddenException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ForbiddenException) ErrorCode() string             { return "ForbiddenException" }
func (e *ForbiddenException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ForbiddenException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *ForbiddenException) HasCode() bool {
	return e.Code != nil
}
func (e *ForbiddenException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ForbiddenException) HasMessage() bool {
	return e.Message != nil
}

type InternalFailureException struct {
	Message *string

	Code *string
}

func (e *InternalFailureException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalFailureException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalFailureException) ErrorCode() string             { return "InternalFailureException" }
func (e *InternalFailureException) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }
func (e *InternalFailureException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InternalFailureException) HasMessage() bool {
	return e.Message != nil
}
func (e *InternalFailureException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InternalFailureException) HasCode() bool {
	return e.Code != nil
}

type InvalidRequestException struct {
	Message *string

	Code *string
}

func (e *InvalidRequestException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidRequestException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidRequestException) ErrorCode() string             { return "InvalidRequestException" }
func (e *InvalidRequestException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *InvalidRequestException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *InvalidRequestException) HasCode() bool {
	return e.Code != nil
}
func (e *InvalidRequestException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *InvalidRequestException) HasMessage() bool {
	return e.Message != nil
}

type PreconditionFailedException struct {
	Message *string

	Code *string
}

func (e *PreconditionFailedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PreconditionFailedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PreconditionFailedException) ErrorCode() string             { return "PreconditionFailedException" }
func (e *PreconditionFailedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *PreconditionFailedException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *PreconditionFailedException) HasMessage() bool {
	return e.Message != nil
}
func (e *PreconditionFailedException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *PreconditionFailedException) HasCode() bool {
	return e.Code != nil
}

type RangeNotSatisfiableException struct {
	Message *string

	Code *string
}

func (e *RangeNotSatisfiableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RangeNotSatisfiableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RangeNotSatisfiableException) ErrorCode() string             { return "RangeNotSatisfiableException" }
func (e *RangeNotSatisfiableException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *RangeNotSatisfiableException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *RangeNotSatisfiableException) HasMessage() bool {
	return e.Message != nil
}
func (e *RangeNotSatisfiableException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *RangeNotSatisfiableException) HasCode() bool {
	return e.Code != nil
}

type ResourceConflictException struct {
	Message *string

	Code *string
}

func (e *ResourceConflictException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceConflictException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceConflictException) ErrorCode() string             { return "ResourceConflictException" }
func (e *ResourceConflictException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceConflictException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *ResourceConflictException) HasCode() bool {
	return e.Code != nil
}
func (e *ResourceConflictException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceConflictException) HasMessage() bool {
	return e.Message != nil
}

type ResourceNotFoundException struct {
	Message *string

	Code *string
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string             { return "ResourceNotFoundException" }
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
func (e *ResourceNotFoundException) GetMessage() string {
	return ptr.ToString(e.Message)
}
func (e *ResourceNotFoundException) HasMessage() bool {
	return e.Message != nil
}
func (e *ResourceNotFoundException) GetCode() string {
	return ptr.ToString(e.Code)
}
func (e *ResourceNotFoundException) HasCode() bool {
	return e.Code != nil
}