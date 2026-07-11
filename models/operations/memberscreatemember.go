package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type MembersCreateMemberResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Member created.
	Member *components.Member
}

func (m *MembersCreateMemberResponse) GetHTTPMeta() components.HTTPMetadata {
	if m == nil {
		return components.HTTPMetadata{}
	}
	return m.HTTPMeta
}

func (m *MembersCreateMemberResponse) GetMember() *components.Member {
	if m == nil {
		return nil
	}
	return m.Member
}
