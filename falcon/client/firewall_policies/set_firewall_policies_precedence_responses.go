// Code generated by go-swagger; DO NOT EDIT.

package firewall_policies

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

// SetFirewallPoliciesPrecedenceReader is a Reader for the SetFirewallPoliciesPrecedence structure.
type SetFirewallPoliciesPrecedenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetFirewallPoliciesPrecedenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetFirewallPoliciesPrecedenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewSetFirewallPoliciesPrecedenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSetFirewallPoliciesPrecedenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewSetFirewallPoliciesPrecedenceTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSetFirewallPoliciesPrecedenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewSetFirewallPoliciesPrecedenceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetFirewallPoliciesPrecedenceOK creates a SetFirewallPoliciesPrecedenceOK with default headers values
func NewSetFirewallPoliciesPrecedenceOK() *SetFirewallPoliciesPrecedenceOK {
	return &SetFirewallPoliciesPrecedenceOK{}
}

/*
SetFirewallPoliciesPrecedenceOK describes a response with status code 200, with default header values.

OK
*/
type SetFirewallPoliciesPrecedenceOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *SetFirewallPoliciesPrecedenceOK) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-precedence/v1][%d] setFirewallPoliciesPrecedenceOK  %+v", 200, o.Payload)
}
func (o *SetFirewallPoliciesPrecedenceOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *SetFirewallPoliciesPrecedenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetFirewallPoliciesPrecedenceBadRequest creates a SetFirewallPoliciesPrecedenceBadRequest with default headers values
func NewSetFirewallPoliciesPrecedenceBadRequest() *SetFirewallPoliciesPrecedenceBadRequest {
	return &SetFirewallPoliciesPrecedenceBadRequest{}
}

/*
SetFirewallPoliciesPrecedenceBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type SetFirewallPoliciesPrecedenceBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *SetFirewallPoliciesPrecedenceBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-precedence/v1][%d] setFirewallPoliciesPrecedenceBadRequest  %+v", 400, o.Payload)
}
func (o *SetFirewallPoliciesPrecedenceBadRequest) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *SetFirewallPoliciesPrecedenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetFirewallPoliciesPrecedenceForbidden creates a SetFirewallPoliciesPrecedenceForbidden with default headers values
func NewSetFirewallPoliciesPrecedenceForbidden() *SetFirewallPoliciesPrecedenceForbidden {
	return &SetFirewallPoliciesPrecedenceForbidden{}
}

/*
SetFirewallPoliciesPrecedenceForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type SetFirewallPoliciesPrecedenceForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *SetFirewallPoliciesPrecedenceForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-precedence/v1][%d] setFirewallPoliciesPrecedenceForbidden  %+v", 403, o.Payload)
}
func (o *SetFirewallPoliciesPrecedenceForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *SetFirewallPoliciesPrecedenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaErrorsOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetFirewallPoliciesPrecedenceTooManyRequests creates a SetFirewallPoliciesPrecedenceTooManyRequests with default headers values
func NewSetFirewallPoliciesPrecedenceTooManyRequests() *SetFirewallPoliciesPrecedenceTooManyRequests {
	return &SetFirewallPoliciesPrecedenceTooManyRequests{}
}

/*
SetFirewallPoliciesPrecedenceTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type SetFirewallPoliciesPrecedenceTooManyRequests struct {

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

func (o *SetFirewallPoliciesPrecedenceTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-precedence/v1][%d] setFirewallPoliciesPrecedenceTooManyRequests  %+v", 429, o.Payload)
}
func (o *SetFirewallPoliciesPrecedenceTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *SetFirewallPoliciesPrecedenceTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewSetFirewallPoliciesPrecedenceInternalServerError creates a SetFirewallPoliciesPrecedenceInternalServerError with default headers values
func NewSetFirewallPoliciesPrecedenceInternalServerError() *SetFirewallPoliciesPrecedenceInternalServerError {
	return &SetFirewallPoliciesPrecedenceInternalServerError{}
}

/*
SetFirewallPoliciesPrecedenceInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type SetFirewallPoliciesPrecedenceInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *SetFirewallPoliciesPrecedenceInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-precedence/v1][%d] setFirewallPoliciesPrecedenceInternalServerError  %+v", 500, o.Payload)
}
func (o *SetFirewallPoliciesPrecedenceInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *SetFirewallPoliciesPrecedenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetFirewallPoliciesPrecedenceDefault creates a SetFirewallPoliciesPrecedenceDefault with default headers values
func NewSetFirewallPoliciesPrecedenceDefault(code int) *SetFirewallPoliciesPrecedenceDefault {
	return &SetFirewallPoliciesPrecedenceDefault{
		_statusCode: code,
	}
}

/*
SetFirewallPoliciesPrecedenceDefault describes a response with status code -1, with default header values.

OK
*/
type SetFirewallPoliciesPrecedenceDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the set firewall policies precedence default response
func (o *SetFirewallPoliciesPrecedenceDefault) Code() int {
	return o._statusCode
}

func (o *SetFirewallPoliciesPrecedenceDefault) Error() string {
	return fmt.Sprintf("[POST /policy/entities/firewall-precedence/v1][%d] setFirewallPoliciesPrecedence default  %+v", o._statusCode, o.Payload)
}
func (o *SetFirewallPoliciesPrecedenceDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *SetFirewallPoliciesPrecedenceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
