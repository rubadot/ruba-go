package components

type SeatClaim struct {
	// Invitation token to claim the seat
	InvitationToken string `json:"invitation_token"`
}

func (s *SeatClaim) GetInvitationToken() string {
	if s == nil {
		return ""
	}
	return s.InvitationToken
}
