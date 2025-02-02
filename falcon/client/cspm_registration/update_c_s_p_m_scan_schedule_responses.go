// Code generated by go-swagger; DO NOT EDIT.

package cspm_registration

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

// UpdateCSPMScanScheduleReader is a Reader for the UpdateCSPMScanSchedule structure.
type UpdateCSPMScanScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCSPMScanScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCSPMScanScheduleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCSPMScanScheduleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCSPMScanScheduleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCSPMScanScheduleTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCSPMScanScheduleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUpdateCSPMScanScheduleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCSPMScanScheduleOK creates a UpdateCSPMScanScheduleOK with default headers values
func NewUpdateCSPMScanScheduleOK() *UpdateCSPMScanScheduleOK {
	return &UpdateCSPMScanScheduleOK{}
}

/*
UpdateCSPMScanScheduleOK describes a response with status code 200, with default header values.

OK
*/
type UpdateCSPMScanScheduleOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationScanScheduleResponseV1
}

func (o *UpdateCSPMScanScheduleOK) Error() string {
	return fmt.Sprintf("[POST /settings/scan-schedule/v1][%d] updateCSPMScanScheduleOK  %+v", 200, o.Payload)
}
func (o *UpdateCSPMScanScheduleOK) GetPayload() *models.RegistrationScanScheduleResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMScanScheduleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationScanScheduleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMScanScheduleBadRequest creates a UpdateCSPMScanScheduleBadRequest with default headers values
func NewUpdateCSPMScanScheduleBadRequest() *UpdateCSPMScanScheduleBadRequest {
	return &UpdateCSPMScanScheduleBadRequest{}
}

/*
UpdateCSPMScanScheduleBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateCSPMScanScheduleBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationScanScheduleResponseV1
}

func (o *UpdateCSPMScanScheduleBadRequest) Error() string {
	return fmt.Sprintf("[POST /settings/scan-schedule/v1][%d] updateCSPMScanScheduleBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCSPMScanScheduleBadRequest) GetPayload() *models.RegistrationScanScheduleResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMScanScheduleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationScanScheduleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMScanScheduleForbidden creates a UpdateCSPMScanScheduleForbidden with default headers values
func NewUpdateCSPMScanScheduleForbidden() *UpdateCSPMScanScheduleForbidden {
	return &UpdateCSPMScanScheduleForbidden{}
}

/*
UpdateCSPMScanScheduleForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateCSPMScanScheduleForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UpdateCSPMScanScheduleForbidden) Error() string {
	return fmt.Sprintf("[POST /settings/scan-schedule/v1][%d] updateCSPMScanScheduleForbidden  %+v", 403, o.Payload)
}
func (o *UpdateCSPMScanScheduleForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCSPMScanScheduleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCSPMScanScheduleTooManyRequests creates a UpdateCSPMScanScheduleTooManyRequests with default headers values
func NewUpdateCSPMScanScheduleTooManyRequests() *UpdateCSPMScanScheduleTooManyRequests {
	return &UpdateCSPMScanScheduleTooManyRequests{}
}

/*
UpdateCSPMScanScheduleTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateCSPMScanScheduleTooManyRequests struct {

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

func (o *UpdateCSPMScanScheduleTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /settings/scan-schedule/v1][%d] updateCSPMScanScheduleTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateCSPMScanScheduleTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCSPMScanScheduleTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCSPMScanScheduleInternalServerError creates a UpdateCSPMScanScheduleInternalServerError with default headers values
func NewUpdateCSPMScanScheduleInternalServerError() *UpdateCSPMScanScheduleInternalServerError {
	return &UpdateCSPMScanScheduleInternalServerError{}
}

/*
UpdateCSPMScanScheduleInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateCSPMScanScheduleInternalServerError struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationScanScheduleResponseV1
}

func (o *UpdateCSPMScanScheduleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /settings/scan-schedule/v1][%d] updateCSPMScanScheduleInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateCSPMScanScheduleInternalServerError) GetPayload() *models.RegistrationScanScheduleResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMScanScheduleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationScanScheduleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMScanScheduleDefault creates a UpdateCSPMScanScheduleDefault with default headers values
func NewUpdateCSPMScanScheduleDefault(code int) *UpdateCSPMScanScheduleDefault {
	return &UpdateCSPMScanScheduleDefault{
		_statusCode: code,
	}
}

/*
UpdateCSPMScanScheduleDefault describes a response with status code -1, with default header values.

OK
*/
type UpdateCSPMScanScheduleDefault struct {
	_statusCode int

	Payload *models.RegistrationScanScheduleResponseV1
}

// Code gets the status code for the update c s p m scan schedule default response
func (o *UpdateCSPMScanScheduleDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCSPMScanScheduleDefault) Error() string {
	return fmt.Sprintf("[POST /settings/scan-schedule/v1][%d] UpdateCSPMScanSchedule default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateCSPMScanScheduleDefault) GetPayload() *models.RegistrationScanScheduleResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMScanScheduleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationScanScheduleResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
