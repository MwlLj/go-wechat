package user

type CUserOpenIdInfo struct {
	OpenId []string `json:"openid"`
}

// is NextOpenId is "" -> over
type CSingleUsers struct {
	ErrCode    int             `json:"errcode"`
	ErrMsg     string          `json:"errmsg"`
	Total      int             `json:"total"`
	Count      int             `json:"count"`
	Data       CUserOpenIdInfo `json:"data"`
	NextOpenId string          `json:"next_openid"`
}

type CGetBlackListSingle struct {
	BeginOpenId string `json:"begin_openid"`
}
