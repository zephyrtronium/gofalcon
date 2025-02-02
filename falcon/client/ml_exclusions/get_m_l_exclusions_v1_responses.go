// Code generated by go-swagger; DO NOT EDIT.

package ml_exclusions

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

// GetMLExclusionsV1Reader is a Reader for the GetMLExclusionsV1 structure.
type GetMLExclusionsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMLExclusionsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMLExclusionsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMLExclusionsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMLExclusionsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMLExclusionsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMLExclusionsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetMLExclusionsV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMLExclusionsV1OK creates a GetMLExclusionsV1OK with default headers values
func NewGetMLExclusionsV1OK() *GetMLExclusionsV1OK {
	return &GetMLExclusionsV1OK{}
}

/*
GetMLExclusionsV1OK describes a response with status code 200, with default header values.

OK
*/
type GetMLExclusionsV1OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesMlExclusionRespV1
}

func (o *GetMLExclusionsV1OK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ml-exclusions/v1][%d] getMLExclusionsV1OK  %+v", 200, o.Payload)
}
func (o *GetMLExclusionsV1OK) GetPayload() *models.ResponsesMlExclusionRespV1 {
	return o.Payload
}

func (o *GetMLExclusionsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesMlExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMLExclusionsV1BadRequest creates a GetMLExclusionsV1BadRequest with default headers values
func NewGetMLExclusionsV1BadRequest() *GetMLExclusionsV1BadRequest {
	return &GetMLExclusionsV1BadRequest{}
}

/*
GetMLExclusionsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetMLExclusionsV1BadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesMlExclusionRespV1
}

func (o *GetMLExclusionsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ml-exclusions/v1][%d] getMLExclusionsV1BadRequest  %+v", 400, o.Payload)
}
func (o *GetMLExclusionsV1BadRequest) GetPayload() *models.ResponsesMlExclusionRespV1 {
	return o.Payload
}

func (o *GetMLExclusionsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesMlExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMLExclusionsV1Forbidden creates a GetMLExclusionsV1Forbidden with default headers values
func NewGetMLExclusionsV1Forbidden() *GetMLExclusionsV1Forbidden {
	return &GetMLExclusionsV1Forbidden{}
}

/*
GetMLExclusionsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMLExclusionsV1Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetMLExclusionsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ml-exclusions/v1][%d] getMLExclusionsV1Forbidden  %+v", 403, o.Payload)
}
func (o *GetMLExclusionsV1Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetMLExclusionsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMLExclusionsV1TooManyRequests creates a GetMLExclusionsV1TooManyRequests with default headers values
func NewGetMLExclusionsV1TooManyRequests() *GetMLExclusionsV1TooManyRequests {
	return &GetMLExclusionsV1TooManyRequests{}
}

/*
GetMLExclusionsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetMLExclusionsV1TooManyRequests struct {

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

func (o *GetMLExclusionsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ml-exclusions/v1][%d] getMLExclusionsV1TooManyRequests  %+v", 429, o.Payload)
}
func (o *GetMLExclusionsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMLExclusionsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMLExclusionsV1InternalServerError creates a GetMLExclusionsV1InternalServerError with default headers values
func NewGetMLExclusionsV1InternalServerError() *GetMLExclusionsV1InternalServerError {
	return &GetMLExclusionsV1InternalServerError{}
}

/*
GetMLExclusionsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMLExclusionsV1InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesMlExclusionRespV1
}

func (o *GetMLExclusionsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ml-exclusions/v1][%d] getMLExclusionsV1InternalServerError  %+v", 500, o.Payload)
}
func (o *GetMLExclusionsV1InternalServerError) GetPayload() *models.ResponsesMlExclusionRespV1 {
	return o.Payload
}

func (o *GetMLExclusionsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesMlExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMLExclusionsV1Default creates a GetMLExclusionsV1Default with default headers values
func NewGetMLExclusionsV1Default(code int) *GetMLExclusionsV1Default {
	return &GetMLExclusionsV1Default{
		_statusCode: code,
	}
}

/*
GetMLExclusionsV1Default describes a response with status code -1, with default header values.

OK
*/
type GetMLExclusionsV1Default struct {
	_statusCode int

	Payload *models.ResponsesMlExclusionRespV1
}

// Code gets the status code for the get m l exclusions v1 default response
func (o *GetMLExclusionsV1Default) Code() int {
	return o._statusCode
}

func (o *GetMLExclusionsV1Default) Error() string {
	return fmt.Sprintf("[GET /policy/entities/ml-exclusions/v1][%d] getMLExclusionsV1 default  %+v", o._statusCode, o.Payload)
}
func (o *GetMLExclusionsV1Default) GetPayload() *models.ResponsesMlExclusionRespV1 {
	return o.Payload
}

func (o *GetMLExclusionsV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesMlExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
