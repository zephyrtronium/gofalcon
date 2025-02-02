// Code generated by go-swagger; DO NOT EDIT.

package sensor_update_policies

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

// DeleteSensorUpdatePoliciesReader is a Reader for the DeleteSensorUpdatePolicies structure.
type DeleteSensorUpdatePoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSensorUpdatePoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSensorUpdatePoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteSensorUpdatePoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSensorUpdatePoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteSensorUpdatePoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSensorUpdatePoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteSensorUpdatePoliciesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSensorUpdatePoliciesOK creates a DeleteSensorUpdatePoliciesOK with default headers values
func NewDeleteSensorUpdatePoliciesOK() *DeleteSensorUpdatePoliciesOK {
	return &DeleteSensorUpdatePoliciesOK{}
}

/*
DeleteSensorUpdatePoliciesOK describes a response with status code 200, with default header values.

OK
*/
type DeleteSensorUpdatePoliciesOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *DeleteSensorUpdatePoliciesOK) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sensor-update/v1][%d] deleteSensorUpdatePoliciesOK  %+v", 200, o.Payload)
}
func (o *DeleteSensorUpdatePoliciesOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSensorUpdatePoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSensorUpdatePoliciesForbidden creates a DeleteSensorUpdatePoliciesForbidden with default headers values
func NewDeleteSensorUpdatePoliciesForbidden() *DeleteSensorUpdatePoliciesForbidden {
	return &DeleteSensorUpdatePoliciesForbidden{}
}

/*
DeleteSensorUpdatePoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteSensorUpdatePoliciesForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *DeleteSensorUpdatePoliciesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sensor-update/v1][%d] deleteSensorUpdatePoliciesForbidden  %+v", 403, o.Payload)
}
func (o *DeleteSensorUpdatePoliciesForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *DeleteSensorUpdatePoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSensorUpdatePoliciesNotFound creates a DeleteSensorUpdatePoliciesNotFound with default headers values
func NewDeleteSensorUpdatePoliciesNotFound() *DeleteSensorUpdatePoliciesNotFound {
	return &DeleteSensorUpdatePoliciesNotFound{}
}

/*
DeleteSensorUpdatePoliciesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteSensorUpdatePoliciesNotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *DeleteSensorUpdatePoliciesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sensor-update/v1][%d] deleteSensorUpdatePoliciesNotFound  %+v", 404, o.Payload)
}
func (o *DeleteSensorUpdatePoliciesNotFound) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSensorUpdatePoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSensorUpdatePoliciesTooManyRequests creates a DeleteSensorUpdatePoliciesTooManyRequests with default headers values
func NewDeleteSensorUpdatePoliciesTooManyRequests() *DeleteSensorUpdatePoliciesTooManyRequests {
	return &DeleteSensorUpdatePoliciesTooManyRequests{}
}

/*
DeleteSensorUpdatePoliciesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteSensorUpdatePoliciesTooManyRequests struct {

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

func (o *DeleteSensorUpdatePoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sensor-update/v1][%d] deleteSensorUpdatePoliciesTooManyRequests  %+v", 429, o.Payload)
}
func (o *DeleteSensorUpdatePoliciesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteSensorUpdatePoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSensorUpdatePoliciesInternalServerError creates a DeleteSensorUpdatePoliciesInternalServerError with default headers values
func NewDeleteSensorUpdatePoliciesInternalServerError() *DeleteSensorUpdatePoliciesInternalServerError {
	return &DeleteSensorUpdatePoliciesInternalServerError{}
}

/*
DeleteSensorUpdatePoliciesInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteSensorUpdatePoliciesInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

func (o *DeleteSensorUpdatePoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sensor-update/v1][%d] deleteSensorUpdatePoliciesInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteSensorUpdatePoliciesInternalServerError) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSensorUpdatePoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteSensorUpdatePoliciesDefault creates a DeleteSensorUpdatePoliciesDefault with default headers values
func NewDeleteSensorUpdatePoliciesDefault(code int) *DeleteSensorUpdatePoliciesDefault {
	return &DeleteSensorUpdatePoliciesDefault{
		_statusCode: code,
	}
}

/*
DeleteSensorUpdatePoliciesDefault describes a response with status code -1, with default header values.

OK
*/
type DeleteSensorUpdatePoliciesDefault struct {
	_statusCode int

	Payload *models.MsaQueryResponse
}

// Code gets the status code for the delete sensor update policies default response
func (o *DeleteSensorUpdatePoliciesDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSensorUpdatePoliciesDefault) Error() string {
	return fmt.Sprintf("[DELETE /policy/entities/sensor-update/v1][%d] deleteSensorUpdatePolicies default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteSensorUpdatePoliciesDefault) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *DeleteSensorUpdatePoliciesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
