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

// PaymentType represents the payment method type.
type PaymentType string

// Payment type constants
const (
	PaymentTypeBankCard PaymentType = "0" // bank card payment
	PaymentTypeAlipay   PaymentType = "1" // Alipay payment
	PaymentTypeWeChat   PaymentType = "2" // WeChat payment
)

// AuthChannel represents the face authentication channel type.
type AuthChannel string

// Auth channel constants
const (
	AuthChannelBaidu     AuthChannel = "01" // 百度云
	AuthChannelAliyun    AuthChannel = "02" // 阿里云
	AuthChannelTencent   AuthChannel = "03" // 腾讯云
	AuthChannelFadada    AuthChannel = "04" // 法大大
	AuthChannelAlipay    AuthChannel = "05" // 支付宝
	AuthChannelVolcano   AuthChannel = "06" // 火山引擎
	AuthChannelHuawei    AuthChannel = "07" // 华为云
	AuthChannelSensetime AuthChannel = "08" // 商汤科技
	AuthChannelMegvii    AuthChannel = "09" // 旷世Face++
	AuthChannelJDCloud   AuthChannel = "10" // 京东智联云
	AuthChannelWechatPay AuthChannel = "11" // 微信支付
	AuthChannelOther     AuthChannel = "12" // 其他活体通道
)

// FaceAuthState represents the face authentication state.
type FaceAuthState string

// Face auth state constants
const (
	FaceAuthStateUnAuth  FaceAuthState = "UN_AUTH" // 未认证
	FaceAuthStateProcess FaceAuthState = "PROCESS" // 认证中
	FaceAuthStateSuccess FaceAuthState = "SUCCESS" // 认证成功
	FaceAuthStateFailed  FaceAuthState = "FAILED"  // 认证失败
	FaceAuthStateExpired FaceAuthState = "EXPIRED" // 认证过期
)

type FunCode struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Function codes
var (
	FunCodePayment           = &FunCode{Code: "6001", Name: "payment"}             // function code for payment
	FunCodePaymentQuery      = &FunCode{Code: "6002", Name: "payment_query"}       // function code for payment query
	FunCodeBalanceQuery      = &FunCode{Code: "6003", Name: "balance_query"}       // function code for balance query
	FunCodeSyncFaceAuth      = &FunCode{Code: "6008", Name: "sync_face_auth"}      // function code for sync face recognition record
	FunCodeFaceAuth          = &FunCode{Code: "6009", Name: "face_auth"}           // function code for face recognition authentication
	FunCodeSignContract      = &FunCode{Code: "6010", Name: "sign_contract"}       // function code for contract signing
	FunCodeSignContractQuery = &FunCode{Code: "6011", Name: "sign_contract_query"} // function code for contract status query
)
