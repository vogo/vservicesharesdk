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
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/vogo/vogo/vos"
	"github.com/vogo/vservicesharesdk/cores"
	"github.com/vogo/vservicesharesdk/freelancers"
)

func TestSilentSign(t *testing.T) {
	// Create client from environment variables
	client := CreateClient(t)

	// Create freelancers service
	freelancerService := freelancers.NewService(client)

	frontPath := vos.EnvString("SS_ID_CARD_FRONT_PATH") // ID front photo path
	backPath := vos.EnvString("SS_ID_CARD_BACK_PATH")   // ID back photo path
	frontBytes, err := os.ReadFile(frontPath)
	if err != nil {
		log.Fatalf("failed to read ID card front photo | err: %v", err)
	}
	backBytes, err := os.ReadFile(backPath)
	if err != nil {
		log.Fatalf("failed to read ID card back photo | err: %v", err)
	}

	// Prepare silent sign request
	resp, err := freelancerService.SignContract(&freelancers.SignContractRequest{
		Name:        vos.EnvString("SS_FREELANCER_NAME"),
		CardNo:      vos.EnvString("SS_FREELANCER_CARD_NO"),
		IdCard:      vos.EnvString("SS_FREELANCER_ID_CARD"),
		Mobile:      vos.EnvString("SS_FREELANCER_MOBILE"),
		PaymentType: cores.PaymentTypeBankCard,
		ProviderId:  vos.EnvInt64("SS_PROVIDER_ID"),
		IdCardPic1:  hex.EncodeToString(frontBytes),  // ID front photo path
		IdCardPic2:  hex.EncodeToString(backBytes),   // ID back photo path
		NotifyUrl:   vos.EnvString("SS_NOTIFY_URL"),  // Optional callback URL
		TagList:     []string{"design", "marketing"}, // Optional skill tags
	})
	if err != nil {
		log.Fatalf("Failed to sign freelancer: %v", err)
	}

	fmt.Printf("Silent Sign Result:\n")
	fmt.Printf("  OtherParam: %s\n", resp.OtherParam)
	fmt.Printf("\nNote: This is an asynchronous operation.\n")
	fmt.Printf("Check the callback URL or use signature query API for final results.\n")
}
