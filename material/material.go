package material

import (
	"encoding/json"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"github.com/MwlLj/go-wechat/private"
	"net/http"
)

type CMaterial struct {
	m_token common.IToken
}

func (this *CMaterial) CreateTmpMaterial(request *common.CCreateTmpMaterialRequest, timeoutMS int64) (*common.CCreateTmpMaterialResponse, error) {
	params := make(map[string]string)
	params[CreateTmpMaterialParamType] = request.MaterialType
	multi := communicate.CMultiData{}
	multi.Type = communicate.MultiDataTypeFile
	multi.FormName = CreateTmpMaterialFormname
	multi.ValueOrFilename = request.Path
	multis := make([]communicate.CMultiData, 1)
	multis = append(multis, multi)
	resBody, err := communicate.UploadFileWithToken(this.m_token, timeoutMS, &multis, &CreateTmpMaterialUrl, &params, nil)
	if err != nil {
		return nil, err
	}
	response := common.CCreateTmpMaterialResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CMaterial) GetTmpMaterial(request *common.CGetTmpMaterialRequest, timeoutMS int64) (*common.CGetTmpMaterialResponse, error) {
	method := http.MethodGet
	params := make(map[string]string)
	params[GetTmpMaterialParamMediaId] = request.MediaId
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTmpMaterialUrl, &method, &params, nil, nil)
	if err != nil {
		return nil, err
	}
	response := CGetTmpMaterialResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	res := common.CGetTmpMaterialResponse{}
	isFind := false
	for _, url := range GetTmpMaterialResUrls {
		r := response.ResUrl.(map[string]interface{})[url]
		if r != nil {
			res.ResUrl = r.(string)
			isFind = true
			break
		}
	}
	if isFind == false {
		return nil, errors.New("type error")
	}
	return &res, nil
}

func (this *CMaterial) GetTmpHDMaterial(request *common.CGetTmpHDMaterialRequest, timeoutMS int64) (*common.CGetTmpHDMaterialResponse, error) {
	method := http.MethodGet
	params := make(map[string]string)
	params[GetTmpHDMaterialParamMediaId] = request.MediaId
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetTmpHDMaterialUrl, &method, &params, nil, nil)
	if err != nil {
		return nil, err
	}
	response := CGetTmpHDMaterialResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	res := common.CGetTmpHDMaterialResponse{}
	isFind := false
	for _, url := range GetTmpMaterialResUrls {
		r := response.ResUrl.(map[string]interface{})[url]
		if r != nil {
			res.ResUrl = r.(string)
			isFind = true
			break
		}
	}
	if isFind == false {
		return nil, errors.New("type error")
	}
	return &res, nil
}

func (this *CMaterial) AddForeverImgTextMaterial(request *common.CAddForeverImgTextMaterialRequest, timeoutMS int64) (*common.CAddForeverImgTextMaterialResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &AddForeverImgTextMaterialUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CAddForeverImgTextMaterialResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CMaterial) UploadImage(path *string, timeoutMS int64) (*common.CUploadImageResponse, error) {
	multi := communicate.CMultiData{}
	multi.Type = communicate.MultiDataTypeFile
	multi.FormName = UploadImageFormname
	multi.ValueOrFilename = *path
	multis := make([]communicate.CMultiData, 1)
	multis = append(multis, multi)
	resBody, err := communicate.UploadFileWithToken(this.m_token, timeoutMS, &multis, &UploadImageUrl, nil, nil)
	if err != nil {
		return nil, err
	}
	response := common.CUploadImageResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CMaterial) UploadVideo(request *common.CUploadVideoRequest, timeoutMS int64) (*common.CUploadVideoResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &UploadVideoUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CUploadVideoResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CMaterial) AddForeverOtherMaterial(request *common.CAddForeverOtherMaterialRequest, timeoutMS int64) (*common.CAddForeverOtherMaterialResponse, error) {
	params := make(map[string]string)
	params[AddOtherMaterialParamType] = request.MaterialType
	multis := make([]communicate.CMultiData, 1)
	multi := communicate.CMultiData{}
	// file
	multi.Type = communicate.MultiDataTypeFile
	multi.FormName = AddOtherMaterialFormnameMedia
	multi.ValueOrFilename = request.Path
	multis = append(multis, multi)
	// desc
	multi.Type = communicate.MultiDataTypeText
	multi.FormName = AddOtherMaterialFormnameDescription
	desc := CAddForeverOtherMaterialFormnameDesc{}
	desc.Title = request.Title
	desc.Introduction = request.Introduction
	b, err := json.Marshal(&desc)
	if err != nil {
		return nil, err
	}
	multi.ValueOrFilename = string(b)
	multis = append(multis, multi)
	resBody, err := communicate.UploadFileWithToken(this.m_token, timeoutMS, &multis, &AddForeverOtherMaterialUrl, &params, nil)
	if err != nil {
		return nil, err
	}
	response := common.CAddForeverOtherMaterialResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CMaterial) GetForeverMaterial(request *common.CGetForeverMaterialRequest, timeoutMS int64) (*common.CGetForeverMaterialResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetForeverMaterialUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetForeverMaterialResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CMaterial) DeleteForeverMaterial(request *common.CDeleteForeverMaterialRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteForeverMaterialUrl, &method, nil, nil, b)
	if err != nil {
		return err
	}
	response := private.CCommonResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func (this *CMaterial) UpdateForeverImgTextMaterial(request *common.CUpdateForeverImgTextMaterialRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &UpdateForeverImgTextMaterialUrl, &method, nil, nil, b)
	if err != nil {
		return err
	}
	response := private.CCommonResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return errors.New(response.ErrMsg)
	}
	return nil
}

func (this *CMaterial) GetMaterialTotal(timeoutMS int64) (*common.CGetMaterialTotalResponse, error) {
	method := http.MethodGet
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetMaterialTotal, &method, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	response := common.CGetMaterialTotalResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CMaterial) GetMaterialList(request *common.CGetMaterialListRequest, timeoutMS int64) (*common.CGetMaterialListResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetMaterialList, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetMaterialListResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func New(token common.IToken) common.IMaterial {
	media := CMaterial{
		m_token: token,
	}
	return &media
}
