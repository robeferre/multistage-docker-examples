// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/jhh3/multistage-docker-examples/go-example/models"
)

// UploadFileOKCode is the HTTP code returned for type UploadFileOK
const UploadFileOKCode int = 200

/*UploadFileOK successful operation

swagger:response uploadFileOK
*/
type UploadFileOK struct {

	/*
	  In: Body
	*/
	Payload *models.APIResponse `json:"body,omitempty"`
}

// NewUploadFileOK creates UploadFileOK with default headers values
func NewUploadFileOK() *UploadFileOK {

	return &UploadFileOK{}
}

// WithPayload adds the payload to the upload file o k response
func (o *UploadFileOK) WithPayload(payload *models.APIResponse) *UploadFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the upload file o k response
func (o *UploadFileOK) SetPayload(payload *models.APIResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UploadFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
