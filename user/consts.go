package user

var (
	GetUserBaseInfoParamOpenId    string = "openid"
	GetUserBaseInfoParamLang      string = "lang"
	GetFollowUsersParamNextOpenId string = "next_openid"
)

var (
	CreateTagUrl        string = "https://api.weixin.qq.com/cgi-bin/tags/create"
	GetTagListUrl       string = "https://api.weixin.qq.com/cgi-bin/tags/get"
	UpdateTagUrl        string = "https://api.weixin.qq.com/cgi-bin/tags/update"
	DeleteTagUrl        string = "https://api.weixin.qq.com/cgi-bin/tags/delete"
	GetTagUserListUrl   string = "https://api.weixin.qq.com/cgi-bin/user/tag/get"
	AddTagToUsersUrl    string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging"
	DeleteTagToUsersUrl string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging"
	GetTagsByUserUrl    string = "https://api.weixin.qq.com/cgi-bin/tags/getidlist"

	UpdateUserRemarkUrl string = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark"

	GetUserBaseInfoUrl      string = "https://api.weixin.qq.com/cgi-bin/user/info"
	GetUserBaseInfoMultiUrl string = "https://api.weixin.qq.com/cgi-bin/user/info/batchget"

	GetFollowUsersUrl string = "https://api.weixin.qq.com/cgi-bin/user/get"

	GetBlackListUsersUrl      string = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist"
	TakeUsersToBlackListUrl   string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist"
	UnTakeUsersToBlackListUrl string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist"
)
