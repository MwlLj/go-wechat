package token

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"io/ioutil"
	"net/http"
	"time"
)

type CTokenJson struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
}

type CToken struct {
	m_tokenJson          CTokenJson
	m_userInfo           common.CUserInfo
	m_isVaild            bool
	m_updateTokenChannel chan bool
}

func (this *CToken) GetUserInfo() *common.CUserInfo {
	return &this.m_userInfo
}

func (this *CToken) GetToken(timeoutMS int64) (token []byte, e error) {
	var isTimeout bool = false
	if this.m_isVaild == false {
		if timeoutMS <= 0 {
			return nil, errors.New("token is vaild")
		} else {
			t := time.After(time.Duration(timeoutMS) * time.Millisecond)
		end:
			for {
				if this.m_isVaild == true {
					isTimeout = false
					break end
				}
				select {
				case <-t:
					isTimeout = true
					break end
				default:
				}
				time.Sleep(100 * time.Millisecond)
			}
		}
	}
	if isTimeout == true {
		return nil, errors.New("timeout")
	}
	return []byte(this.m_tokenJson.AccessToken), nil
}

func (this *CToken) UpdateToken(timeoutMS int64) (token []byte, e error) {
	this.m_updateTokenChannel <- true
	this.m_isVaild = false
	return this.GetToken(timeoutMS)
}

func (this *CToken) init(info *common.CUserInfo) {
	this.m_userInfo = *info
	this.m_updateTokenChannel = make(chan bool, 1)
	this.m_isVaild = false
	go func() {
		for {
			err := this.sendTokenRequest(info)
			if err != nil || this.m_tokenJson.AccessToken == "" {
				fmt.Println(err)
				time.Sleep(100 * time.Millisecond)
				continue
			}
			select {
			case <-this.m_updateTokenChannel:
				fmt.Println("[INFO] update token request")
			case <-time.After(time.Duration(this.m_tokenJson.ExpiresIn) * time.Second):
				this.m_isVaild = false
				fmt.Println("[INFO] token timeout")
			}
		}
	}()
}

func (this *CToken) sendTokenRequest(info *common.CUserInfo) error {
	url := GetTokenUrl
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		this.m_isVaild = false
		return err
	}
	values := req.URL.Query()
	values.Add(GrantType, ClientCredential)
	values.Add(AppId, info.AppId)
	values.Add(AppSecret, info.AppSecret)
	req.URL.RawQuery = values.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		this.m_isVaild = false
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		this.m_isVaild = false
		return err
	}
	err = json.Unmarshal(body, &this.m_tokenJson)
	if err != nil {
		this.m_isVaild = false
		return err
	}
	if this.m_tokenJson.ErrCode != common.ErrorCodeSuccess {
		this.m_isVaild = false
		return errors.New(this.m_tokenJson.ErrMsg)
	}
	this.m_isVaild = true
	return nil
}

func New(info *common.CUserInfo) common.IToken {
	t := CToken{}
	t.init(info)
	return &t
}
