package operations

import (
	"github.com/Rubadot/ruba-go/models/components"
)

type MembersGetMemberRequest struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (m *MembersGetMemberRequest) GetID() string {
	if m == nil {
		return ""
	}
	return m.ID
}

type MembersGetMemberResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Member retrieved.
	Member *components.Member
}

func (m *MembersGetMemberResponse) GetHTTPMeta() components.HTTPMetadata {
	if m == nil {
		return components.HTTPMetadata{}
	}
	return m.HTTPMeta
}

func (m *MembersGetMemberResponse) GetMember() *components.Member {
	if m == nil {
		return nil
	}
	return m.Member
}
