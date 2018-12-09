package shop

import (
	"encoding/json"
	"errors"
	"github.com/MwlLj/go-wechat/common"
	"github.com/MwlLj/go-wechat/communicate"
	"github.com/MwlLj/go-wechat/private"
	"net/http"
	"strings"
)

type CShop struct {
	m_token common.IToken
}

func (this *CShop) stockSkuList2SkuInfo(stockSkuList *[]common.CStockSkuInfo) *string {
	result := ""
	i := 0
	for _, info := range *stockSkuList {
		item := strings.Join([]string{info.Id, info.Vid}, ":")
		if i == 0 {
			result = item
		} else {
			result = strings.Join([]string{result, item}, ";")
		}
		i++
	}
	return &result
}

func (this *CShop) AddCommodity(request *common.CAddCommodityRequest, timeoutMS int64) (*common.CAddCommodityResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &AddCommodityUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CAddCommodityResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CShop) DeleteCommodity(request *common.CDeleteCommodityRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &DeleteCommodityUrl, &method, nil, nil, b)
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

func (this *CShop) UpdateCommodity(request *common.CUpdateCommodityRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &UpdateCommodityUrl, &method, nil, nil, b)
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

func (this *CShop) GetCommodity(request *common.CGetCommodityRequest, timeoutMS int64) (*common.CGetCommodityResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetCommodityUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetCommodityResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CShop) GetCommodityByStatus(request *common.CGetCommodityByStatusRequest, timeoutMS int64) (*common.CGetCommodityByStatusResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetCommodityByStatusUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetCommodityByStatusResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CShop) UpdateCommodityStatus(request *common.CUpdateCommodityStatusRequest, timeoutMS int64) error {
	b, err := json.Marshal(request)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &UpdateCommodityStatusUrl, &method, nil, nil, b)
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

func (this *CShop) GetSubClassesByClassify(request *common.CGetSubClassesByClassifyRequest, timeoutMS int64) (*common.CGetSubClassesByClassifyResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetSubClassesByClassifyUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetSubClassesByClassifyResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CShop) GetAllSkuByClassify(request *common.CGetAllSkuByClassifyRequest, timeoutMS int64) (*common.CGetAllSkuByClassifyResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetAllSkuByClassifyUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetAllSkuByClassifyResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CShop) GetAllPropertyByClassify(request *common.CGetAllPropertyByClassifyRequest, timeoutMS int64) (*common.CGetAllPropertyByClassifyResponse, error) {
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &GetAllPropertyByClassifyUrl, &method, nil, nil, b)
	if err != nil {
		return nil, err
	}
	response := common.CGetAllPropertyByClassifyResponse{}
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return nil, err
	}
	if response.ErrCode != private.ErrorCodeSuccess {
		return nil, errors.New(response.ErrMsg)
	}
	return &response, nil
}

func (this *CShop) AddStock(request *common.CAddStockRequest, timeoutMS int64) error {
	skuInfo := this.stockSkuList2SkuInfo(&request.SkuInfo)
	req := CAddStockRequest{}
	req.ProductId = request.ProductId
	req.Quantity = request.Quantity
	req.SkuInfo = *skuInfo
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &AddStockUrl, &method, nil, nil, b)
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

func (this *CShop) ReduceStock(request *common.CReduceStockRequest, timeoutMS int64) error {
	skuInfo := this.stockSkuList2SkuInfo(&request.SkuInfo)
	req := CReduceStockRequest{}
	req.ProductId = request.ProductId
	req.Quantity = request.Quantity
	req.SkuInfo = *skuInfo
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}
	method := http.MethodPost
	resBody, err := communicate.SendRequestWithToken(this.m_token, timeoutMS, &ReduceStockUrl, &method, nil, nil, b)
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

func New(token common.IToken) common.IShop {
	shop := CShop{m_token: token}
	return &shop
}
