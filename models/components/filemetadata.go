// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type Meta struct {
}

// FileMetadata - File metadata record.
type FileMetadata struct {
	// Unique identifier generated by the system for the file metadata record.
	ID *string `json:"id,omitempty"`
	// Kind of data being ingested. Must follow the naming
	//   convention:data-Partition-Id}:dataset-name}:record-type}:version}.
	Kind  string `json:"kind"`
	Legal Legal  `json:"legal"`
	// The file data container containing all necessary details of the file record
	Data FileData `json:"data"`
	// A named list of entities in the data lake as a dictionary item.
	Ancestry *Ancestry         `json:"ancestry,omitempty"`
	Meta     []map[string]Meta `json:"meta,omitempty"`
	Tags     map[string]string `json:"tags,omitempty"`
	ACL      ACL               `json:"acl"`
}

func (o *FileMetadata) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *FileMetadata) GetKind() string {
	if o == nil {
		return ""
	}
	return o.Kind
}

func (o *FileMetadata) GetLegal() Legal {
	if o == nil {
		return Legal{}
	}
	return o.Legal
}

func (o *FileMetadata) GetData() FileData {
	if o == nil {
		return FileData{}
	}
	return o.Data
}

func (o *FileMetadata) GetAncestry() *Ancestry {
	if o == nil {
		return nil
	}
	return o.Ancestry
}

func (o *FileMetadata) GetMeta() []map[string]Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *FileMetadata) GetTags() map[string]string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *FileMetadata) GetACL() ACL {
	if o == nil {
		return ACL{}
	}
	return o.ACL
}
