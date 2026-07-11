package components

import (
	"github.com/Rubadot/ruba-go/internal/utils"
)

// BenefitDiscordProperties - Properties for a benefit of type `discord`.
type BenefitDiscordProperties struct {
	// The ID of the Discord server.
	GuildID string `json:"guild_id"`
	// The ID of the Discord role to grant.
	RoleID string `json:"role_id"`
	// Whether to kick the member from the Discord server on revocation.
	KickMember bool   `json:"kick_member"`
	GuildToken string `json:"guild_token"`
}

func (b BenefitDiscordProperties) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(b, "", false)
}

func (b *BenefitDiscordProperties) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &b, "", false, []string{"guild_id", "role_id", "kick_member", "guild_token"}); err != nil {
		return err
	}
	return nil
}

func (b *BenefitDiscordProperties) GetGuildID() string {
	if b == nil {
		return ""
	}
	return b.GuildID
}

func (b *BenefitDiscordProperties) GetRoleID() string {
	if b == nil {
		return ""
	}
	return b.RoleID
}

func (b *BenefitDiscordProperties) GetKickMember() bool {
	if b == nil {
		return false
	}
	return b.KickMember
}

func (b *BenefitDiscordProperties) GetGuildToken() string {
	if b == nil {
		return ""
	}
	return b.GuildToken
}
