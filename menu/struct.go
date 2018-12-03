package menu

type CButtonData struct {
	SubButton           []CButtonData `json:"sub_button"`
	Type                string        `json:"type"`
	Name                string        `json:"name"`
	Key                 string        `json:"key"`
	Url                 string        `json:"url"`
	MediaId             string        `json:"media_id"`
	MiniProgramAppId    string        `json:"appid"`
	MiniProgramPagePath string        `json:"pagepath"`
}

type CMenuData struct {
	Buttons []CButtonData `json:"button"`
}
