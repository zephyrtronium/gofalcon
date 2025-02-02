// Code generated by go-swagger; DO NOT EDIT.

package quick_scan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new quick scan API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for quick scan API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetScans(params *GetScansParams, opts ...ClientOption) (*GetScansOK, error)

	GetScansAggregates(params *GetScansAggregatesParams, opts ...ClientOption) (*GetScansAggregatesOK, error)

	QuerySubmissionsMixin0(params *QuerySubmissionsMixin0Params, opts ...ClientOption) (*QuerySubmissionsMixin0OK, error)

	ScanSamples(params *ScanSamplesParams, opts ...ClientOption) (*ScanSamplesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
GetScans checks the status of a volume scan time required for analysis increases with the number of samples in a volume but usually it should take less than 1 minute
*/
func (a *Client) GetScans(params *GetScansParams, opts ...ClientOption) (*GetScansOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScansParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScans",
		Method:             "GET",
		PathPattern:        "/scanner/entities/scans/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScansReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetScansOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetScansDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetScansAggregates gets scans aggregations as specified via json in request body
*/
func (a *Client) GetScansAggregates(params *GetScansAggregatesParams, opts ...ClientOption) (*GetScansAggregatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetScansAggregatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetScansAggregates",
		Method:             "POST",
		PathPattern:        "/scanner/aggregates/scans/GET/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScansAggregatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetScansAggregatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetScansAggregates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
QuerySubmissionsMixin0 finds i ds for submitted scans by providing an f q l filter and paging details returns a set of volume i ds that match your criteria
*/
func (a *Client) QuerySubmissionsMixin0(params *QuerySubmissionsMixin0Params, opts ...ClientOption) (*QuerySubmissionsMixin0OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQuerySubmissionsMixin0Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "QuerySubmissionsMixin0",
		Method:             "GET",
		PathPattern:        "/scanner/queries/scans/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &QuerySubmissionsMixin0Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QuerySubmissionsMixin0OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*QuerySubmissionsMixin0Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ScanSamples submits a volume of files for ml scanning time required for analysis increases with the number of samples in a volume but usually it should take less than 1 minute
*/
func (a *Client) ScanSamples(params *ScanSamplesParams, opts ...ClientOption) (*ScanSamplesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScanSamplesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ScanSamples",
		Method:             "POST",
		PathPattern:        "/scanner/entities/scans/v1",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScanSamplesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ScanSamplesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ScanSamplesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
