package identity

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

type Identity struct {
	secretKey []byte
}

func NewIdentity(serverSecret, salt string) *Identity {
	return &Identity{
		secretKey: pbkdf2.Key([]byte(serverSecret), []byte(salt), 480000, 32, sha256.New),
	}
}

func (identity *Identity) GCMAESEncrypt(plain []byte) ([]byte, []byte, error) {
	c, err := aes.NewCipher(identity.secretKey)
	if err != nil {
		return nil, nil, fmt.Errorf("new cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, nil, fmt.Errorf("gcm: %w", err)
	}
	iv := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return nil, nil, fmt.Errorf("read_full: %w", err)
	}
	return iv, gcm.Seal(nil, iv, plain, nil), nil
}

func (identity *Identity) GCMAESDecrypt(iv []byte, enc []byte) ([]byte, error) {
	c, err := aes.NewCipher(identity.secretKey)
	if err != nil {
		return nil, fmt.Errorf("new cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, fmt.Errorf("gcm: %w", err)
	}
	plain, err := gcm.Open(nil, iv, enc, nil)
	if err != nil {
		return nil, fmt.Errorf("gcp_open: %w", err)
	}
	return plain, nil
}
