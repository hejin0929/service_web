package pkg

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-kratos/kratos/v2/errors"
	"strconv"
	"time"
)

const (
	reason string = "UNAUTHORIZED"
)

var (
	ErrMissingJwtToken        = errors.Unauthorized(reason, "JWT token is missing")
	ErrMissingKeyFunc         = errors.Unauthorized(reason, "keyFunc is missing")
	ErrTokenInvalid           = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired           = errors.Unauthorized(reason, "JWT token has expired")
	ErrTokenParseFail         = errors.Unauthorized(reason, "Fail to parse JWT token ")
	ErrUnSupportSigningMethod = errors.Unauthorized(reason, "Wrong signing method")
	ErrWrongContext           = errors.Unauthorized(reason, "Wrong context for middleware")
	ErrNeedTokenProvider      = errors.Unauthorized(reason, "Token provider is missing")
	ErrSignToken              = errors.Unauthorized(reason, "Can not sign token.Is the key correct?")
	ErrGetKey                 = errors.Unauthorized(reason, "Can not get key while signing token")
)

// 自定义中间件，用于验证 JWT Token
//func JWTAuthMiddleware(secretKey string) middleware.Middleware {
//	return func(handler middleware.Handler) middleware.Handler {
//		return func(ctx context.Context, req interface{}) (interface{}, error) {
//
//			// 从 HTTP 请求的 Header 中获取 Token
//			header, ok := http.FromServerContext(ctx)
//			if !ok {
//				return nil, kratos.ErrInternal
//			}
//			tokenString := header.Request.Header.Get("Authorization")
//
//			// 验证 Token
//			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//				// 指定验证的算法和密钥
//				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//				}
//				return []byte(secretKey), nil
//			})
//			if err != nil || !token.Valid {
//				return nil, kratos.ErrUnauthorized
//			}
//
//			// 验证通过，继续处理请求
//			return handler(ctx, req)
//		}
//	}
//}

// JWTClaims 定义 JWTClaims 结构体作为 JWT Token 的 Claims
type JWTClaims struct {
	jwt.StandardClaims
	UserID   string `json:"user_id"`
	Username string `json:"username"`
}

// GenerateToken 生成 JWT Token
func GenerateToken(userID, username string) (string, error) {
	// 创建一个新的 JWTClaims 对象，设置有效期和自定义的用户信息
	claims := &JWTClaims{
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Add(time.Hour).Unix(), // 设置有效期为 1 小时
			Issuer:    username,
		},
		UserID:   userID,
		Username: username,
	}

	// 使用 HMAC 签名方法创建 JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用自定义的密钥对 Token 进行签名
	secretKey := []byte("password-hello-word")

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ParseToken 解析和验证 Token
func ParseToken(tokenString, secretKey string) (jwt.MapClaims, error) {
	// 解析 Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 指定验证的算法和密钥
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 Token 是否有效
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// 提取声明（Claims）
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse token claims")
	}

	// 从声明中提取时间戳
	expirationTimeFloat64, ok := claims["exp"].(float64)
	if !ok {
		return nil, fmt.Errorf("failed to extract expiration time")
	}

	expirationTime := time.Unix(int64(expirationTimeFloat64), 0).Unix()

	if time.Now().Unix() > expirationTime {
		return nil, fmt.Errorf("token已过期")
	}
	return claims, nil
}

// UseUserID 在业务逻辑层使用上下文中的用户ID
func UseUserID(ctx context.Context) (id int, err error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok {
		// 无法获取用户ID
		return 0, fmt.Errorf("ID为空")
	}

	return strconv.Atoi(userID)
	// 使用用户ID进行业务逻辑处理
	// ...
}
