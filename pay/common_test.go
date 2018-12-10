package pay

import (
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"testing"
)

func TestSignValidateByStruct(t *testing.T) {
	request := CPayByPaymentCodeRequest{}
	request.Body = "hello"
	key := "test"
	sign, err := signValidateByStruct(&request, &key, &common.SignEncrypyTypeMD5)
	if err != nil {
		return
	}
	fmt.Println(*sign)
}

func TestGenRandomString(t *testing.T) {
	fmt.Println(string(genRandomString(32)))
}
