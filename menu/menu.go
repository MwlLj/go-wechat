package menu

import (
	"github.com/MwlLj/go-wechat/common"
)

type CMenu struct {
	m_menuData common.CMenuData
}

func (this *CMenu) AddButton(button *common.IButton) {
	this.m_menuData.Buttons = append(this.m_menuData.Buttons, *button)
}

func (this *CMenu) Create() {
	this.m_menuData.Buttons = []common.CButtonData{}
}

func NewMenu() common.IMenu {
	menu := CMenu{}
	return &menu
}
