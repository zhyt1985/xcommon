package common

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

/**
 * @Description: 获取token字符串
 * @param secretKey jwt 秘钥
 * @param seconds 过期时间，单位秒
 * @param customClaim 自定义要存储的值
 * @return string 返回的token
 * @return error 错误信息
	说明：如果使用go-zero框架
	go-zero从jwt token解析后会将用户生成token时传入的kv原封不动的放在http.Request的Context中，因此我们可以通过Context就可以拿到你想要的值
 	logx.Infof("userId: %v",l.ctx.Value("userId"))// 这里的key和生成jwt token时传入的key一致
*/
func GetJwtToken(secretKey string, seconds int64, customClaim map[string]interface{}) (string, error) {
	claims := make(jwt.MapClaims)
	iat := time.Now().Unix()
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	for key, value := range customClaim {
		claims[key] = value
	}
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
