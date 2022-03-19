package jwt

import (
	"bytes"
	"crypto"
	"crypto/hmac"
	_ "github.com/pedroalbanese/gogost/gost34112012256"
	_ "github.com/pedroalbanese/gogost/gost34112012512"
	"errors"
)

// Implements the HMAC-SHA family of signing methods signing methods
type SigningMethodHMAC struct {
	Name string
	Hash crypto.Hash
}

// Specific instances for HS256 and company
var (
	SigningMethodHG256 *SigningMethodHMAC
	SigningMethodHG512 *SigningMethodHMAC
)

func init() {
	// HG256
	SigningMethodHG256 = &SigningMethodHMAC{"HG256", crypto.GOSTR34112012256}
	RegisterSigningMethod(SigningMethodHG256.Alg(), func() SigningMethod {
		return SigningMethodHG256
	})

	// HG512
	SigningMethodHG512 = &SigningMethodHMAC{"HG512", crypto.GOSTR34112012512}
	RegisterSigningMethod(SigningMethodHG512.Alg(), func() SigningMethod {
		return SigningMethodHG512
	})
}

func (m *SigningMethodHMAC) Alg() string {
	return m.Name
}

func (m *SigningMethodHMAC) Verify(signingString, signature string, key interface{}) error {
	if keyBytes, ok := key.([]byte); ok {
		var sig []byte
		var err error
		if sig, err = DecodeSegment(signature); err == nil {
			if !m.Hash.Available() {
				return ErrHashUnavailable
			}

			hasher := hmac.New(m.Hash.New, keyBytes)
			hasher.Write([]byte(signingString))

			if !bytes.Equal(sig, hasher.Sum(nil)) {
				err = errors.New("signature is invalid")
			}
		}
		return err
	}

	return ErrInvalidKey
}

func (m *SigningMethodHMAC) Sign(signingString string, key interface{}) (string, error) {
	if keyBytes, ok := key.([]byte); ok {
		if !m.Hash.Available() {
			return "", ErrHashUnavailable
		}

		hasher := hmac.New(m.Hash.New, keyBytes)
		hasher.Write([]byte(signingString))

		return EncodeSegment(hasher.Sum(nil)), nil
	}

	return "", ErrInvalidKey
}
