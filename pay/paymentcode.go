package pay

import (
	"encoding/xml"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"github.com/MwlLj/go-wechat/private"
	"net/http"
	"strconv"
)

type CPayByPaymentCode struct {
	m_token common.IToken
}

func (this *CPayByPaymentCode) Commit(request *common.CPayByPaymentCodeRequest, timeoutMS int64) (*common.CPayByPaymentCodeResponse, error) {
	req := CPayByPaymentCodeRequest{}
	info := this.m_token.GetUserInfo()
	req.AppId = info.AppId
	req.Attach = request.Attach
	req.AuthCode = request.AuthCode
	req.Body = request.Body
	req.Detail = request.Detail
	req.DeviceInfo = request.DeviceInfo
	req.FeeType = request.FeeType
	req.GoodsTag = request.GoodsTag
	req.LimitPay = request.LimitPay
	req.MchId = request.MchId
	req.NonceStr = string(genRandomString(32))
	req.OutTradeNo = request.OutTradeNo
	req.Receipt = request.Receipt
	req.SceneInfo = request.SceneInfo
	req.SignType = request.SignType
	req.SpbillCreateIp = request.SpbillCreateIp
	req.TimeExpire = request.TimeExpire
	req.TimeStart = request.TimeStart
	req.TotalFee = request.TotalFee
	sign, err := signValidateByStruct(&req, &request.Key, &request.SignType)
	if err != nil {
		return nil, err
	}
	req.Sign = *sign
	b, err := xml.Marshal(&req)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &PayByPaymentCodeCommitUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CPayByPaymentCodeResponse{}
	err = xml.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ReturnCode != strconv.FormatInt(int64(private.ErrorCodeSuccess), 10) {
		return nil, errors.New(response.ReturnMsg)
	}
	return &response, nil
}

func (this *CPayByPaymentCode) QueryOrder() {
}

func NewPayByPaymentCode(token common.IToken) common.IPayByPaymentCode {
	pay := CPayByPaymentCode{m_token: token}
	return &pay
}
