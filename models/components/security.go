package components

type Security struct {
	AccessToken *string `security:"scheme,type=http,subtype=bearer,name=Authorization,env=ruba_access_token"`
}

func (s *Security) GetAccessToken() *string {
	if s == nil {
		return nil
	}
	return s.AccessToken
}
