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

// UpdateCSPMAzureTenantDefaultSubscriptionIDReader is a Reader for the UpdateCSPMAzureTenantDefaultSubscriptionID structure.
type UpdateCSPMAzureTenantDefaultSubscriptionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUpdateCSPMAzureTenantDefaultSubscriptionIDCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCSPMAzureTenantDefaultSubscriptionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateCSPMAzureTenantDefaultSubscriptionIDCreated creates a UpdateCSPMAzureTenantDefaultSubscriptionIDCreated with default headers values
func NewUpdateCSPMAzureTenantDefaultSubscriptionIDCreated() *UpdateCSPMAzureTenantDefaultSubscriptionIDCreated {
	return &UpdateCSPMAzureTenantDefaultSubscriptionIDCreated{}
}

/*
UpdateCSPMAzureTenantDefaultSubscriptionIDCreated describes a response with status code 201, with default header values.

Created
*/
type UpdateCSPMAzureTenantDefaultSubscriptionIDCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDCreated) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/default-subscription-id/v1][%d] updateCSPMAzureTenantDefaultSubscriptionIdCreated  %+v", 201, o.Payload)
}
func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDCreated) GetPayload() *models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest creates a UpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest with default headers values
func NewUpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest() *UpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest {
	return &UpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest{}
}

/*
UpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/default-subscription-id/v1][%d] updateCSPMAzureTenantDefaultSubscriptionIdBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest) GetPayload() *models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMAzureTenantDefaultSubscriptionIDForbidden creates a UpdateCSPMAzureTenantDefaultSubscriptionIDForbidden with default headers values
func NewUpdateCSPMAzureTenantDefaultSubscriptionIDForbidden() *UpdateCSPMAzureTenantDefaultSubscriptionIDForbidden {
	return &UpdateCSPMAzureTenantDefaultSubscriptionIDForbidden{}
}

/*
UpdateCSPMAzureTenantDefaultSubscriptionIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateCSPMAzureTenantDefaultSubscriptionIDForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDForbidden) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/default-subscription-id/v1][%d] updateCSPMAzureTenantDefaultSubscriptionIdForbidden  %+v", 403, o.Payload)
}
func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests creates a UpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests with default headers values
func NewUpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests() *UpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests {
	return &UpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests{}
}

/*
UpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests struct {

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

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/default-subscription-id/v1][%d] updateCSPMAzureTenantDefaultSubscriptionIdTooManyRequests  %+v", 429, o.Payload)
}
func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError creates a UpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError with default headers values
func NewUpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError() *UpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError {
	return &UpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError{}
}

/*
UpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type UpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /cloud-connect-cspm-azure/entities/default-subscription-id/v1][%d] updateCSPMAzureTenantDefaultSubscriptionIdInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError) GetPayload() *models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1 {
	return o.Payload
}

func (o *UpdateCSPMAzureTenantDefaultSubscriptionIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.RegistrationAzureTenantDefaultSubscriptionIDResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
