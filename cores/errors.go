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

import "fmt"

// APIError represents an error returned by the ServiceShare API.
type APIError struct {
	Code    string
	Message string
}

// Error implements the error interface.
func (e *APIError) Error() string {
	return fmt.Sprintf("API Error [%s]: %s", e.Code, e.Message)
}

// NewAPIError creates a new APIError.
func NewAPIError(code, message string) *APIError {
	return &APIError{
		Code:    code,
		Message: message,
	}
}

// Common SDK errors
var (
	ErrInvalidConfig      = fmt.Errorf("invalid configuration")
	ErrEncryptionFailed   = fmt.Errorf("encryption failed")
	ErrDecryptionFailed   = fmt.Errorf("decryption failed")
	ErrSignatureFailed    = fmt.Errorf("signature generation failed")
	ErrVerificationFailed = fmt.Errorf("signature verification failed")
	ErrRequestFailed      = fmt.Errorf("request failed")
	ErrInvalidResponse    = fmt.Errorf("invalid response")
	ErrInvalidKey         = fmt.Errorf("invalid key format")
)
