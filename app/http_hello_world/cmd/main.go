package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"learn-go/common/tool"
	"math/big"
	"os"
	"time"
)

var currentPath = tool.ProjectPath + "app/http_hello_world/https/"

func main() {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	// 创建证书模板
	template := x509.Certificate{
		SerialNumber: big.NewInt(2019), // 序列号
		Subject: pkix.Name{
			Organization: []string{"lewaimai"},  // 组织名称
			Country:      []string{"china"},     // 国家
			Province:     []string{"GuangDng"},  // 省份
			Locality:     []string{"ShenZheng"}, // 城市
		},
		NotBefore:             time.Now(),                                                   // 生效时间
		NotAfter:              time.Now().Add(365 * 24 * time.Hour),                         // 过期时间，例如一年后过期
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // 密钥用途
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // 扩展密钥用途，例如服务器认证
		BasicConstraintsValid: true,                                                         // 表示这是一个CA证书
	}
	// 使用模板生成证书
	certBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		panic(err)
	}

	// 将私钥保存到文件
	privateKeyFile, err := os.Create(currentPath + "key.pem")
	if err != nil {
		panic(err)
	}
	defer privateKeyFile.Close()
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		panic(err)
	}

	// 将证书保存到文件
	certOut, err := os.Create(currentPath + "cert.pem")
	if err != nil {
		panic(err)
	}
	defer certOut.Close()
	certPEM := &pem.Block{Type: "CERTIFICATE", Bytes: certBytes}
	pem.Encode(certOut, certPEM)

	fmt.Println("create success")
}
