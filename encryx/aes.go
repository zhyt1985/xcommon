package encryx

import (
	"crypto/aes"
	"crypto/sha1"
	"errors"
)

type Aes struct {
	Key []byte //密钥
	Iv  []byte //iv
}

func NewAes(key string, iv string) *Aes {
	return &Aes{
		Key: []byte(key),
		Iv:  []byte(iv),
	}
}

// content：test123
// encryptKey：123456
// 加密结果为：668C826342B8703D86E8BBF404610499
// 此时就和 java 结果相对应了，解密也一样对 key 加一步处理就行
func (s *Aes) AesSha1EncryptECB(src []byte) ([]byte, error) {
	key, err := AesSha1prng(s.Key, 128) // 比示例一多出这一步
	if err != nil {
		return nil, err
	}

	cipher, _ := aes.NewCipher(generateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(src); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted, nil
}

func (s *Aes) AesSha1DecryptECB(encrypted []byte) ([]byte, error) {
	key, err := AesSha1prng(s.Key, 128) // 比示例一多出这一步
	if err != nil {
		return nil, err
	}

	cipher, _ := aes.NewCipher(generateKey(key))
	decrypted := make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim], nil
}

// 模拟 java SHA1PRNG 处理
func AesSha1prng(keyBytes []byte, encryptLength int) ([]byte, error) {
	hashs := Sha1(Sha1(keyBytes))
	maxLen := len(hashs)
	realLen := encryptLength / 8
	if realLen > maxLen {
		return nil, errors.New("invalid length!")
	}

	return hashs[0:realLen], nil
}

func Sha1(data []byte) []byte {
	h := sha1.New()
	h.Write(data)
	return h.Sum(nil)
}

func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}
