// Code generated by go-swagger; DO NOT EDIT.

package overwatch_dashboard

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

// AggregatesEventsCollectionsReader is a Reader for the AggregatesEventsCollections structure.
type AggregatesEventsCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatesEventsCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatesEventsCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregatesEventsCollectionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregatesEventsCollectionsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewAggregatesEventsCollectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregatesEventsCollectionsOK creates a AggregatesEventsCollectionsOK with default headers values
func NewAggregatesEventsCollectionsOK() *AggregatesEventsCollectionsOK {
	return &AggregatesEventsCollectionsOK{}
}

/*
AggregatesEventsCollectionsOK describes a response with status code 200, with default header values.

OK
*/
type AggregatesEventsCollectionsOK struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

func (o *AggregatesEventsCollectionsOK) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsOK  %+v", 200, o.Payload)
}
func (o *AggregatesEventsCollectionsOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregatesEventsCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatesEventsCollectionsForbidden creates a AggregatesEventsCollectionsForbidden with default headers values
func NewAggregatesEventsCollectionsForbidden() *AggregatesEventsCollectionsForbidden {
	return &AggregatesEventsCollectionsForbidden{}
}

/*
AggregatesEventsCollectionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregatesEventsCollectionsForbidden struct {

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

func (o *AggregatesEventsCollectionsForbidden) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsForbidden  %+v", 403, o.Payload)
}
func (o *AggregatesEventsCollectionsForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesEventsCollectionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatesEventsCollectionsTooManyRequests creates a AggregatesEventsCollectionsTooManyRequests with default headers values
func NewAggregatesEventsCollectionsTooManyRequests() *AggregatesEventsCollectionsTooManyRequests {
	return &AggregatesEventsCollectionsTooManyRequests{}
}

/*
AggregatesEventsCollectionsTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregatesEventsCollectionsTooManyRequests struct {

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

func (o *AggregatesEventsCollectionsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] aggregatesEventsCollectionsTooManyRequests  %+v", 429, o.Payload)
}
func (o *AggregatesEventsCollectionsTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregatesEventsCollectionsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregatesEventsCollectionsDefault creates a AggregatesEventsCollectionsDefault with default headers values
func NewAggregatesEventsCollectionsDefault(code int) *AggregatesEventsCollectionsDefault {
	return &AggregatesEventsCollectionsDefault{
		_statusCode: code,
	}
}

/*
AggregatesEventsCollectionsDefault describes a response with status code -1, with default header values.

OK
*/
type AggregatesEventsCollectionsDefault struct {
	_statusCode int

	Payload *models.MsaAggregatesResponse
}

// Code gets the status code for the aggregates events collections default response
func (o *AggregatesEventsCollectionsDefault) Code() int {
	return o._statusCode
}

func (o *AggregatesEventsCollectionsDefault) Error() string {
	return fmt.Sprintf("[POST /overwatch-dashboards/aggregates/events-collections/GET/v1][%d] AggregatesEventsCollections default  %+v", o._statusCode, o.Payload)
}
func (o *AggregatesEventsCollectionsDefault) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregatesEventsCollectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
