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
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"

	"github.com/vogo/vogo/vlog"
)

// Client represents the ServiceShare API client.
type Client struct {
	config     *Config
	httpClient *http.Client
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

// NewClient creates a new ServiceShare API client.
func NewClient(config *Config) (*Client, error) {
	// Validate configuration
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Parse private key
	privateKey, err := ParsePrivateKey(config.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}

	// Parse platform public key
	publicKey, err := ParsePublicKey(config.PlatformPublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse platform public key: %w", err)
	}

	// Create HTTP client
	httpClient := &http.Client{
		Timeout: config.Timeout,
	}

	return &Client{
		config:     config,
		httpClient: httpClient,
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// generateRequestID generates a unique request ID using timestamp and random number.
func (c *Client) generateRequestID() string {
	timestamp := time.Now().Unix()
	random := rand.Intn(1000000)
	return fmt.Sprintf("%d%06d", timestamp, random)
}

// Do executes an API request with encryption and signing.
// Returns decrypted response data as JSON string.
func (c *Client) Do(funCode string, reqData interface{}) (string, error) {
	// 1. Generate unique request ID
	reqId := c.generateRequestID()

	// 2. Marshal request data to JSON
	reqDataJSON, err := json.Marshal(reqData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request data: %w", err)
	}

	vlog.Infof("service share api request | funCode: %s | reqId: %s | reqData: %s", funCode, reqId, string(reqDataJSON))

	// 3. Encrypt request data with AES
	encryptedData, err := EncryptAES(string(reqDataJSON), c.config.AesKey)
	if err != nil {
		return "", fmt.Errorf("failed to encrypt request data: %w", err)
	}

	// 4. Create request message
	requestMsg := &RequestMessage{
		ReqId:   reqId,
		FunCode: funCode,
		MerId:   c.config.MerchantID,
		Version: c.config.Version,
		ReqData: encryptedData,
	}

	// 5. Sign the encrypted data
	signature, err := Sign(encryptedData, c.privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign request: %w", err)
	}
	requestMsg.Sign = signature

	// 6. Marshal request message to JSON
	requestJSON, err := requestMsg.ToJSON()
	if err != nil {
		return "", err
	}

	// 7. Send HTTP POST request
	req, err := http.NewRequest("POST", c.config.BaseURL, bytes.NewReader(requestJSON))
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	req.Header.Set("Accept", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("%w: %v", ErrRequestFailed, err)
	}
	defer resp.Body.Close()

	// 8. Read response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("%w: failed to read response: %v", ErrRequestFailed, err)
	}

	// Check HTTP status
	if resp.StatusCode != http.StatusOK {
		vlog.Errorf("service share api response not ok | funCode: %s | reqId: %s | status_code: %d | respBody: %s", funCode, reqId, resp.StatusCode, string(respBody))

		return "", fmt.Errorf("%w: HTTP %d", ErrRequestFailed, resp.StatusCode)
	}

	vlog.Infof("service share api response | funCode: %s | reqId: %s | respBody: %s", funCode, reqId, string(respBody))

	// 9. Parse response message
	responseMsg, err := ParseResponseMessage(respBody)
	if err != nil {
		vlog.Errorf("service share api response parse failed | funCode: %s | reqId: %s | respBody: %s | err: %v", funCode, reqId, string(respBody), err)

		return "", fmt.Errorf("%w: %v", ErrInvalidResponse, err)
	}

	// 10. Check API response code
	if !responseMsg.IsSuccess() {
		return "", responseMsg.GetError()
	}

	// 11. Verify response signature (if present)
	if responseMsg.Sign != "" && responseMsg.ResData != "" {
		if err := Verify(responseMsg.ResData, responseMsg.Sign, c.publicKey); err != nil {
			return "", fmt.Errorf("response signature verification failed: %w", err)
		}
	}

	// 12. Decrypt response data (if present)
	if responseMsg.ResData == "" {
		return "", nil
	}

	decryptedData, err := DecryptAES(responseMsg.ResData, c.config.AesKey)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt response data: %w", err)
	}

	vlog.Infof("service share api response decrypt | funCode: %s | reqId: %s | respBody: %s | decryptedData: %s", funCode, reqId, string(respBody), decryptedData)

	return decryptedData, nil
}

// VerifyAndDecryptNotification verifies the signature and decrypts the notification data.
// It accepts the raw JSON body of the notification request.
func (c *Client) VerifyAndDecryptNotification(body []byte) (string, error) {
	// 1. Parse request message
	var reqMsg RequestMessage
	if err := json.Unmarshal(body, &reqMsg); err != nil {
		return "", fmt.Errorf("failed to unmarshal notification envelope: %w", err)
	}

	// 2. Verify signature
	// The notification is signed by the platform, so we verify with the platform's public key.
	if reqMsg.Sign == "" {
		return "", fmt.Errorf("missing signature in notification")
	}
	// Note: The signature is generated based on the encrypted ReqData
	if err := Verify(reqMsg.ReqData, reqMsg.Sign, c.publicKey); err != nil {
		return "", fmt.Errorf("notification signature verification failed: %w", err)
	}

	// 3. Decrypt data
	if reqMsg.ReqData == "" {
		return "", nil
	}

	decryptedData, err := DecryptAES(reqMsg.ReqData, c.config.AesKey)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt notification data: %w", err)
	}

	return decryptedData, nil
}
