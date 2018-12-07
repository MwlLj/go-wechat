package user

type CFollowUserInfo struct {
	OpenId []string `json:"openid"`
}

// is NextOpenId is "" -> over
type CSingleFollowUsers struct {
	ErrCode    int             `json:"errcode"`
	ErrMsg     string          `json:"errmsg"`
	Total      int             `json:"total"`
	Count      int             `json:"count"`
	Data       CFollowUserInfo `json:"data"`
	NextOpenId string          `json:"next_openid"`
}
