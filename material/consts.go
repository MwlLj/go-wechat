package material

var (
	CreateTmpMaterialParamType          string = "type"
	CreateTmpMaterialFormname           string = "media"
	GetTmpMaterialParamMediaId          string = "media_id"
	GetTmpHDMaterialParamMediaId        string = "media_id"
	UploadImageFormname                 string = "media"
	AddOtherMaterialParamType           string = "type"
	AddOtherMaterialFormnameMedia       string = "media"
	AddOtherMaterialFormnameDescription string = "description"
)

var (
	CreateTmpMaterialUrl            string = "https://api.weixin.qq.com/cgi-bin/media/upload"
	GetTmpMaterialUrl               string = "https://api.weixin.qq.com/cgi-bin/media/get"
	GetTmpHDMaterialUrl             string = "https://api.weixin.qq.com/cgi-bin/media/get"
	AddForeverImgTextMaterialUrl    string = "https://api.weixin.qq.com/cgi-bin/material/add_news"
	UploadImageUrl                  string = "https://api.weixin.qq.com/cgi-bin/media/uploadimg"
	AddForeverOtherMaterialUrl      string = "https://api.weixin.qq.com/cgi-bin/material/add_material"
	GetForeverMaterialUrl           string = "https://api.weixin.qq.com/cgi-bin/material/get_material"
	DeleteForeverMaterialUrl        string = "https://api.weixin.qq.com/cgi-bin/material/del_material"
	UpdateForeverImgTextMaterialUrl string = "https://api.weixin.qq.com/cgi-bin/material/update_news"
	GetMaterialTotal                string = "https://api.weixin.qq.com/cgi-bin/material/get_materialcount"
	GetMaterialList                 string = "https://api.weixin.qq.com/cgi-bin/material/batchget_material"
)
