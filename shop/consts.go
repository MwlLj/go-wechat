package shop

var (
	AddCommodityUrl             string = "https://api.weixin.qq.com/merchant/create"
	DeleteCommodityUrl          string = "https://api.weixin.qq.com/merchant/del"
	GetCommodityUrl             string = "https://api.weixin.qq.com/merchant/get"
	GetCommodityByStatusUrl     string = "https://api.weixin.qq.com/merchant/getbystatus"
	UpdateCommodityStatusUrl    string = "https://api.weixin.qq.com/merchant/modproductstatus"
	GetSubClassesByClassifyUrl  string = "https://api.weixin.qq.com/merchant/category/getsub"
	GetAllSkuByClassifyUrl      string = "https://api.weixin.qq.com/merchant/category/getsku"
	GetAllPropertyByClassifyUrl string = "https://api.weixin.qq.com/merchant/category/getproperty"

	AddStockUrl    string = "https://api.weixin.qq.com/merchant/stock/add"
	ReduceStockUrl string = "https://api.weixin.qq.com/merchant/stock/reduce"
)
