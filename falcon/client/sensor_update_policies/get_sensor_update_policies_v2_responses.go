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

// GetSensorUpdatePoliciesV2Reader is a Reader for the GetSensorUpdatePoliciesV2 structure.
type GetSensorUpdatePoliciesV2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSensorUpdatePoliciesV2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSensorUpdatePoliciesV2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSensorUpdatePoliciesV2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSensorUpdatePoliciesV2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSensorUpdatePoliciesV2TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSensorUpdatePoliciesV2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSensorUpdatePoliciesV2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSensorUpdatePoliciesV2OK creates a GetSensorUpdatePoliciesV2OK with default headers values
func NewGetSensorUpdatePoliciesV2OK() *GetSensorUpdatePoliciesV2OK {
	return &GetSensorUpdatePoliciesV2OK{}
}

/*
GetSensorUpdatePoliciesV2OK describes a response with status code 200, with default header values.

OK
*/
type GetSensorUpdatePoliciesV2OK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

func (o *GetSensorUpdatePoliciesV2OK) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v2][%d] getSensorUpdatePoliciesV2OK  %+v", 200, o.Payload)
}
func (o *GetSensorUpdatePoliciesV2OK) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesV2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorUpdatePoliciesV2Forbidden creates a GetSensorUpdatePoliciesV2Forbidden with default headers values
func NewGetSensorUpdatePoliciesV2Forbidden() *GetSensorUpdatePoliciesV2Forbidden {
	return &GetSensorUpdatePoliciesV2Forbidden{}
}

/*
GetSensorUpdatePoliciesV2Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSensorUpdatePoliciesV2Forbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaErrorsOnly
}

func (o *GetSensorUpdatePoliciesV2Forbidden) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v2][%d] getSensorUpdatePoliciesV2Forbidden  %+v", 403, o.Payload)
}
func (o *GetSensorUpdatePoliciesV2Forbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesV2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorUpdatePoliciesV2NotFound creates a GetSensorUpdatePoliciesV2NotFound with default headers values
func NewGetSensorUpdatePoliciesV2NotFound() *GetSensorUpdatePoliciesV2NotFound {
	return &GetSensorUpdatePoliciesV2NotFound{}
}

/*
GetSensorUpdatePoliciesV2NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetSensorUpdatePoliciesV2NotFound struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

func (o *GetSensorUpdatePoliciesV2NotFound) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v2][%d] getSensorUpdatePoliciesV2NotFound  %+v", 404, o.Payload)
}
func (o *GetSensorUpdatePoliciesV2NotFound) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesV2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorUpdatePoliciesV2TooManyRequests creates a GetSensorUpdatePoliciesV2TooManyRequests with default headers values
func NewGetSensorUpdatePoliciesV2TooManyRequests() *GetSensorUpdatePoliciesV2TooManyRequests {
	return &GetSensorUpdatePoliciesV2TooManyRequests{}
}

/*
GetSensorUpdatePoliciesV2TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetSensorUpdatePoliciesV2TooManyRequests struct {

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

func (o *GetSensorUpdatePoliciesV2TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v2][%d] getSensorUpdatePoliciesV2TooManyRequests  %+v", 429, o.Payload)
}
func (o *GetSensorUpdatePoliciesV2TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesV2TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSensorUpdatePoliciesV2InternalServerError creates a GetSensorUpdatePoliciesV2InternalServerError with default headers values
func NewGetSensorUpdatePoliciesV2InternalServerError() *GetSensorUpdatePoliciesV2InternalServerError {
	return &GetSensorUpdatePoliciesV2InternalServerError{}
}

/*
GetSensorUpdatePoliciesV2InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetSensorUpdatePoliciesV2InternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

func (o *GetSensorUpdatePoliciesV2InternalServerError) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v2][%d] getSensorUpdatePoliciesV2InternalServerError  %+v", 500, o.Payload)
}
func (o *GetSensorUpdatePoliciesV2InternalServerError) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesV2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSensorUpdatePoliciesV2Default creates a GetSensorUpdatePoliciesV2Default with default headers values
func NewGetSensorUpdatePoliciesV2Default(code int) *GetSensorUpdatePoliciesV2Default {
	return &GetSensorUpdatePoliciesV2Default{
		_statusCode: code,
	}
}

/*
GetSensorUpdatePoliciesV2Default describes a response with status code -1, with default header values.

OK
*/
type GetSensorUpdatePoliciesV2Default struct {
	_statusCode int

	Payload *models.ResponsesSensorUpdatePoliciesV2
}

// Code gets the status code for the get sensor update policies v2 default response
func (o *GetSensorUpdatePoliciesV2Default) Code() int {
	return o._statusCode
}

func (o *GetSensorUpdatePoliciesV2Default) Error() string {
	return fmt.Sprintf("[GET /policy/entities/sensor-update/v2][%d] getSensorUpdatePoliciesV2 default  %+v", o._statusCode, o.Payload)
}
func (o *GetSensorUpdatePoliciesV2Default) GetPayload() *models.ResponsesSensorUpdatePoliciesV2 {
	return o.Payload
}

func (o *GetSensorUpdatePoliciesV2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponsesSensorUpdatePoliciesV2)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
