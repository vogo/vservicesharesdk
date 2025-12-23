/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cores

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"strings"
)

// Sign signs data using RSA-SHA1 and returns Base64-encoded signature.
func Sign(data string, privateKey *rsa.PrivateKey) (string, error) {
	if privateKey == nil {
		return "", fmt.Errorf("%w: private key is nil", ErrSignatureFailed)
	}

	// Calculate SHA1 hash
	h := sha1.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	// Sign with RSA
	signature, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, hashed)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrSignatureFailed, err)
	}

	// Base64 encode
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Verify verifies RSA-SHA1 signature.
func Verify(data, signatureStr string, publicKey *rsa.PublicKey) error {
	if publicKey == nil {
		return fmt.Errorf("%w: public key is nil", ErrVerificationFailed)
	}

	// Base64 decode signature
	signature, err := base64.StdEncoding.DecodeString(signatureStr)
	if err != nil {
		return fmt.Errorf("%w: invalid base64 signature: %v", ErrVerificationFailed, err)
	}

	// Calculate SHA1 hash
	h := sha1.New()
	h.Write([]byte(data))
	hashed := h.Sum(nil)

	// Verify signature
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA1, hashed, signature)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrVerificationFailed, err)
	}

	return nil
}

// ParsePrivateKey parses RSA private key from PEM format (PKCS8) or raw base64.
func ParsePrivateKey(pemKey string) (*rsa.PrivateKey, error) {
	// Normalize line endings and format
	pemKey = strings.ReplaceAll(pemKey, "\r\n", "\n")
	pemKey = strings.TrimSpace(pemKey)

	var keyBytes []byte

	// Try parsing as PEM first
	block, _ := pem.Decode([]byte(pemKey))
	if block != nil {
		keyBytes = block.Bytes
	} else {
		// If PEM decode fails, try raw base64
		decoded, err := base64.StdEncoding.DecodeString(pemKey)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to decode PEM block or base64: %v", ErrInvalidKey, err)
		}
		keyBytes = decoded
	}

	// Try PKCS8 format first
	privateKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err == nil {
		if rsaKey, ok := privateKey.(*rsa.PrivateKey); ok {
			return rsaKey, nil
		}
		return nil, fmt.Errorf("%w: not an RSA private key", ErrInvalidKey)
	}

	// Try PKCS1 format as fallback
	rsaKey, err := x509.ParsePKCS1PrivateKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidKey, err)
	}

	return rsaKey, nil
}

// ParsePublicKey parses RSA public key from PEM format or raw base64.
func ParsePublicKey(pemKey string) (*rsa.PublicKey, error) {
	// Normalize line endings and format
	pemKey = strings.ReplaceAll(pemKey, "\r\n", "\n")
	pemKey = strings.TrimSpace(pemKey)

	var keyBytes []byte

	// Try parsing as PEM first
	block, _ := pem.Decode([]byte(pemKey))
	if block != nil {
		keyBytes = block.Bytes
	} else {
		// If PEM decode fails, try raw base64
		decoded, err := base64.StdEncoding.DecodeString(pemKey)
		if err != nil {
			return nil, fmt.Errorf("%w: failed to decode PEM block or base64: %v", ErrInvalidKey, err)
		}
		keyBytes = decoded
	}

	// Try PKIX format
	publicKey, err := x509.ParsePKIXPublicKey(keyBytes)
	if err == nil {
		if rsaKey, ok := publicKey.(*rsa.PublicKey); ok {
			return rsaKey, nil
		}
		return nil, fmt.Errorf("%w: not an RSA public key", ErrInvalidKey)
	}

	// Try PKCS1 format as fallback
	rsaKey, err := x509.ParsePKCS1PublicKey(keyBytes)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInvalidKey, err)
	}

	return rsaKey, nil
}
