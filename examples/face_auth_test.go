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

package examples

import (
	"fmt"
	"log"
	"testing"

	"github.com/vogo/vogo/vos"
	"github.com/vogo/vservicesharesdk/freelancers"
)

func TestFaceAuth(t *testing.T) {
	// Create client from environment variables
	client := CreateClient(t)

	// Create freelancers service
	freelancerService := freelancers.NewService(client)

	// Prepare face auth request
	// Note: User must have a successfully signed contract before initiating face auth
	resp, err := freelancerService.FaceAuth(&freelancers.FaceAuthRequest{
		Name:            vos.EnvString("SS_FREELANCER_NAME"),
		IdCard:          vos.EnvString("SS_FREELANCER_ID_CARD"),
		Mobile:          vos.EnvString("SS_FREELANCER_MOBILE"),
		RedirectUrl:     vos.EnvString("SS_FACE_AUTH_REDIRECT_URL"),      // optional
		RedirectBtnName: vos.EnvString("SS_FACE_AUTH_REDIRECT_BTN_NAME"), // optional
	})
	if err != nil {
		log.Fatalf("Failed to initiate face auth: %v", err)
	}

	fmt.Printf("Face Auth Result:\n")
	fmt.Printf("  H5 URL: %s\n", resp.Url)
	fmt.Printf("\nNote: The H5 URL is valid for one day.\n")
}
