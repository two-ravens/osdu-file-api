// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/two-ravens/osdu-file-api/models/components"
)

type RevokeURLResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
}

func (o *RevokeURLResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}
