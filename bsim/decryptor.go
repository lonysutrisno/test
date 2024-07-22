package bsim

import (
	"bytes"
	"context"
	"crypto/des"
	"encoding/hex"
	"strconv"
)

type desDecryptor struct {
	secretKey string
}

func NewDesDecryptor() DesDecryptor {
	return &desDecryptor{
		secretKey: "3816c875da2f2540",
	}
}

type DesDecryptor interface {
	Decrypt(ctx context.Context, val string) (string, error)
	Encrypt(ctx context.Context, val string) (string, error)
}

func (d *desDecryptor) Decrypt(ctx context.Context, val string) (string, error) {
	return d.desDecrypt(val)
}

func (d *desDecryptor) Encrypt(ctx context.Context, val string) (string, error) {
	return d.desEncrypt(val)
}

func (d *desDecryptor) desDecrypt(val string) (string, error) {
	crypted, err := d.getBytes(val)
	if err != nil {
		return "", err
	}

	keyByte, err := d.getBytes(d.secretKey)
	if err != nil {
		return "", err
	}

	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	if len(crypted)%bs != 0 {
		return "", err
	}
	origData := make([]byte, len(crypted))
	dst := origData
	for len(crypted) > 0 {
		block.Decrypt(dst, crypted[:bs])
		crypted = crypted[bs:]
		dst = dst[bs:]
	}
	origData = d.pkcs5UnPadding(origData)
	return string(origData), nil
}

func (d *desDecryptor) desEncrypt(val string) (string, error) {
	keyByte, err := d.getBytes(d.secretKey)
	if err != nil {
		return "", err
	}

	block, err := des.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	origData := d.pkcs5Padding([]byte(val), bs)
	crypted := make([]byte, len(origData))
	dst := crypted
	for len(origData) > 0 {
		block.Encrypt(dst, origData[:bs])
		origData = origData[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(crypted), nil
}

func (*desDecryptor) pkcs5Padding(origData []byte, blockSize int) []byte {
	padding := blockSize - len(origData)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(origData, padText...)
}

func (*desDecryptor) pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func (*desDecryptor) getBytes(str string) ([]byte, error) {
	strbyte := make([]byte, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		in, err := strconv.ParseUint(str[i:i+2], 16, 8)
		if err != nil {
			return nil, err
		}
		strbyte[i/2] = byte(in)
	}
	return strbyte, nil
}
