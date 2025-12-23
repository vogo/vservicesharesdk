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
	"github.com/vogo/vservicesharesdk/examples/common"
	"github.com/vogo/vservicesharesdk/freelancers"
)

func main() {
	// Create client from environment variables
	client := common.CreateClient()

	// Create freelancers service
	freelancerService := freelancers.NewService(client)

	// Query contract status
	resp, err := freelancerService.QueryContract(&freelancers.ContractQueryRequest{
		Name:       vos.EnvString("SS_FREELANCER_NAME"),
		IdCard:     vos.EnvString("SS_FREELANCER_ID_CARD"),
		Mobile:     vos.EnvString("SS_FREELANCER_MOBILE"),
		ProviderId: vos.EnvString("SS_PROVIDER_ID"),
	})
	if err != nil {
		log.Fatalf("Failed to query contract: %v", err)
	}

	// Display contract status
	stateNames := map[freelancers.ContractState]string{
		freelancers.ContractStateUnsigned:  "Unsigned",
		freelancers.ContractStateSigned:    "Signed",
		freelancers.ContractStateNotFound:  "Not Found",
		freelancers.ContractStatePending:   "Pending",
		freelancers.ContractStateFailed:    "Failed",
		freelancers.ContractStateCancelled: "Cancelled",
	}

	fmt.Printf("Contract Query Result:\n")
	fmt.Printf("  Name: %s\n", resp.Name)
	fmt.Printf("  ID Card: %s\n", resp.IdCard)
	fmt.Printf("  Mobile: %s\n", resp.Mobile)
	fmt.Printf("  Card No: %s\n", resp.CardNo)
	fmt.Printf("  Provider ID: %s\n", resp.ProviderId)
	fmt.Printf("  State: %s (%d)\n", stateNames[resp.State], resp.State)
	if resp.RetMsg != "" {
		fmt.Printf("  Message: %s\n", resp.RetMsg)
	}
}
