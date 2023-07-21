package web

import (
	"backend/x/identity"
	"backend/x/messages"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"google.golang.org/protobuf/proto"
)

type CookieConfig struct {
	Name          string `yaml:"name"`
	ExpiryTimeSec int    `yaml:"expiry_time_sec"`
	Path          string `yaml:"path"`
	Domain        string `yaml:"domain"`
	Secure        bool   `yaml:"secure"`
	HttpOnly      bool   `yaml:"http_only"`
}

func DeleteCookie(name string) *http.Cookie {
	return &http.Cookie{
		Name:   name,
		Value:  "",
		MaxAge: -1,
	}
}

func CreateCookie(config *CookieConfig, value []byte) (*http.Cookie, error) {
	expires := time.Time{}
	if config.ExpiryTimeSec != 0 {
		expires = time.Now().Add(time.Duration(config.ExpiryTimeSec) * time.Second)
	}
	return &http.Cookie{
		Name:     config.Name,
		Value:    base64.URLEncoding.EncodeToString(value),
		Expires:  expires,
		Path:     config.Path,
		Domain:   config.Domain,
		Secure:   config.Secure,
		HttpOnly: config.HttpOnly,
	}, nil
}

func ParseCookieValue(value string) ([]byte, error) {
	byteValue, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return nil, fmt.Errorf("base64_decode: %w", err)
	}
	return byteValue, nil
}

func CreateEncryptedCookie(config *CookieConfig, identity *identity.Identity, value []byte) (*http.Cookie, error) {
	expires := time.Time{}
	if config.ExpiryTimeSec != 0 {
		expires = time.Now().Add(time.Duration(config.ExpiryTimeSec) * time.Second)
	}
	iv, enc, err := identity.GCMAESEncrypt(value)
	if err != nil {
		return nil, fmt.Errorf("gcm_encrypt: %w", err)
	}
	serializedVal, err := proto.Marshal(&messages.EncryptedCookieValue{
		Iv:         iv,
		EncMessage: enc,
	})
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	return &http.Cookie{
		Name:     config.Name,
		Value:    base64.URLEncoding.EncodeToString(serializedVal),
		Expires:  expires,
		Path:     config.Path,
		Domain:   config.Domain,
		Secure:   config.Secure,
		HttpOnly: config.HttpOnly,
	}, nil
}

func ParseEncryptedCookieValue(identity *identity.Identity, value string) ([]byte, error) {
	serializedMessage, err := base64.URLEncoding.DecodeString(value)
	if err != nil {
		return nil, fmt.Errorf("base64_decode: %w", err)
	}
	encValueMessage := messages.EncryptedCookieValue{}
	if err := proto.Unmarshal(serializedMessage, &encValueMessage); err != nil {
		return nil, fmt.Errorf("unmrshal: %w", err)
	}
	plain, err := identity.GCMAESDecrypt(encValueMessage.Iv, encValueMessage.EncMessage)
	if err != nil {
		return nil, fmt.Errorf("gcm_decrypt: %w", err)
	}
	return plain, nil
}
