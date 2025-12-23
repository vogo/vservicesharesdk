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

package main

import (
	"fmt"
	"log"

	"github.com/vogo/vogo/vos"
	"github.com/vogo/vservicesharesdk/cores"
	"github.com/vogo/vservicesharesdk/examples/common"
	"github.com/vogo/vservicesharesdk/freelancers"
)

func main() {
	// Create client from environment variables
	client := common.CreateClient()

	// Create freelancers service
	freelancerService := freelancers.NewService(client)

	// Prepare silent sign request
	resp, err := freelancerService.SilentSign(&freelancers.SilentSignRequest{
		Name:        vos.EnvString("SS_FREELANCER_NAME"),
		CardNo:      vos.EnvString("SS_FREELANCER_CARD_NO"),
		IdCard:      vos.EnvString("SS_FREELANCER_ID_CARD"),
		Mobile:      vos.EnvString("SS_FREELANCER_MOBILE"),
		PaymentType: cores.PaymentTypeBankCard, // 0=Bank, 1=Alipay, 2=WeChat
		ProviderId:  vos.EnvString("SS_PROVIDER_ID"),
		IdCardPic1:  vos.EnvString("SS_ID_CARD_FRONT_HEX"), // ID front photo in hex format
		IdCardPic2:  vos.EnvString("SS_ID_CARD_BACK_HEX"),  // ID back photo in hex format
		NotifyUrl:   vos.EnvString("SS_NOTIFY_URL"),        // Optional callback URL
		TagList:     []string{"design", "marketing"},       // Optional skill tags
	})
	if err != nil {
		log.Fatalf("Failed to sign freelancer: %v", err)
	}

	fmt.Printf("Silent Sign Result:\n")
	fmt.Printf("  OtherParam: %s\n", resp.OtherParam)
	fmt.Printf("\nNote: This is an asynchronous operation.\n")
	fmt.Printf("Check the callback URL or use signature query API for final results.\n")
}
