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

var (
	MultiDataTypeFile string = "file"
	MultiDataTypeText string = "text"
)

type CMultiData struct {
	Type            string
	FormName        string
	ValueOrFilename string
}

func SendRequestWithToken(itoken common.IToken, timeoutMS int64, url *string, method *string, params *map[string]string, headers *map[string]string, payload []byte) (resBody []byte, e error) {
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
	if params != nil {
		for k, v := range *params {
			values.Add(k, v)
		}
	}
	req.URL.RawQuery = values.Encode()
	if headers != nil {
		for k, v := range *headers {
			req.Header.Add(k, v)
		}
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func UploadFileWithToken(itoken common.IToken, timeoutMS int64, multiData *[]CMultiData, targetUrl *string, params *map[string]string, headers *map[string]string) (res []byte, e error) {
	token, err := itoken.GetToken(timeoutMS)
	if err != nil {
		return nil, err
	}
	return SendMultiFormData(token, multiData, targetUrl, params, headers)
}

func SendMultiFormData(token []byte, multiData *[]CMultiData, targetUrl *string, params *map[string]string, headers *map[string]string) (res []byte, e error) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	for _, multi := range *multiData {
		switch multi.Type {
		case MultiDataTypeFile:
			fileWriter, err := bodyWriter.CreateFormFile(multi.FormName, multi.ValueOrFilename)
			if err != nil {
				return nil, err
			}
			fh, err := os.Open(multi.ValueOrFilename)
			if err != nil {
				return nil, err
			}
			defer fh.Close()
			_, err = io.Copy(fileWriter, fh)
			if err != nil {
				return nil, err
			}
		case MultiDataTypeText:
			valueWriter, err := bodyWriter.CreateFormField(multi.FormName)
			if err != nil {
				return nil, err
			}
			valueReader := bytes.NewReader([]byte(multi.ValueOrFilename))
			_, err = io.Copy(valueWriter, valueReader)
			if err != nil {
				return nil, err
			}
		}
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	request, err := http.NewRequest(http.MethodPost, *targetUrl, bodyBuf)
	values := request.URL.Query()
	values.Add(AccessToken, string(token))
	if params != nil {
		for k, v := range *params {
			values.Add(k, v)
		}
	}
	request.URL.RawQuery = values.Encode()
	request.Header.Set("Content-Type", contentType)
	if headers != nil {
		for k, v := range *headers {
			request.Header.Add(k, v)
		}
	}
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
