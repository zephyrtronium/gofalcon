// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// UpdateRuleGroupMixin0Reader is a Reader for the UpdateRuleGroupMixin0 structure.
type UpdateRuleGroupMixin0Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRuleGroupMixin0Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRuleGroupMixin0OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRuleGroupMixin0Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRuleGroupMixin0NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateRuleGroupMixin0TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateRuleGroupMixin0Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRuleGroupMixin0OK creates a UpdateRuleGroupMixin0OK with default headers values
func NewUpdateRuleGroupMixin0OK() *UpdateRuleGroupMixin0OK {
	return &UpdateRuleGroupMixin0OK{}
}

/*
UpdateRuleGroupMixin0OK describes a response with status code 200, with default header values.

OK
*/
type UpdateRuleGroupMixin0OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIRuleGroupsResponse
}

func (o *UpdateRuleGroupMixin0OK) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rule-groups/v1][%d] updateRuleGroupMixin0OK  %+v", 200, o.Payload)
}
func (o *UpdateRuleGroupMixin0OK) GetPayload() *models.APIRuleGroupsResponse {
	return o.Payload
}

func (o *UpdateRuleGroupMixin0OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.APIRuleGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupMixin0Forbidden creates a UpdateRuleGroupMixin0Forbidden with default headers values
func NewUpdateRuleGroupMixin0Forbidden() *UpdateRuleGroupMixin0Forbidden {
	return &UpdateRuleGroupMixin0Forbidden{}
}

/*
UpdateRuleGroupMixin0Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRuleGroupMixin0Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UpdateRuleGroupMixin0Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rule-groups/v1][%d] updateRuleGroupMixin0Forbidden  %+v", 403, o.Payload)
}
func (o *UpdateRuleGroupMixin0Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupMixin0Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupMixin0NotFound creates a UpdateRuleGroupMixin0NotFound with default headers values
func NewUpdateRuleGroupMixin0NotFound() *UpdateRuleGroupMixin0NotFound {
	return &UpdateRuleGroupMixin0NotFound{}
}

/*
UpdateRuleGroupMixin0NotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateRuleGroupMixin0NotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UpdateRuleGroupMixin0NotFound) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rule-groups/v1][%d] updateRuleGroupMixin0NotFound  %+v", 404, o.Payload)
}
func (o *UpdateRuleGroupMixin0NotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupMixin0NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupMixin0TooManyRequests creates a UpdateRuleGroupMixin0TooManyRequests with default headers values
func NewUpdateRuleGroupMixin0TooManyRequests() *UpdateRuleGroupMixin0TooManyRequests {
	return &UpdateRuleGroupMixin0TooManyRequests{}
}

/*
UpdateRuleGroupMixin0TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateRuleGroupMixin0TooManyRequests struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UpdateRuleGroupMixin0TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rule-groups/v1][%d] updateRuleGroupMixin0TooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateRuleGroupMixin0TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRuleGroupMixin0TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRuleGroupMixin0Default creates a UpdateRuleGroupMixin0Default with default headers values
func NewUpdateRuleGroupMixin0Default(code int) *UpdateRuleGroupMixin0Default {
	return &UpdateRuleGroupMixin0Default{
		_statusCode: code,
	}
}

/*
UpdateRuleGroupMixin0Default describes a response with status code -1, with default header values.

OK
*/
type UpdateRuleGroupMixin0Default struct {
	_statusCode int

	Payload *models.APIRuleGroupsResponse
}

// Code gets the status code for the update rule group mixin0 default response
func (o *UpdateRuleGroupMixin0Default) Code() int {
	return o._statusCode
}

func (o *UpdateRuleGroupMixin0Default) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rule-groups/v1][%d] update-rule-groupMixin0 default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateRuleGroupMixin0Default) GetPayload() *models.APIRuleGroupsResponse {
	return o.Payload
}

func (o *UpdateRuleGroupMixin0Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIRuleGroupsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
