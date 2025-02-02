// Code generated by go-swagger; DO NOT EDIT.

package mssp

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

// GetCIDGroupMembersByReader is a Reader for the GetCIDGroupMembersBy structure.
type GetCIDGroupMembersByReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCIDGroupMembersByReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCIDGroupMembersByOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewGetCIDGroupMembersByMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCIDGroupMembersByBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetCIDGroupMembersByForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetCIDGroupMembersByTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetCIDGroupMembersByDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCIDGroupMembersByOK creates a GetCIDGroupMembersByOK with default headers values
func NewGetCIDGroupMembersByOK() *GetCIDGroupMembersByOK {
	return &GetCIDGroupMembersByOK{}
}

/*
GetCIDGroupMembersByOK describes a response with status code 200, with default header values.

OK
*/
type GetCIDGroupMembersByOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

func (o *GetCIDGroupMembersByOK) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByOK  %+v", 200, o.Payload)
}
func (o *GetCIDGroupMembersByOK) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupMembersByOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupMembersByMultiStatus creates a GetCIDGroupMembersByMultiStatus with default headers values
func NewGetCIDGroupMembersByMultiStatus() *GetCIDGroupMembersByMultiStatus {
	return &GetCIDGroupMembersByMultiStatus{}
}

/*
GetCIDGroupMembersByMultiStatus describes a response with status code 207, with default header values.

Multi-Status
*/
type GetCIDGroupMembersByMultiStatus struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainCIDGroupMembersResponseV1
}

func (o *GetCIDGroupMembersByMultiStatus) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByMultiStatus  %+v", 207, o.Payload)
}
func (o *GetCIDGroupMembersByMultiStatus) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupMembersByMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCIDGroupMembersByBadRequest creates a GetCIDGroupMembersByBadRequest with default headers values
func NewGetCIDGroupMembersByBadRequest() *GetCIDGroupMembersByBadRequest {
	return &GetCIDGroupMembersByBadRequest{}
}

/*
GetCIDGroupMembersByBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetCIDGroupMembersByBadRequest struct {

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

func (o *GetCIDGroupMembersByBadRequest) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByBadRequest  %+v", 400, o.Payload)
}
func (o *GetCIDGroupMembersByBadRequest) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCIDGroupMembersByForbidden creates a GetCIDGroupMembersByForbidden with default headers values
func NewGetCIDGroupMembersByForbidden() *GetCIDGroupMembersByForbidden {
	return &GetCIDGroupMembersByForbidden{}
}

/*
GetCIDGroupMembersByForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetCIDGroupMembersByForbidden struct {

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

func (o *GetCIDGroupMembersByForbidden) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByForbidden  %+v", 403, o.Payload)
}
func (o *GetCIDGroupMembersByForbidden) GetPayload() *models.MsaErrorsOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCIDGroupMembersByTooManyRequests creates a GetCIDGroupMembersByTooManyRequests with default headers values
func NewGetCIDGroupMembersByTooManyRequests() *GetCIDGroupMembersByTooManyRequests {
	return &GetCIDGroupMembersByTooManyRequests{}
}

/*
GetCIDGroupMembersByTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetCIDGroupMembersByTooManyRequests struct {

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

func (o *GetCIDGroupMembersByTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIdGroupMembersByTooManyRequests  %+v", 429, o.Payload)
}
func (o *GetCIDGroupMembersByTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetCIDGroupMembersByTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCIDGroupMembersByDefault creates a GetCIDGroupMembersByDefault with default headers values
func NewGetCIDGroupMembersByDefault(code int) *GetCIDGroupMembersByDefault {
	return &GetCIDGroupMembersByDefault{
		_statusCode: code,
	}
}

/*
GetCIDGroupMembersByDefault describes a response with status code -1, with default header values.

OK
*/
type GetCIDGroupMembersByDefault struct {
	_statusCode int

	Payload *models.DomainCIDGroupMembersResponseV1
}

// Code gets the status code for the get c ID group members by default response
func (o *GetCIDGroupMembersByDefault) Code() int {
	return o._statusCode
}

func (o *GetCIDGroupMembersByDefault) Error() string {
	return fmt.Sprintf("[GET /mssp/entities/cid-group-members/v1][%d] getCIDGroupMembersBy default  %+v", o._statusCode, o.Payload)
}
func (o *GetCIDGroupMembersByDefault) GetPayload() *models.DomainCIDGroupMembersResponseV1 {
	return o.Payload
}

func (o *GetCIDGroupMembersByDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DomainCIDGroupMembersResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
