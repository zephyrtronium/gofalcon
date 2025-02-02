// Code generated by go-swagger; DO NOT EDIT.

package intel

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

// QueryIntelRuleIdsReader is a Reader for the QueryIntelRuleIds structure.
type QueryIntelRuleIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryIntelRuleIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryIntelRuleIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryIntelRuleIdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewQueryIntelRuleIdsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryIntelRuleIdsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryIntelRuleIdsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewQueryIntelRuleIdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryIntelRuleIdsOK creates a QueryIntelRuleIdsOK with default headers values
func NewQueryIntelRuleIdsOK() *QueryIntelRuleIdsOK {
	return &QueryIntelRuleIdsOK{}
}

/*
QueryIntelRuleIdsOK describes a response with status code 200, with default header values.

OK
*/
type QueryIntelRuleIdsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *QueryIntelRuleIdsOK) Error() string {
	return fmt.Sprintf("[GET /intel/queries/rules/v1][%d] queryIntelRuleIdsOK  %+v", 200, o.Payload)
}
func (o *QueryIntelRuleIdsOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryIntelRuleIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelRuleIdsBadRequest creates a QueryIntelRuleIdsBadRequest with default headers values
func NewQueryIntelRuleIdsBadRequest() *QueryIntelRuleIdsBadRequest {
	return &QueryIntelRuleIdsBadRequest{}
}

/*
QueryIntelRuleIdsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type QueryIntelRuleIdsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryIntelRuleIdsBadRequest) Error() string {
	return fmt.Sprintf("[GET /intel/queries/rules/v1][%d] queryIntelRuleIdsBadRequest  %+v", 400, o.Payload)
}
func (o *QueryIntelRuleIdsBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryIntelRuleIdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelRuleIdsForbidden creates a QueryIntelRuleIdsForbidden with default headers values
func NewQueryIntelRuleIdsForbidden() *QueryIntelRuleIdsForbidden {
	return &QueryIntelRuleIdsForbidden{}
}

/*
QueryIntelRuleIdsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryIntelRuleIdsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *QueryIntelRuleIdsForbidden) Error() string {
	return fmt.Sprintf("[GET /intel/queries/rules/v1][%d] queryIntelRuleIdsForbidden  %+v", 403, o.Payload)
}
func (o *QueryIntelRuleIdsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIntelRuleIdsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelRuleIdsTooManyRequests creates a QueryIntelRuleIdsTooManyRequests with default headers values
func NewQueryIntelRuleIdsTooManyRequests() *QueryIntelRuleIdsTooManyRequests {
	return &QueryIntelRuleIdsTooManyRequests{}
}

/*
QueryIntelRuleIdsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryIntelRuleIdsTooManyRequests struct {

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

func (o *QueryIntelRuleIdsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /intel/queries/rules/v1][%d] queryIntelRuleIdsTooManyRequests  %+v", 429, o.Payload)
}
func (o *QueryIntelRuleIdsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIntelRuleIdsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelRuleIdsInternalServerError creates a QueryIntelRuleIdsInternalServerError with default headers values
func NewQueryIntelRuleIdsInternalServerError() *QueryIntelRuleIdsInternalServerError {
	return &QueryIntelRuleIdsInternalServerError{}
}

/*
QueryIntelRuleIdsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type QueryIntelRuleIdsInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *QueryIntelRuleIdsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /intel/queries/rules/v1][%d] queryIntelRuleIdsInternalServerError  %+v", 500, o.Payload)
}
func (o *QueryIntelRuleIdsInternalServerError) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *QueryIntelRuleIdsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIntelRuleIdsDefault creates a QueryIntelRuleIdsDefault with default headers values
func NewQueryIntelRuleIdsDefault(code int) *QueryIntelRuleIdsDefault {
	return &QueryIntelRuleIdsDefault{
		_statusCode: code,
	}
}

/*
QueryIntelRuleIdsDefault describes a response with status code -1, with default header values.

OK
*/
type QueryIntelRuleIdsDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the query intel rule ids default response
func (o *QueryIntelRuleIdsDefault) Code() int {
	return o._statusCode
}

func (o *QueryIntelRuleIdsDefault) Error() string {
	return fmt.Sprintf("[GET /intel/queries/rules/v1][%d] QueryIntelRuleIds default  %+v", o._statusCode, o.Payload)
}
func (o *QueryIntelRuleIdsDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryIntelRuleIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
