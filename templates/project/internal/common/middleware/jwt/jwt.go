/**
 * Created by Wangwei on 2019-06-06 10:52.
 */

package jwt

import (
	"errors"
	"goadmin/internal/common"
	"goadmin/internal/conf"
	"time"

	JwtGo "github.com/dgrijalva/jwt-go"

	"strings"
)

//JWT签名结构
type JWT struct {
	JwtSecret []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Token invalid!")
)

type CustomClaims struct {
	EmployeeId int64  `json:"employee_id"` //员工或用户id
	Account    string `json:"account"`     //员工或用户账号
	RoleId     int64  `json:"role_id"`     //员工角色
	OrgTypeId  int64  `json:"org_type_id"` //机构类型id
	OrgId      int64  `json:"org_id"`      //机构id
	OrgName    string `json:"org_name"`    //机构名称
	ClientId   string `json:"client_id"`   //客户端id
	JwtGo.StandardClaims
}

func (jwt *JWT) getJwtSecret() []byte {
	if len(jwt.JwtSecret) == 0 {
		jwt.JwtSecret = []byte(conf.AppConf.JwtSecret)
	}

	return jwt.JwtSecret
}

// 生成员工/用户登录token
func (jwt *JWT) GenToken(userId int64, account string, roleId, orgTypeId, orgId int64, orgName, clientId string) (string, int64, error) {
	clientId = common.FormatClientId(clientId)
	exp := GetExp(clientId)
	claims := CustomClaims{
		EmployeeId: userId,
		Account:    account,
		ClientId:   clientId,
		RoleId:     roleId,
		OrgTypeId:  orgTypeId,
		OrgId:      orgId,
		OrgName:    orgName,
		StandardClaims: JwtGo.StandardClaims{
			ExpiresAt: exp.Unix(),
			Issuer:    "",
		},
	}

	token := JwtGo.NewWithClaims(JwtGo.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwt.getJwtSecret())
	if err != nil {
		return "", 0, err
	}

	return tokenStr, exp.Unix(), nil
}

// 将token解密
func (jwt *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	token, err := JwtGo.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *JwtGo.Token) (interface{}, error) {
		return jwt.JwtSecret, nil
	})

	if err != nil {
		if ve, ok := err.(*JwtGo.ValidationError); ok {
			if ve.Errors&JwtGo.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&JwtGo.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&JwtGo.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

// 设置token的有效期时间
func GetExp(clientId string) time.Time {
	var exp time.Time
	now := time.Now()
	if strings.HasPrefix(clientId, "webadmin") {
		exp = now.Add(30 * 24 * time.Hour)
	} else if strings.ContainsAny(clientId, "mp") {
		exp = now.Add(365 * 10 * 24 * time.Hour)
	} else if strings.HasPrefix(clientId, "h5") {
		exp = now.Add(365 * 24 * time.Hour)
		//exp = now.AddDate(0,0,365)
	} else {
		exp = now.Add(30 * 24 * time.Hour)
	}

	return exp
}
