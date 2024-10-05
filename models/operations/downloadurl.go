// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/two-ravens/osdu-file-api/models/components"
)

type DownloadURLRequest struct {
	// File Metadata record Id.
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// The Time for which Signed URL to be valid. Accepted Regex patterns are "^[0-9]+M$", "^[0-9]+H$", "^[0-9]+D$" denoting Integer values in Minutes, Hours, Days respectively. In absence of this parameter the URL would be valid for 1 Hour.
	ExpiryTime *string `queryParam:"style=form,explode=true,name=expiryTime"`
	// Tenant Id
	DataPartitionID string `header:"style=simple,explode=false,name=data-partition-id"`
}

func (o *DownloadURLRequest) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *DownloadURLRequest) GetExpiryTime() *string {
	if o == nil {
		return nil
	}
	return o.ExpiryTime
}

func (o *DownloadURLRequest) GetDataPartitionID() string {
	if o == nil {
		return ""
	}
	return o.DataPartitionID
}

type DownloadURLResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Body     []byte
}

func (o *DownloadURLResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DownloadURLResponse) GetBody() []byte {
	if o == nil {
		return nil
	}
	return o.Body
}
