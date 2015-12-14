package setting

type OAuthInfo struct {
	ClientId, ClientSecret string
	Scopes                 []string
	AuthUrl, TokenUrl      string
	Enabled                bool
	AllowedDomains         []string
	ApiUrl                 string
	AllowSignup            bool
}

type OAuther struct {
	GitHub, Google, Twitter, Uoko bool
	OAuthInfos              map[string]*OAuthInfo
}

var OAuthService *OAuther
