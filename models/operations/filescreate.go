package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type FilesCreateResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// File created.
	FileUpload *components.FileUpload
}

func (f *FilesCreateResponse) GetHTTPMeta() components.HTTPMetadata {
	if f == nil {
		return components.HTTPMetadata{}
	}
	return f.HTTPMeta
}

func (f *FilesCreateResponse) GetFileUpload() *components.FileUpload {
	if f == nil {
		return nil
	}
	return f.FileUpload
}
