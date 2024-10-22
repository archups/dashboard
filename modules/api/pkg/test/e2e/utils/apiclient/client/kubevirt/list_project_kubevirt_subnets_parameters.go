// Code generated by go-swagger; DO NOT EDIT.

package kubevirt

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListProjectKubevirtSubnetsParams creates a new ListProjectKubevirtSubnetsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListProjectKubevirtSubnetsParams() *ListProjectKubevirtSubnetsParams {
	return &ListProjectKubevirtSubnetsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectKubevirtSubnetsParamsWithTimeout creates a new ListProjectKubevirtSubnetsParams object
// with the ability to set a timeout on a request.
func NewListProjectKubevirtSubnetsParamsWithTimeout(timeout time.Duration) *ListProjectKubevirtSubnetsParams {
	return &ListProjectKubevirtSubnetsParams{
		timeout: timeout,
	}
}

// NewListProjectKubevirtSubnetsParamsWithContext creates a new ListProjectKubevirtSubnetsParams object
// with the ability to set a context for a request.
func NewListProjectKubevirtSubnetsParamsWithContext(ctx context.Context) *ListProjectKubevirtSubnetsParams {
	return &ListProjectKubevirtSubnetsParams{
		Context: ctx,
	}
}

// NewListProjectKubevirtSubnetsParamsWithHTTPClient creates a new ListProjectKubevirtSubnetsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListProjectKubevirtSubnetsParamsWithHTTPClient(client *http.Client) *ListProjectKubevirtSubnetsParams {
	return &ListProjectKubevirtSubnetsParams{
		HTTPClient: client,
	}
}

/*
ListProjectKubevirtSubnetsParams contains all the parameters to send to the API endpoint

	for the list project kubevirt subnets operation.

	Typically these are written to a http.Request.
*/
type ListProjectKubevirtSubnetsParams struct {

	// Credential.
	Credential *string

	// Kubeconfig.
	Kubeconfig *string

	// VPCName.
	VPCName *string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list project kubevirt subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectKubevirtSubnetsParams) WithDefaults() *ListProjectKubevirtSubnetsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list project kubevirt subnets params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListProjectKubevirtSubnetsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) WithTimeout(timeout time.Duration) *ListProjectKubevirtSubnetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) WithContext(ctx context.Context) *ListProjectKubevirtSubnetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) WithHTTPClient(client *http.Client) *ListProjectKubevirtSubnetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCredential adds the credential to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) WithCredential(credential *string) *ListProjectKubevirtSubnetsParams {
	o.SetCredential(credential)
	return o
}

// SetCredential adds the credential to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) SetCredential(credential *string) {
	o.Credential = credential
}

// WithKubeconfig adds the kubeconfig to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) WithKubeconfig(kubeconfig *string) *ListProjectKubevirtSubnetsParams {
	o.SetKubeconfig(kubeconfig)
	return o
}

// SetKubeconfig adds the kubeconfig to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) SetKubeconfig(kubeconfig *string) {
	o.Kubeconfig = kubeconfig
}

// WithVPCName adds the vPCName to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) WithVPCName(vPCName *string) *ListProjectKubevirtSubnetsParams {
	o.SetVPCName(vPCName)
	return o
}

// SetVPCName adds the vPCName to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) SetVPCName(vPCName *string) {
	o.VPCName = vPCName
}

// WithProjectID adds the projectID to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) WithProjectID(projectID string) *ListProjectKubevirtSubnetsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list project kubevirt subnets params
func (o *ListProjectKubevirtSubnetsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectKubevirtSubnetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Credential != nil {

		// header param Credential
		if err := r.SetHeaderParam("Credential", *o.Credential); err != nil {
			return err
		}
	}

	if o.Kubeconfig != nil {

		// header param Kubeconfig
		if err := r.SetHeaderParam("Kubeconfig", *o.Kubeconfig); err != nil {
			return err
		}
	}

	if o.VPCName != nil {

		// header param VPCName
		if err := r.SetHeaderParam("VPCName", *o.VPCName); err != nil {
			return err
		}
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
