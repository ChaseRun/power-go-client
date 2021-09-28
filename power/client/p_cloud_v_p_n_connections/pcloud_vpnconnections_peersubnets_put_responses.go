// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_v_p_n_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudVpnconnectionsPeersubnetsPutReader is a Reader for the PcloudVpnconnectionsPeersubnetsPut structure.
type PcloudVpnconnectionsPeersubnetsPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudVpnconnectionsPeersubnetsPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPcloudVpnconnectionsPeersubnetsPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPcloudVpnconnectionsPeersubnetsPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPcloudVpnconnectionsPeersubnetsPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPcloudVpnconnectionsPeersubnetsPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPcloudVpnconnectionsPeersubnetsPutUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPcloudVpnconnectionsPeersubnetsPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPcloudVpnconnectionsPeersubnetsPutOK creates a PcloudVpnconnectionsPeersubnetsPutOK with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutOK() *PcloudVpnconnectionsPeersubnetsPutOK {
	return &PcloudVpnconnectionsPeersubnetsPutOK{}
}

/*PcloudVpnconnectionsPeersubnetsPutOK handles this case with default header values.

OK
*/
type PcloudVpnconnectionsPeersubnetsPutOK struct {
	Payload *models.PeerSubnets
}

func (o *PcloudVpnconnectionsPeersubnetsPutOK) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutOK  %+v", 200, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PeerSubnets)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutBadRequest creates a PcloudVpnconnectionsPeersubnetsPutBadRequest with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutBadRequest() *PcloudVpnconnectionsPeersubnetsPutBadRequest {
	return &PcloudVpnconnectionsPeersubnetsPutBadRequest{}
}

/*PcloudVpnconnectionsPeersubnetsPutBadRequest handles this case with default header values.

Bad Request
*/
type PcloudVpnconnectionsPeersubnetsPutBadRequest struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutUnauthorized creates a PcloudVpnconnectionsPeersubnetsPutUnauthorized with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutUnauthorized() *PcloudVpnconnectionsPeersubnetsPutUnauthorized {
	return &PcloudVpnconnectionsPeersubnetsPutUnauthorized{}
}

/*PcloudVpnconnectionsPeersubnetsPutUnauthorized handles this case with default header values.

Unauthorized
*/
type PcloudVpnconnectionsPeersubnetsPutUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutForbidden creates a PcloudVpnconnectionsPeersubnetsPutForbidden with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutForbidden() *PcloudVpnconnectionsPeersubnetsPutForbidden {
	return &PcloudVpnconnectionsPeersubnetsPutForbidden{}
}

/*PcloudVpnconnectionsPeersubnetsPutForbidden handles this case with default header values.

Forbidden
*/
type PcloudVpnconnectionsPeersubnetsPutForbidden struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutForbidden  %+v", 403, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutUnprocessableEntity creates a PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutUnprocessableEntity() *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity {
	return &PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity{}
}

/*PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity handles this case with default header values.

Unprocessable Entity
*/
type PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudVpnconnectionsPeersubnetsPutInternalServerError creates a PcloudVpnconnectionsPeersubnetsPutInternalServerError with default headers values
func NewPcloudVpnconnectionsPeersubnetsPutInternalServerError() *PcloudVpnconnectionsPeersubnetsPutInternalServerError {
	return &PcloudVpnconnectionsPeersubnetsPutInternalServerError{}
}

/*PcloudVpnconnectionsPeersubnetsPutInternalServerError handles this case with default header values.

Internal Server Error
*/
type PcloudVpnconnectionsPeersubnetsPutInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/vpn-connections/{vpn_connection_id}/peer-subnets][%d] pcloudVpnconnectionsPeersubnetsPutInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudVpnconnectionsPeersubnetsPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}