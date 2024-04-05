package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	// ID トークンから、ヘッダー・ペイロードを入手するプログラム

	// 変数idTokenには、先ほど取得した自分のIDトークンを代入
	idToken := `eyJhbGciOiJSUzI1NiIsImtpZCI6IjkzNGE1ODE2NDY4Yjk1NzAzOTUzZDE0ZTlmMTVkZjVkMDlhNDAxZTQiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJodHRwczovL2FjY291bnRzLmdvb2dsZS5jb20iLCJhenAiOiIzMDE0MjUyMzk3NjgtMG81c3Zva3YwZHY3ZTZkYTVtOTR2OTQ0NmQ2aGFrcWguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJhdWQiOiIzMDE0MjUyMzk3NjgtMG81c3Zva3YwZHY3ZTZkYTVtOTR2OTQ0NmQ2aGFrcWguYXBwcy5nb29nbGV1c2VyY29udGVudC5jb20iLCJzdWIiOiIxMDgyMDU5NjA5ODk5NjkyNTgwMTUiLCJub25jZSI6IjExMTExIiwibmJmIjoxNzEyMjQxNTM0LCJuYW1lIjoia2lub2tpa2kiLCJwaWN0dXJlIjoiaHR0cHM6Ly9saDMuZ29vZ2xldXNlcmNvbnRlbnQuY29tL2EvQUNnOG9jS1dZYW85Y09FQzB2c0Z0c2NNN3dDaVVUV0NtdTB2bXNMajM3RmVpdXBvQTdBclJtQlU9czk2LWMiLCJnaXZlbl9uYW1lIjoia2lub2tpa2kiLCJpYXQiOjE3MTIyNDE4MzQsImV4cCI6MTcxMjI0NTQzNCwianRpIjoiNWZmMmExNDE3MmFkOTQ4MGExMGUyYzgyNDMxNjliZGYwYzY4YjY5ZiJ9.nfAKbO6hxiJZAid_cY-ZlRJntqpxBbU4kqyOM9DcRBPYEjC16Jx4bqxoBMmnAZNYeyFGULK3qZ6v20BskSSUaYjLQMvoJjb0RU4cJI3mGbeQenyfQQcpqsLVpDys0Z1GZnb-B8Bhn-zyreX4NHigqATdMTzHJQevx_dXB9qRITjQDBMzFiCa1Ev3gmrVGAc2lziiEkNCBLcfUEvgaZ-x33dJDAeyZjT6ev6NAq7fb2n-pqMlss0u4DSmAyX6S73cd-AB_A2_ZM6wXswhjbXKx0j6PXbUOZ-_TOT7IvoAKH8grpZAbbhC7ZhSDUrIJ5Uyg2Vbr9r85Lne8ko7ben71g`

	dataArray := strings.Split(idToken, ".")
	header, payload, sig := dataArray[0], dataArray[1], dataArray[2]

	// headerをbase64 decodeする
	headerData, err := base64.RawURLEncoding.DecodeString(header)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// payloadをbase64 decodeする
	payloadData, err := base64.RawURLEncoding.DecodeString(payload)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	// 公開鍵構造体を作成
	N := "uhWRpJ3PNZaiBmq3P91A6QB0b28LeQvV-HI0TAEcN5nffQPm94w-hY2S6mThb7xXLCGHcP3bhpWl31giZJFlvzHe6db-TsPl8HSLgLIjMbMT8iYWqZPa2eodijEJrkO6SPex5jHLzSwGsoRdSfW8hFeTFQk8xtPXm7GlEEo9mFEKUAaArT9acdE8h53VR7ZkJkipiLCtx0rhySA2W4rEAcinLG3ApG709pOw6sVjA2IAQmZVYrfQ7curmFqKWL_F534kDhQJL2hMdrubhHcqCxetyi_U7WDWDkYCJ_CetjDsI0yfwB2sR01vn6LuDDo6ho8pWJcHOOvXYUnSMFAlew"
	E := "AQAB"

	dn, _ := base64.RawURLEncoding.DecodeString(N)
	de, _ := base64.RawURLEncoding.DecodeString(E)

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(dn),
		E: int(new(big.Int).SetBytes(de).Int64()),
	}

	// 検証するデータ
	message := sha256.Sum256([]byte(header + "." + payload))

	// 署名をbase64 decodeする
	sigData, err := base64.RawURLEncoding.DecodeString(sig)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	if err := rsa.VerifyPKCS1v15(pk, crypto.SHA256, message[:], sigData); err != nil {
		fmt.Println("invalid token")
	} else {
		fmt.Println("valid token")
		fmt.Println("header: ", string(headerData))
		fmt.Println("payload: ", string(payloadData))
	}
}
