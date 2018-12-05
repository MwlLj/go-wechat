package user

var (
	CreateTagUrl        string = "https://api.weixin.qq.com/cgi-bin/tags/create"
	GetTagListUrl       string = "https://api.weixin.qq.com/cgi-bin/tags/get"
	UpdateTagUrl        string = "https://api.weixin.qq.com/cgi-bin/tags/update"
	DeleteTagUrl        string = "https://api.weixin.qq.com/cgi-bin/tags/delete"
	GetTagUserListUrl   string = "https://api.weixin.qq.com/cgi-bin/user/tag/get"
	AddTagToUsersUrl    string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging"
	DeleteTagToUsersUrl string = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging"
	GetTagsByUserUrl    string = "https://api.weixin.qq.com/cgi-bin/tags/getidlist"
)
