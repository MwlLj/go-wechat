package utils

import (
	"github.com/MwlLj/go-wechat/common"
)

func ResXml2Message(resXml *common.CWxResXml) *common.CMessage {
	msg := common.CMessage{}
	msg.CreateTime = resXml.CreateTime
	msg.MsgType = string(resXml.MsgType)
	msg.MsgId = resXml.MsgId
	msg.Content = string(resXml.Content)
	msg.PicUrl = string(resXml.PicUrl)
	msg.MediaId = resXml.MediaId
	msg.Format = string(resXml.Format)
	msg.ThumbMediaId = string(resXml.ThumbMediaId)
	msg.LocationX = resXml.Location_X
	msg.LocationY = resXml.Location_Y
	msg.Scale = resXml.Scale
	msg.Label = string(resXml.Label)
	msg.Title = string(resXml.Title)
	msg.Description = string(resXml.Description)
	msg.Url = string(resXml.Url)
	return &msg
}

type CMessage2ResXmlExt struct {
	ToUserName   common.CData
	FromUserName common.CData
}

func Message2ResXml(msg *common.CMessage, ext *CMessage2ResXmlExt) *common.CWxResXml {
	resXml := common.CWxResXml{}
	resXml.ToUserName = ext.ToUserName
	resXml.FromUserName = ext.FromUserName
	resXml.CreateTime = msg.CreateTime
	resXml.MsgType = common.CData(msg.MsgType)
	resXml.MsgId = msg.MsgId
	resXml.Content = common.CData(msg.Content)
	resXml.PicUrl = common.CData(msg.PicUrl)
	resXml.MediaId = msg.MediaId
	resXml.Format = common.CData(msg.Format)
	resXml.ThumbMediaId = common.CData(msg.ThumbMediaId)
	resXml.Location_X = msg.LocationX
	resXml.Location_Y = msg.LocationY
	resXml.Scale = msg.Scale
	resXml.Label = common.CData(msg.Label)
	resXml.Title = common.CData(msg.Title)
	resXml.Description = common.CData(msg.Description)
	resXml.Url = common.CData(msg.Url)
	return &resXml
}
