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
	OutTradenn     string   `xml:"out_trade_no"`
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

type CPayByPaymentCodeResponse struct {
	XMLName            xml.Name `xml:"xml"`
	ReturnCode         string   `xml:"return_code"`
	ReturnMsg          string   `xml:"return_msg"`
	Appid              string   `xml:"appid"`
	MchId              string   `xml:"mch_id"`
	DeviceInfo         string   `xml:"device_info"`
	NonceStr           string   `xml:"nonce_str"`
	Sign               string   `xml:"sign"`
	ResultCode         string   `xml:"result_code"`
	ErrCode            string   `xml:"err_code"`
	ErrCodeDes         string   `xml:"err_code_des"`
	OpenId             string   `xml:"openid"`
	IsSubscribe        string   `xml:"is_subscribe"`
	TradeType          string   `xml:"trade_type"`
	BankType           string   `xml:"bank_type"`
	FeeType            string   `xml:"fee_type"`
	TotalFee           int      `xml:"total_fee"`
	SettlementTotalFee int      `xml:"settlement_total_fee"`
	CouponFee          int      `xml:"coupon_fee"`
	CashFeeType        string   `xml:"cash_fee_type"`
	CashFee            int      `xml:"cash_fee"`
	TransactionId      string   `xml:"transaction_id"`
	OutTradeNo         string   `xml:"out_trade_no"`
	Attach             string   `xml:"attach"`
	TimeEnd            string   `xml:"time_end"`
	PromotionDetail    string   `xml:"promotion_detail"`
}
