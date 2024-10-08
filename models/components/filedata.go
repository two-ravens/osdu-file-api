// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Endian - Endianness of binary value. Enumeration- \BIG\ \LITTLE\.  If absent applications will need to interpret from context indicators.
type Endian string

const (
	EndianBig    Endian = "BIG"
	EndianLittle Endian = "LITTLE"
)

func (e Endian) ToPointer() *Endian {
	return &e
}
func (e *Endian) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BIG":
		fallthrough
	case "LITTLE":
		*e = Endian(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Endian: %v", v)
	}
}

// ExtensionProperties - File DMS Extension Properties
type ExtensionProperties struct {
}

// FileData - The file data container containing all necessary details of the file record
type FileData struct {
	// An optional name of the dataset, e.g. a user friendly file or file collection name.
	Name *string `json:"Name,omitempty"`
	// An optional, textual description of the dataset.
	Description *string `json:"Description,omitempty"`
	// Total size of the dataset in bytes; for files it is the same as declared in FileSourceInfo.FileSize or the sum of all individual files. Implemented as string. The value must be convertible to a long integer (sizes can become very large).
	TotalSize *string `json:"TotalSize,omitempty"`
	// Encoding Format Type ID
	EncodingFormatTypeID *string `json:"EncodingFormatTypeID,omitempty"`
	// Schema Format Type ID
	SchemaFormatTypeID *string `json:"SchemaFormatTypeID,omitempty"`
	// Resource Home Region ID
	ResourceHomeRegionID *string `json:"ResourceHomeRegionID,omitempty"`
	// Resource Host Region IDs
	ResourceHostRegionIDs []string `json:"ResourceHostRegionIDs,omitempty"`
	// Resource Curation Status
	ResourceCurationStatus *string `json:"ResourceCurationStatus,omitempty"`
	// Resource Lifecycle Status
	ResourceLifecycleStatus *string `json:"ResourceLifecycleStatus,omitempty"`
	// Resource Security Classification
	ResourceSecurityClassification *string `json:"ResourceSecurityClassification,omitempty"`
	// Source
	Source *string `json:"Source,omitempty"`
	// Dataset Properties
	DatasetProperties DatasetProperties `json:"DatasetProperties"`
	// Existence Kind
	ExistenceKind *string `json:"ExistenceKind,omitempty"`
	// Endianness of binary value. Enumeration- \BIG\ \LITTLE\.  If absent applications will need to interpret from context indicators.
	Endian *Endian `json:"Endian,omitempty"`
	// MD5 checksum of file bytes - a 32 byte hexadecimal number
	Checksum *string `json:"Checksum,omitempty"`
	// File DMS Extension Properties
	ExtensionProperties map[string]ExtensionProperties `json:"ExtensionProperties,omitempty"`
}

func (o *FileData) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *FileData) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *FileData) GetTotalSize() *string {
	if o == nil {
		return nil
	}
	return o.TotalSize
}

func (o *FileData) GetEncodingFormatTypeID() *string {
	if o == nil {
		return nil
	}
	return o.EncodingFormatTypeID
}

func (o *FileData) GetSchemaFormatTypeID() *string {
	if o == nil {
		return nil
	}
	return o.SchemaFormatTypeID
}

func (o *FileData) GetResourceHomeRegionID() *string {
	if o == nil {
		return nil
	}
	return o.ResourceHomeRegionID
}

func (o *FileData) GetResourceHostRegionIDs() []string {
	if o == nil {
		return nil
	}
	return o.ResourceHostRegionIDs
}

func (o *FileData) GetResourceCurationStatus() *string {
	if o == nil {
		return nil
	}
	return o.ResourceCurationStatus
}

func (o *FileData) GetResourceLifecycleStatus() *string {
	if o == nil {
		return nil
	}
	return o.ResourceLifecycleStatus
}

func (o *FileData) GetResourceSecurityClassification() *string {
	if o == nil {
		return nil
	}
	return o.ResourceSecurityClassification
}

func (o *FileData) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *FileData) GetDatasetProperties() DatasetProperties {
	if o == nil {
		return DatasetProperties{}
	}
	return o.DatasetProperties
}

func (o *FileData) GetExistenceKind() *string {
	if o == nil {
		return nil
	}
	return o.ExistenceKind
}

func (o *FileData) GetEndian() *Endian {
	if o == nil {
		return nil
	}
	return o.Endian
}

func (o *FileData) GetChecksum() *string {
	if o == nil {
		return nil
	}
	return o.Checksum
}

func (o *FileData) GetExtensionProperties() map[string]ExtensionProperties {
	if o == nil {
		return nil
	}
	return o.ExtensionProperties
}
