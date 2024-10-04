// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type GetFileMetadataByIDRequest struct {
	// File metadata record Id.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Tenant Id
	DataPartitionID string `header:"style=simple,explode=false,name=data-partition-id"`
}

func (o *GetFileMetadataByIDRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetFileMetadataByIDRequest) GetDataPartitionID() string {
	if o == nil {
		return ""
	}
	return o.DataPartitionID
}

type GetFileMetadataByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Body     []byte
}

func (o *GetFileMetadataByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetFileMetadataByIDResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
