package pay

import (
	"encoding/xml"
)

type CPayByPaymentCodeRequest struct {
	XMLName        xml.Name `xml:"xml"`
	AppId          string   `xml:"appid"`
	MchId          string   `xml:"mch_id"`
	DeviceInfo     string   `xml:"device_info"`
	NonceStr       string   `xml:"nonce_str"`
	Sign           string   `xml:"sign"`
	SignType       string   `xml:"sign_type"`
	Body           string   `xml:"body"`
	Detail         string   `xml:"detail"`
	Attach         string   `xml:"attach"`
	OutTradeNo     string   `xml:"out_trade_no"`
	TotalFee       int      `xml:"total_fee"`
	FeeType        string   `xml:"fee_type"`
	SpbillCreateIp string   `xml:"spbill_create_ip"`
	GoodsTag       string   `xml:"goods_tag"`
	LimitPay       string   `xml:"limit_pay"`
	TimeStart      string   `xml:"time_start"`
	TimeExpire     string   `xml:"time_expire"`
	Receipt        string   `xml:"receipt"`
	AuthCode       string   `xml:"auth_code"`
	SceneInfo      string   `xml:"scene_info"`
}
