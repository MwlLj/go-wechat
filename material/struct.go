package material

type CGetTmpMaterialResponse struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	ResUrl  interface{} `json:"-"`
}

type CGetTmpHDMaterialResponse struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  string      `json:"errmsg"`
	ResUrl  interface{} `json:"-"`
}

type CAddForeverOtherMaterialFormnameDesc struct {
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}
