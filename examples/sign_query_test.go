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

func TestSignQuery(t *testing.T) {
	// Create client from environment variables
	client := CreateClient(t)

	// Create freelancers service
	freelancerService := freelancers.NewService(client)

	// Query sign status
	resp, err := freelancerService.QuerySign(&freelancers.SignQueryRequest{
		Name:       vos.EnvString("SS_FREELANCER_NAME"),
		IdCard:     vos.EnvString("SS_FREELANCER_ID_CARD"),
		Mobile:     vos.EnvString("SS_FREELANCER_MOBILE"),
		ProviderId: vos.EnvInt64("SS_PROVIDER_ID"),
	})
	if err != nil {
		log.Fatalf("Failed to query sign: %v", err)
	}

	// Display sign status
	stateNames := map[freelancers.SignState]string{
		freelancers.SignStateUnsigned:  "Unsigned",
		freelancers.SignStateSigned:    "Signed",
		freelancers.SignStateNotFound:  "Not Found",
		freelancers.SignStatePending:   "Pending",
		freelancers.SignStateFailed:    "Failed",
		freelancers.SignStateCancelled: "Cancelled",
	}

	fmt.Printf("Sign Query Result:\n")
	fmt.Printf("  Name: %s\n", resp.Name)
	fmt.Printf("  ID Card: %s\n", resp.IdCard)
	fmt.Printf("  Mobile: %s\n", resp.Mobile)
	fmt.Printf("  Card No: %s\n", resp.CardNo)
	fmt.Printf("  Provider ID: %d\n", resp.ProviderId)
	fmt.Printf("  State: %s (%d)\n", stateNames[resp.State], resp.State)
	if resp.RetMsg != "" {
		fmt.Printf("  Message: %s\n", resp.RetMsg)
	}
}
