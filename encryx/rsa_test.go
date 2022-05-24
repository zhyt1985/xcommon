package encryx

import (
	"fmt"
	"testing"
)

func TestRSA(t *testing.T) {
	//服务端生成一对RSA秘钥，私钥放在服务端（不可泄露），公钥下发给客户端。
	prvkey, pubkey := GenRsaKey()
	fmt.Printf("公钥： %s", string(pubkey))
	publicKey, err := BytesToPublicKey(pubkey)
	bytesToPrivateKey, err := BytesToPrivateKey(prvkey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	//客户端使用随机函数生成 key。
	key := "123456789abcdefg"
	//- 客户端使用随机的 key 对传输的数据用AES进行加密。
	encryptCBCContent := AesEncryptCBC([]byte("我是测试数据"), []byte(key))
	fmt.Printf("客户端用 key:%s 加密后的数据： %s", key, string(encryptCBCContent))
	println()
	//- 使用服务端给的公钥对 key进行加密。
	encryptWithPublicKey, err := EncryptWithPublicKey([]byte(key), publicKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("加密后的公钥：%s", string(encryptWithPublicKey))
	println()

	//- 客户端将使用AES加密的数据 以及使用 RSA公钥加密的key 一起发给服务端。
	//- 服务端拿到数据后，先使用私钥对加密的随机key进行解密，
	clientKey, err := DecryptWithPrivateKey(encryptWithPublicKey, bytesToPrivateKey)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("服务端发过来的加密key 用私钥解密后是： %s", string(clientKey))
	println()
	//解密成功即可确定是客户端发来的数据，没有经过他人修改，
	//然后使用解密成功的随机key对使用AES加密的数据进行解密，获取最终的数据。
	decryptCBCContent := AesDecryptCBC(encryptCBCContent, clientKey)
	fmt.Printf("解密后的数据为 ：  %s", string(decryptCBCContent))

}
