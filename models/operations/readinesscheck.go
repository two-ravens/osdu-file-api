// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"openapi/models/components"
)

type ReadinessCheckRequest struct {
	// Tenant Id
	DataPartitionID string `header:"style=simple,explode=false,name=data-partition-id"`
}

func (o *ReadinessCheckRequest) GetDataPartitionID() string {
	if o == nil {
		return ""
	}
	return o.DataPartitionID
}

type ReadinessCheckResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	Res *string
}

func (o *ReadinessCheckResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *ReadinessCheckResponse) GetRes() *string {
	if o == nil {
		return nil
	}
	return o.Res
}