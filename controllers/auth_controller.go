package controllers

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"time"
)

type JWTKeys struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var jwtConfig *JWTKeys = nil

func InitJWTConfiguration() *JWTKeys {
	if jwtConfig == nil {
		jwtConfig = &JWTKeys{}
		jwtConfig.privateKey = getPrivateKey()
		jwtConfig.PublicKey = getPublicKey()
	}
	return jwtConfig
}

func (jwtConfig *JWTKeys) GenerateToken(userId string) (string, error) {
	token := jwt.New(jwt.SigningMethodRS512)
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = userId
	tokenString, err := token.SignedString(jwtConfig.privateKey)
	if err != nil {

		panic(err)
		return "", err
	}
	return tokenString, nil
}

func AuthenticateUser(w http.ResponseWriter, r *http.Request, nextMethod http.HandlerFunc) {
	jwtConfig := InitJWTConfiguration()
	token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
		return jwtConfig.PublicKey, nil
	})
	if err != nil || !token.Valid {
		responseError := ErrUnauthorized
		responseError.Error = err
		WriteError(w, responseError)
		return
	}
	nextMethod(w, r)
}

func getPrivateKey() *rsa.PrivateKey {
	privateKeyFile, err := os.Open("keys/demo.rsa")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	privateKeyFile.Close()

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.Open("keys/demo.rsa.pub")
	if err != nil {
		panic(err)
	}

	pemfileinfo, _ := publicKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)

	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)

	data, _ := pem.Decode([]byte(pembytes))

	publicKeyFile.Close()

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
