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
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"fmt"
)

// EncryptAES encrypts plaintext using AES-256-ECB mode with PKCS5 padding.
// Returns Base64-encoded ciphertext.
func EncryptAES(plaintext, key string) (string, error) {
	if len(key) != 32 {
		return "", fmt.Errorf("%w: AES-256 key must be 32 bytes", ErrEncryptionFailed)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrEncryptionFailed, err)
	}

	// Apply PKCS5 padding
	plaintextBytes := []byte(plaintext)
	paddedData := pkcs5Padding(plaintextBytes, block.BlockSize())

	// Encrypt using ECB mode (manual block-by-block encryption)
	ciphertext := make([]byte, len(paddedData))
	blockSize := block.BlockSize()
	for i := 0; i < len(paddedData); i += blockSize {
		block.Encrypt(ciphertext[i:i+blockSize], paddedData[i:i+blockSize])
	}

	// Base64 encode the result
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES decrypts Base64-encoded ciphertext using AES-256-ECB mode.
// Returns plaintext after removing PKCS5 padding.
func DecryptAES(ciphertext, key string) (string, error) {
	if len(key) != 32 {
		return "", fmt.Errorf("%w: AES-256 key must be 32 bytes", ErrDecryptionFailed)
	}

	// Base64 decode
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("%w: invalid base64: %v", ErrDecryptionFailed, err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrDecryptionFailed, err)
	}

	if len(ciphertextBytes)%block.BlockSize() != 0 {
		return "", fmt.Errorf("%w: ciphertext is not a multiple of block size", ErrDecryptionFailed)
	}

	// Decrypt using ECB mode (manual block-by-block decryption)
	plaintext := make([]byte, len(ciphertextBytes))
	blockSize := block.BlockSize()
	for i := 0; i < len(ciphertextBytes); i += blockSize {
		block.Decrypt(plaintext[i:i+blockSize], ciphertextBytes[i:i+blockSize])
	}

	// Remove PKCS5 padding
	unpaddedData, err := pkcs5UnPadding(plaintext)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrDecryptionFailed, err)
	}

	return string(unpaddedData), nil
}

// pkcs5Padding applies PKCS5 padding to data.
func pkcs5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// pkcs5UnPadding removes PKCS5 padding from data.
func pkcs5UnPadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, fmt.Errorf("data is empty")
	}

	padding := int(data[length-1])
	if padding > length || padding == 0 {
		return nil, fmt.Errorf("invalid padding")
	}

	// Verify all padding bytes are correct
	for i := length - padding; i < length; i++ {
		if data[i] != byte(padding) {
			return nil, fmt.Errorf("invalid padding bytes")
		}
	}

	return data[:length-padding], nil
}
