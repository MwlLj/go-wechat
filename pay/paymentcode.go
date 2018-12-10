package pay

import (
	"github.com/MwlLj/go-wechat/common"
)

type CPayByPaymentCode struct {
	m_token common.IToken
}

func (this *CPayByPaymentCode) Commit(request *common.CPayByPaymentCodeRequest, timeoutMS int64) {
}

func NewPayByPaymentCode(token common.IToken) common.IPayByPaymentCode {
	pay := CPayByPaymentCode{m_token: token}
	return &pay
}
