/*
 * Copyright 2023 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type RsaService struct {
	ServerRsaKey *rsa.PrivateKey
	ClientRsaKey *rsa.PublicKey
}

func NewRsaService() *RsaService {
	return &RsaService{}
}

// 1. 生成密钥对
func (r *RsaService) GenerateKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	// 生成私钥
	priv, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	return priv, &priv.PublicKey, nil
}

func (r *RsaService) Build(clientKey *rsa.PublicKey, serverKey *rsa.PrivateKey) (err error) {
	r.ServerRsaKey = serverKey
	r.ClientRsaKey = clientKey
	//r.ClientRsaKey = &r.ServerRsaKey.PublicKey
	return nil
}

func (r *RsaService) SavePrivateKey(filename string, priv *rsa.PrivateKey) ([]byte, error) {
	if _, err := os.Stat(filename); err == nil {
		// 文件存在，不进行保存
		fmt.Println("Private key file already exists, skipping save.")
		return nil, nil
	}
	privFile, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer privFile.Close()

	privBytes := x509.MarshalPKCS1PrivateKey(priv)
	if err != nil {
		return nil, err
	}

	err = pem.Encode(privFile, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	})
	return privBytes, err
}

func (r *RsaService) SavePublicKey(filename string, pub *rsa.PublicKey) ([]byte, error) {
	if _, err := os.Stat(filename); err == nil {
		// 文件存在，不进行保存
		fmt.Println("Public key file already exists, skipping save.")
		return nil, nil
	}
	pubFile, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	defer pubFile.Close()

	pubBytes, err := x509.MarshalPKIXPublicKey(pub)
	if err != nil {
		return nil, err
	}

	err = pem.Encode(pubFile, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubBytes,
	})
	return pubBytes, err
}

// 2. 加密过程
func (r *RsaService) EncryptMessage(pub *rsa.PublicKey, message string) ([]byte, error) {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(message))
	if err != nil {
		return nil, err
	}
	return ciphertext, nil
}

// 3. 解密过程
func (r *RsaService) DecryptMessage(priv *rsa.PrivateKey, ciphertext []byte) (string, error) {
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}

// 4. 签名过程
func (r *RsaService) SignMessage(priv *rsa.PrivateKey, message string) ([]byte, error) {
	hashed := sha256.Sum256([]byte(message))
	signature, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashed[:])
	if err != nil {
		return nil, err
	}
	return signature, nil
}

// 5. 验签过程
func (r *RsaService) VerifySignature(pub *rsa.PublicKey, message string, signature []byte) error {
	hashed := sha256.Sum256([]byte(message))
	err := rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], signature)
	return err
}

// Crypt Encrypt the password using crypto/bcrypt
func Crypt(password string) (string, error) {
	// Generate "cost" factor for the bcrypt algorithm
	cost := 5

	// Hash password with bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(hashedPassword), err
}

// VerifyPassword Verify the password is consistent with the hashed password in the database
func VerifyPassword(password, hashedPassword string) (error, bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err, false
	}
	return nil, true
}
