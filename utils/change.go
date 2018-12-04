package utils

import (
	"github.com/MwlLj/go-wechat/common"
)

type CCommonExt struct {
	ToUserName   common.CData
	FromUserName common.CData
}

func ResXml2DataCommunicate(resXml *common.CWxResXml) *common.CDataCommunicate {
	communicate := common.CDataCommunicate{}
	communicate.ToUserName = string(resXml.ToUserName)
	communicate.FromUserName = string(resXml.FromUserName)
	return &communicate
}

func DataCommunicate2ResXml(communicate *common.CDataCommunicate) *common.CWxResXml {
	resXml := common.CWxResXml{}
	resXml.ToUserName = common.CData(communicate.ToUserName)
	resXml.FromUserName = common.CData(communicate.FromUserName)
	return &resXml
}

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

func Message2ResXml(msg *common.CMessage, ext *CCommonExt) *common.CWxResXml {
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

func ResXml2Event(resXml *common.CWxResXml) *common.CEvent {
	event := common.CEvent{}
	event.CreateTime = resXml.CreateTime
	event.MsgType = string(resXml.MsgType)
	event.Event = string(resXml.Event)
	event.EventKey = string(resXml.EventKey)
	event.Ticket = string(resXml.Ticket)
	event.Latitude = resXml.Latitude
	event.Longitude = resXml.Longitude
	event.Precision = resXml.Precision
	event.Status = string(resXml.Status)
	return &event
}

func Event2ResXml(event *common.CEvent, ext *CCommonExt) *common.CWxResXml {
	resXml := common.CWxResXml{}
	resXml.ToUserName = ext.ToUserName
	resXml.FromUserName = ext.FromUserName
	resXml.CreateTime = event.CreateTime
	resXml.MsgType = common.CData(event.MsgType)
	resXml.Event = common.CData(event.Event)
	resXml.EventKey = common.CData(event.EventKey)
	resXml.Ticket = common.CData(event.Ticket)
	resXml.Latitude = event.Latitude
	resXml.Longitude = event.Longitude
	resXml.Precision = event.Precision
	resXml.Status = common.CData(event.Status)
	return &resXml
}
