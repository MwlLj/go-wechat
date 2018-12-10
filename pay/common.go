package pay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/MwlLj/go-wechat/common"
	"math/rand"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

var _ = fmt.Println

func struct2map(data interface{}) (*[]string, *map[string]string) {
	kvs := make(map[string]string)
	var keys []string
	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		field := t.Field(i)
		tag := string(field.Tag)
		reg, err := regexp.Compile(`xml:"(.*?)"`)
		if err != nil {
			continue
		}
		rs := reg.FindStringSubmatch(tag)
		if len(rs) != 2 {
			continue
		}
		tagName := rs[1]
		typeString := field.Type.String()
		va := v.Field(i).Interface()
		var value string
		if typeString == "string" {
			value = va.(string)
		} else if typeString == "int64" {
			value = strconv.FormatInt(va.(int64), 10)
		} else if typeString == "int" || typeString == "int32" {
			value = strconv.Itoa(va.(int))
		} else if typeString == "float64" {
			value = strconv.FormatFloat(va.(float64), 'f', 30, 32)
		} else if typeString == "bool" {
			value = strconv.FormatBool(va.(bool))
		} else if typeString == "uint64" {
			value = strconv.FormatUint(va.(uint64), 10)
		} else if typeString == "unit" || typeString == "uint32" {
			value = strconv.Itoa(va.(int))
		}
		if tagName == SignField || value == "" {
			continue
		}
		keys = append(keys, tagName)
		kvs[tagName] = value
	}
	return &keys, &kvs
}

func signValidateByStruct(request *CPayByPaymentCodeRequest, key *string, encrypyType *string) (*string, error) {
	keys, kvs := struct2map(*request)
	joinStrings := ""
	sort.Strings(*keys)
	// join kvs
	i := 0
	for _, k := range *keys {
		item := strings.Join([]string{k, (*kvs)[k]}, "=")
		if i == 0 {
			joinStrings = item
		} else {
			joinStrings = strings.Join([]string{joinStrings, item}, "&")
		}
		i++
	}
	// join key
	joinStrings = strings.Join([]string{joinStrings, "&", SignKeyField, "=", *key}, "/")
	var sign string
	// MD5
	if *encrypyType == common.SignEncrypyTypeHMACSHA256 {
		// hmac sha256
		mac := hmac.New(sha256.New, []byte(*key))
		mac.Write([]byte(joinStrings))
		cipherStr := mac.Sum(nil)
		hmacSign := hex.EncodeToString(cipherStr)
		sign = strings.ToUpper(hmacSign)
	} else if *encrypyType == common.SignEncrypyTypeMD5 {
		h := md5.New()
		h.Write([]byte(joinStrings))
		cipherStr := h.Sum(nil)
		md5Sign := hex.EncodeToString(cipherStr)
		sign = strings.ToUpper(md5Sign)
	} else {
		return nil, errors.New("encrypy type not support")
	}
	return &sign, nil
}

func signValidate(kvs *map[string]string, key *string, encrypyType *string) (*string, error) {
	// signValue := ""
	var keys []string
	for k, v := range *kvs {
		if k == SignField {
			// signValue = v
		} else {
			if v != "" {
				keys = append(keys, k)
			}
		}
	}
	joinStrings := ""
	sort.Strings(keys)
	// join kvs
	i := 0
	for _, k := range keys {
		item := strings.Join([]string{k, (*kvs)[k]}, "=")
		if i == 0 {
			joinStrings = item
		} else {
			joinStrings = strings.Join([]string{joinStrings, item}, "&")
		}
		i++
	}
	// join key
	joinStrings = strings.Join([]string{joinStrings, "&", SignKeyField, "=", *key}, "/")
	var sign string
	// MD5
	if *encrypyType == common.SignEncrypyTypeHMACSHA256 {
		// hmac sha256
		mac := hmac.New(sha256.New, []byte(*key))
		mac.Write([]byte(joinStrings))
		cipherStr := mac.Sum(nil)
		hmacSign := hex.EncodeToString(cipherStr)
		sign = strings.ToUpper(hmacSign)
	} else if *encrypyType == common.SignEncrypyTypeMD5 {
		h := md5.New()
		h.Write([]byte(joinStrings))
		cipherStr := h.Sum(nil)
		md5Sign := hex.EncodeToString(cipherStr)
		sign = strings.ToUpper(md5Sign)
	} else {
		return nil, errors.New("encrypy type not support")
	}
	return &sign, nil
}

func genRandomString(length int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}
