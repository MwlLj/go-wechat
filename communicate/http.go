package communicate

import (
	"bytes"
	"github.com/MwlLj/go-wechat/common"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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

func UploadFileWithToken(itoken common.IToken, timeoutMS int64, formname *string, filename *string, targetUrl *string) (res []byte, e error) {
	token, err := itoken.GetToken(timeoutMS)
	if err != nil {
		return nil, err
	}
	return UploadFile(token, formname, filename, targetUrl)
}

func UploadFile(token []byte, formname *string, filename *string, targetUrl *string) (res []byte, e error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	fileWriter, err := bodyWriter.CreateFormFile(*formname, *filename)
	if err != nil {
		return nil, err
	}
	fh, err := os.Open(*filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	request, err := http.NewRequest(http.MethodPost, *targetUrl, bodyBuf)
	values := request.URL.Query()
	values.Add(AccessToken, string(token))
	request.URL.RawQuery = values.Encode()
	request.Header.Set("Content-Type", contentType)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return resBody, nil
}
