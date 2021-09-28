// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPcloudVpnconnectionsOptionsGetParams creates a new PcloudVpnconnectionsOptionsGetParams object
// with the default values initialized.
func NewPcloudVpnconnectionsOptionsGetParams() *PcloudVpnconnectionsOptionsGetParams {
	var ()
	return &PcloudVpnconnectionsOptionsGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVpnconnectionsOptionsGetParamsWithTimeout creates a new PcloudVpnconnectionsOptionsGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudVpnconnectionsOptionsGetParamsWithTimeout(timeout time.Duration) *PcloudVpnconnectionsOptionsGetParams {
	var ()
	return &PcloudVpnconnectionsOptionsGetParams{

		timeout: timeout,
	}
}

// NewPcloudVpnconnectionsOptionsGetParamsWithContext creates a new PcloudVpnconnectionsOptionsGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudVpnconnectionsOptionsGetParamsWithContext(ctx context.Context) *PcloudVpnconnectionsOptionsGetParams {
	var ()
	return &PcloudVpnconnectionsOptionsGetParams{

		Context: ctx,
	}
}

// NewPcloudVpnconnectionsOptionsGetParamsWithHTTPClient creates a new PcloudVpnconnectionsOptionsGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudVpnconnectionsOptionsGetParamsWithHTTPClient(client *http.Client) *PcloudVpnconnectionsOptionsGetParams {
	var ()
	return &PcloudVpnconnectionsOptionsGetParams{
		HTTPClient: client,
	}
}

/*PcloudVpnconnectionsOptionsGetParams contains all the parameters to send to the API endpoint
for the pcloud vpnconnections options get operation typically these are written to a http.Request
*/
type PcloudVpnconnectionsOptionsGetParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud vpnconnections options get params
func (o *PcloudVpnconnectionsOptionsGetParams) WithTimeout(timeout time.Duration) *PcloudVpnconnectionsOptionsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud vpnconnections options get params
func (o *PcloudVpnconnectionsOptionsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud vpnconnections options get params
func (o *PcloudVpnconnectionsOptionsGetParams) WithContext(ctx context.Context) *PcloudVpnconnectionsOptionsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud vpnconnections options get params
func (o *PcloudVpnconnectionsOptionsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud vpnconnections options get params
func (o *PcloudVpnconnectionsOptionsGetParams) WithHTTPClient(client *http.Client) *PcloudVpnconnectionsOptionsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud vpnconnections options get params
func (o *PcloudVpnconnectionsOptionsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud vpnconnections options get params
func (o *PcloudVpnconnectionsOptionsGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudVpnconnectionsOptionsGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud vpnconnections options get params
func (o *PcloudVpnconnectionsOptionsGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVpnconnectionsOptionsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}