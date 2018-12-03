package menu

import (
	"github.com/MwlLj/go-wechat/common"
)

type CButton struct {
	m_buttonData common.CButtonData
	m_subButtons []CButton
}

func (this *CButton) AddSubButton(button common.IButton) {
}

func (this *CButton) Data() *common.CButtonData {
}
