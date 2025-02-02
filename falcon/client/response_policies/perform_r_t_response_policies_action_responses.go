// Code generated by go-swagger; DO NOT EDIT.

package response_policies

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

// PerformRTResponsePoliciesActionReader is a Reader for the PerformRTResponsePoliciesAction structure.
type PerformRTResponsePoliciesActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformRTResponsePoliciesActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformRTResponsePoliciesActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPerformRTResponsePoliciesActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPerformRTResponsePoliciesActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPerformRTResponsePoliciesActionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPerformRTResponsePoliciesActionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPerformRTResponsePoliciesActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewPerformRTResponsePoliciesActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformRTResponsePoliciesActionOK creates a PerformRTResponsePoliciesActionOK with default headers values
func NewPerformRTResponsePoliciesActionOK() *PerformRTResponsePoliciesActionOK {
	return &PerformRTResponsePoliciesActionOK{}
}

/*
PerformRTResponsePoliciesActionOK describes a response with status code 200, with default header values.

OK
*/
type PerformRTResponsePoliciesActionOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

func (o *PerformRTResponsePoliciesActionOK) Error() string {
	return fmt.Sprintf("[POST /policy/entities/response-actions/v1][%d] performRTResponsePoliciesActionOK  %+v", 200, o.Payload)
}
func (o *PerformRTResponsePoliciesActionOK) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *PerformRTResponsePoliciesActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformRTResponsePoliciesActionBadRequest creates a PerformRTResponsePoliciesActionBadRequest with default headers values
func NewPerformRTResponsePoliciesActionBadRequest() *PerformRTResponsePoliciesActionBadRequest {
	return &PerformRTResponsePoliciesActionBadRequest{}
}

/*
PerformRTResponsePoliciesActionBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PerformRTResponsePoliciesActionBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

func (o *PerformRTResponsePoliciesActionBadRequest) Error() string {
	return fmt.Sprintf("[POST /policy/entities/response-actions/v1][%d] performRTResponsePoliciesActionBadRequest  %+v", 400, o.Payload)
}
func (o *PerformRTResponsePoliciesActionBadRequest) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *PerformRTResponsePoliciesActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformRTResponsePoliciesActionForbidden creates a PerformRTResponsePoliciesActionForbidden with default headers values
func NewPerformRTResponsePoliciesActionForbidden() *PerformRTResponsePoliciesActionForbidden {
	return &PerformRTResponsePoliciesActionForbidden{}
}

/*
PerformRTResponsePoliciesActionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PerformRTResponsePoliciesActionForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *PerformRTResponsePoliciesActionForbidden) Error() string {
	return fmt.Sprintf("[POST /policy/entities/response-actions/v1][%d] performRTResponsePoliciesActionForbidden  %+v", 403, o.Payload)
}
func (o *PerformRTResponsePoliciesActionForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *PerformRTResponsePoliciesActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewPerformRTResponsePoliciesActionNotFound creates a PerformRTResponsePoliciesActionNotFound with default headers values
func NewPerformRTResponsePoliciesActionNotFound() *PerformRTResponsePoliciesActionNotFound {
	return &PerformRTResponsePoliciesActionNotFound{}
}

/*
PerformRTResponsePoliciesActionNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PerformRTResponsePoliciesActionNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

func (o *PerformRTResponsePoliciesActionNotFound) Error() string {
	return fmt.Sprintf("[POST /policy/entities/response-actions/v1][%d] performRTResponsePoliciesActionNotFound  %+v", 404, o.Payload)
}
func (o *PerformRTResponsePoliciesActionNotFound) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *PerformRTResponsePoliciesActionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformRTResponsePoliciesActionTooManyRequests creates a PerformRTResponsePoliciesActionTooManyRequests with default headers values
func NewPerformRTResponsePoliciesActionTooManyRequests() *PerformRTResponsePoliciesActionTooManyRequests {
	return &PerformRTResponsePoliciesActionTooManyRequests{}
}

/*
PerformRTResponsePoliciesActionTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type PerformRTResponsePoliciesActionTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

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

func (o *PerformRTResponsePoliciesActionTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /policy/entities/response-actions/v1][%d] performRTResponsePoliciesActionTooManyRequests  %+v", 429, o.Payload)
}
func (o *PerformRTResponsePoliciesActionTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *PerformRTResponsePoliciesActionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

// NewPerformRTResponsePoliciesActionInternalServerError creates a PerformRTResponsePoliciesActionInternalServerError with default headers values
func NewPerformRTResponsePoliciesActionInternalServerError() *PerformRTResponsePoliciesActionInternalServerError {
	return &PerformRTResponsePoliciesActionInternalServerError{}
}

/*
PerformRTResponsePoliciesActionInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PerformRTResponsePoliciesActionInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesRTResponsePoliciesV1
}

func (o *PerformRTResponsePoliciesActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /policy/entities/response-actions/v1][%d] performRTResponsePoliciesActionInternalServerError  %+v", 500, o.Payload)
}
func (o *PerformRTResponsePoliciesActionInternalServerError) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *PerformRTResponsePoliciesActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

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

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformRTResponsePoliciesActionDefault creates a PerformRTResponsePoliciesActionDefault with default headers values
func NewPerformRTResponsePoliciesActionDefault(code int) *PerformRTResponsePoliciesActionDefault {
	return &PerformRTResponsePoliciesActionDefault{
		_statusCode: code,
	}
}

/*
PerformRTResponsePoliciesActionDefault describes a response with status code -1, with default header values.

OK
*/
type PerformRTResponsePoliciesActionDefault struct {
	_statusCode int

	Payload *models.ResponsesRTResponsePoliciesV1
}

// Code gets the status code for the perform r t response policies action default response
func (o *PerformRTResponsePoliciesActionDefault) Code() int {
	return o._statusCode
}

func (o *PerformRTResponsePoliciesActionDefault) Error() string {
	return fmt.Sprintf("[POST /policy/entities/response-actions/v1][%d] performRTResponsePoliciesAction default  %+v", o._statusCode, o.Payload)
}
func (o *PerformRTResponsePoliciesActionDefault) GetPayload() *models.ResponsesRTResponsePoliciesV1 {
	return o.Payload
}

func (o *PerformRTResponsePoliciesActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesRTResponsePoliciesV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
