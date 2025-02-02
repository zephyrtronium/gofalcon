// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// RevokeUserRoleIdsReader is a Reader for the RevokeUserRoleIds structure.
type RevokeUserRoleIdsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RevokeUserRoleIdsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRevokeUserRoleIdsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRevokeUserRoleIdsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRevokeUserRoleIdsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRevokeUserRoleIdsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRevokeUserRoleIdsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRevokeUserRoleIdsOK creates a RevokeUserRoleIdsOK with default headers values
func NewRevokeUserRoleIdsOK() *RevokeUserRoleIdsOK {
	return &RevokeUserRoleIdsOK{}
}

/*
RevokeUserRoleIdsOK describes a response with status code 200, with default header values.

OK
*/
type RevokeUserRoleIdsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainUserRoleIDsResponse
}

func (o *RevokeUserRoleIdsOK) Error() string {
	return fmt.Sprintf("[DELETE /user-roles/entities/user-roles/v1][%d] revokeUserRoleIdsOK  %+v", 200, o.Payload)
}
func (o *RevokeUserRoleIdsOK) GetPayload() *models.DomainUserRoleIDsResponse {
	return o.Payload
}

func (o *RevokeUserRoleIdsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainUserRoleIDsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeUserRoleIdsBadRequest creates a RevokeUserRoleIdsBadRequest with default headers values
func NewRevokeUserRoleIdsBadRequest() *RevokeUserRoleIdsBadRequest {
	return &RevokeUserRoleIdsBadRequest{}
}

/*
RevokeUserRoleIdsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RevokeUserRoleIdsBadRequest struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

func (o *RevokeUserRoleIdsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /user-roles/entities/user-roles/v1][%d] revokeUserRoleIdsBadRequest  %+v", 400, o.Payload)
}
func (o *RevokeUserRoleIdsBadRequest) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *RevokeUserRoleIdsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeUserRoleIdsForbidden creates a RevokeUserRoleIdsForbidden with default headers values
func NewRevokeUserRoleIdsForbidden() *RevokeUserRoleIdsForbidden {
	return &RevokeUserRoleIdsForbidden{}
}

/*
RevokeUserRoleIdsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RevokeUserRoleIdsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

func (o *RevokeUserRoleIdsForbidden) Error() string {
	return fmt.Sprintf("[DELETE /user-roles/entities/user-roles/v1][%d] revokeUserRoleIdsForbidden  %+v", 403, o.Payload)
}
func (o *RevokeUserRoleIdsForbidden) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *RevokeUserRoleIdsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRevokeUserRoleIdsTooManyRequests creates a RevokeUserRoleIdsTooManyRequests with default headers values
func NewRevokeUserRoleIdsTooManyRequests() *RevokeUserRoleIdsTooManyRequests {
	return &RevokeUserRoleIdsTooManyRequests{}
}

/*
RevokeUserRoleIdsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RevokeUserRoleIdsTooManyRequests struct {

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

func (o *RevokeUserRoleIdsTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /user-roles/entities/user-roles/v1][%d] revokeUserRoleIdsTooManyRequests  %+v", 429, o.Payload)
}
func (o *RevokeUserRoleIdsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RevokeUserRoleIdsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRevokeUserRoleIdsDefault creates a RevokeUserRoleIdsDefault with default headers values
func NewRevokeUserRoleIdsDefault(code int) *RevokeUserRoleIdsDefault {
	return &RevokeUserRoleIdsDefault{
		_statusCode: code,
	}
}

/*
RevokeUserRoleIdsDefault describes a response with status code -1, with default header values.

OK
*/
type RevokeUserRoleIdsDefault struct {
	_statusCode int

	Payload *models.DomainUserRoleIDsResponse
}

// Code gets the status code for the revoke user role ids default response
func (o *RevokeUserRoleIdsDefault) Code() int {
	return o._statusCode
}

func (o *RevokeUserRoleIdsDefault) Error() string {
	return fmt.Sprintf("[DELETE /user-roles/entities/user-roles/v1][%d] RevokeUserRoleIds default  %+v", o._statusCode, o.Payload)
}
func (o *RevokeUserRoleIdsDefault) GetPayload() *models.DomainUserRoleIDsResponse {
	return o.Payload
}

func (o *RevokeUserRoleIdsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainUserRoleIDsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
