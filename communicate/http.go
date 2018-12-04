package communicate

import (
	"github.com/MwlLj/go-wechat/common"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendRequestWithToken(itoken common.IToken, timeoutMS int64, url *string, method *string, payload []byte) (resBody []byte, e error) {
	var err error = nil
	token, err := itoken.GetToken(timeoutMS)
	if err != nil {
		return nil, err
	}
	var req *http.Request = nil
	if payload == nil {
		req, err = http.NewRequest(*method, *url, nil)
	} else {
		reader := strings.NewReader(string(payload))
		req, err = http.NewRequest(*method, *url, reader)
	}
	if err != nil {
		return nil, err
	}
	values := req.URL.Query()
	values.Add(AccessToken, string(token))
	req.URL.RawQuery = values.Encode()
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}
