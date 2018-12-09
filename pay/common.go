package pay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func signValidate(kvs *map[string]string, key *string) *string {
	signValue := ""
	var keys []string
	for k, v := range *kvs {
		if k == SignField {
			signValue = v
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
	for k := range keys {
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
	// MD5
	h := md5.New()
	h.Write([]byte(joinStrings))
	cipherStr := h.Sum(nil)
	md5Sign := hex.EncodeToString(cipherStr)
	md5Sign = strings.ToUpper(md5Sign)
	// hmac sha1
	/*
		mac := hmac.New(sha1.New, []byte(SignSha1Key))
		mac.Write([]byte(joinStrings))
		cipherStr := mac.Sum(nil)
		hmacSign := hex.EncodeToString(cipherStr)
		hmacSign = strings.ToUpper(hmacSign)
	*/
	return &md5Sign
}

func getRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
