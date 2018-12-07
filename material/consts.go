package material

var (
	CreateTmpMaterialParamType          string   = "type"
	CreateTmpMaterialFormname           string   = "media"
	GetTmpMaterialResUrls               []string = []string{"video_url", "image_url", "voice_url", "thumb_url", "news_url"}
	GetTmpMaterialParamMediaId          string   = "media_id"
	GetTmpHDMaterialParamMediaId        string   = "media_id"
	UploadImageFormname                 string   = "media"
	AddOtherMaterialParamType           string   = "type"
	AddOtherMaterialFormnameMedia       string   = "media"
	AddOtherMaterialFormnameDescription string   = "description"
)

var (
	CreateTmpMaterialUrl            string = "https://api.weixin.qq.com/cgi-bin/media/upload"
	GetTmpMaterialUrl               string = "https://api.weixin.qq.com/cgi-bin/media/get"
	GetTmpHDMaterialUrl             string = "https://api.weixin.qq.com/cgi-bin/media/get"
	AddForeverImgTextMaterialUrl    string = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	UploadImageUrl                  string = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	UploadVideoUrl                  string = "https://api.weixin.qq.com/cgi-bin/media/uploadvideo"
	AddForeverOtherMaterialUrl      string = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	GetForeverMaterialUrl           string = "https://api.weixin.qq.com/cgi-bin/material/get_material"
	DeleteForeverMaterialUrl        string = "https://api.weixin.qq.com/cgi-bin/material/del_material"
	UpdateForeverImgTextMaterialUrl string = "https://api.weixin.qq.com/cgi-bin/material/update_news"
	GetMaterialTotal                string = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount"
	GetMaterialList                 string = "https://api.weixin.qq.com/cgi-bin/material/batchget_material"
)
