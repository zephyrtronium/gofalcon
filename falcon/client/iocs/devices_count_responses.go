// Code generated by go-swagger; DO NOT EDIT.

package iocs

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

// DevicesCountReader is a Reader for the DevicesCount structure.
type DevicesCountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DevicesCountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDevicesCountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDevicesCountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDevicesCountTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDevicesCountDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDevicesCountOK creates a DevicesCountOK with default headers values
func NewDevicesCountOK() *DevicesCountOK {
	return &DevicesCountOK{}
}

/*
DevicesCountOK describes a response with status code 200, with default header values.

OK
*/
type DevicesCountOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIMsaReplyIOCDevicesCount
}

func (o *DevicesCountOK) Error() string {
	return fmt.Sprintf("[GET /indicators/aggregates/devices-count/v1][%d] devicesCountOK  %+v", 200, o.Payload)
}
func (o *DevicesCountOK) GetPayload() *models.APIMsaReplyIOCDevicesCount {
	return o.Payload
}

func (o *DevicesCountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIMsaReplyIOCDevicesCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDevicesCountForbidden creates a DevicesCountForbidden with default headers values
func NewDevicesCountForbidden() *DevicesCountForbidden {
	return &DevicesCountForbidden{}
}

/*
DevicesCountForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DevicesCountForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *DevicesCountForbidden) Error() string {
	return fmt.Sprintf("[GET /indicators/aggregates/devices-count/v1][%d] devicesCountForbidden  %+v", 403, o.Payload)
}
func (o *DevicesCountForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DevicesCountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDevicesCountTooManyRequests creates a DevicesCountTooManyRequests with default headers values
func NewDevicesCountTooManyRequests() *DevicesCountTooManyRequests {
	return &DevicesCountTooManyRequests{}
}

/*
DevicesCountTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DevicesCountTooManyRequests struct {

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

func (o *DevicesCountTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /indicators/aggregates/devices-count/v1][%d] devicesCountTooManyRequests  %+v", 429, o.Payload)
}
func (o *DevicesCountTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DevicesCountTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDevicesCountDefault creates a DevicesCountDefault with default headers values
func NewDevicesCountDefault(code int) *DevicesCountDefault {
	return &DevicesCountDefault{
		_statusCode: code,
	}
}

/*
DevicesCountDefault describes a response with status code -1, with default header values.

OK
*/
type DevicesCountDefault struct {
	_statusCode int

	Payload *models.APIMsaReplyIOCDevicesCount
}

// Code gets the status code for the devices count default response
func (o *DevicesCountDefault) Code() int {
	return o._statusCode
}

func (o *DevicesCountDefault) Error() string {
	return fmt.Sprintf("[GET /indicators/aggregates/devices-count/v1][%d] DevicesCount default  %+v", o._statusCode, o.Payload)
}
func (o *DevicesCountDefault) GetPayload() *models.APIMsaReplyIOCDevicesCount {
	return o.Payload
}

func (o *DevicesCountDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIMsaReplyIOCDevicesCount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
