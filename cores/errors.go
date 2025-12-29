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

// Is checks if the target error is an APIError with the same code.
func (e *APIError) Is(target error) bool {
	t, ok := target.(*APIError)
	if !ok {
		return false
	}
	return e.Code == t.Code
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

// API Business Errors
var (
	ErrApiSuccess                      = NewAPIError("0000", "当前请求处理成功")
	ErrApiUnknown                      = NewAPIError("6000", "当前请求处理未明，请核实")
	ErrApiParamError                   = NewAPIError("6001", "参数错误")
	ErrApiInvalidAmount                = NewAPIError("6002", "无效交易金额")
	ErrApiCustomerNotFound             = NewAPIError("6003", "客户信息不存在")
	ErrApiCustomerStatusNotOpen        = NewAPIError("6004", "客户状态未开通")
	ErrApiCustomerKeyEmpty             = NewAPIError("6005", "客户秘钥为空")
	ErrApiSignVerifyFailed             = NewAPIError("6006", "请求数据验签失败")
	ErrApiDecryptFailed                = NewAPIError("6007", "请求数据解密失败")
	ErrApiMerchantBlacklisted          = NewAPIError("6008", "商户在黑名单不允许交易")
	ErrApiNoRiskControlInfo            = NewAPIError("6009", "无客户风控信息")
	ErrApiAccountInvalid               = NewAPIError("6010", "无客户账户信息或账户状态无效")
	ErrApiIPNotWhitelisted             = NewAPIError("6011", "客户请求地址未配置白名单")
	ErrApiBatchNoDuplicate             = NewAPIError("6012", "客户批次号重复,请确认批次信息")
	ErrApiAmountLimitExceeded          = NewAPIError("6013", "付款金额超限")
	ErrApiSaveFailed                   = NewAPIError("6014", "信息入库失败")
	ErrApiFeeCalculationError          = NewAPIError("6015", "计算客户手续费出错或客户手续费率不存在")
	ErrApiAlreadySigned                = NewAPIError("6016", "该用户信息已经做过签约")
	ErrApiPaymentMethodNotConfigured   = NewAPIError("6017", "客户付款方式未配置")
	ErrApiPermissionDenied             = NewAPIError("6018", "客户未开通该权限")
	ErrApiInsufficientBalance          = NewAPIError("6019", "商户余额不足")
	ErrApiOrderNotFound                = NewAPIError("6020", "未查询到订单")
	ErrApiNotSignedWithServiceCompany  = NewAPIError("6021", "客户未签约此落地服务公司")
	ErrApiSignAuthFailed               = NewAPIError("6022", "签约信息鉴权失败")
	ErrApiBillFileNotFound             = NewAPIError("6023", "对账文件不存在")
	ErrApiNameEmpty                    = NewAPIError("6024", "姓名不能为空")
	ErrApiIdCardEmpty                  = NewAPIError("6025", "身份证号不能为空")
	ErrApiServiceIdEmpty               = NewAPIError("6026", "服务商 Id 不能为空")
	ErrApiNotSignedWithServiceProvider = NewAPIError("6027", "用户未在该服务商签约")
	ErrApiPlatformProviderNotFound     = NewAPIError("6028", "未查询到对应的平台服务商")
	ErrApiPlatformProviderUnavailable  = NewAPIError("6029", "该平台服务商不可用")
	ErrApiCustomerIdEmpty              = NewAPIError("6030", "客户id不能为空")
	ErrApiBatchNoEmpty                 = NewAPIError("6031", "客户批次号不能为空")
	ErrApiBatchNoNotFound              = NewAPIError("6032", "该客户批次号不存在")
	ErrApiOrderNoNotFound              = NewAPIError("6033", "客户订单号或者订单流水号不存在")
	ErrApiTotalCountMismatch           = NewAPIError("6034", "付款总笔数和明细不一致")
	ErrApiTotalAmountMismatch          = NewAPIError("6035", "付款总金额和明细不一致")
	ErrApiMultiServiceProviders        = NewAPIError("6036", "批量付款只能选择一个服务商")
	ErrApiSigningInProgress            = NewAPIError("6037", "该用户签约中")
	ErrApiApiSigningNotSupported       = NewAPIError("6038", "该客户不支持API接口签约")
	ErrApiIdCardImagesRequired         = NewAPIError("6039", "服务商需要上传身份证正反面图片")
	ErrApiTaskCodeRequired             = NewAPIError("6040", "服务商需要上传任务编码")
	ErrApiTaskNotFound                 = NewAPIError("6041", "不存在该任务")
	ErrApiRequestTooFrequent           = NewAPIError("6042", "请求频繁请稍后再试")
	ErrApiThreeElementAuthFailed       = NewAPIError("6043", "三要素认证失败")
	ErrApiNotSignedWithProvider        = NewAPIError("6044", "该客户未签约此服务商")
	ErrApiInvoiceCategoryNotFound      = NewAPIError("6045", "未查询到可开票类目信息")
	ErrApiInvoiceInfoNotFound          = NewAPIError("6046", "未查询到该客户在该服务商开票信息")
	ErrApiRiskAuditRequired            = NewAPIError("6047", "该客户订单需要待风控审核后才能下发")
	ErrApiRiskAuditFailed              = NewAPIError("6048", "风控审核未通过")
	ErrApiRecordNotFound               = NewAPIError("6049", "未查询到符合条件的记录")
	ErrApiTaskStatusError              = NewAPIError("6050", "任务状态有误")
	ErrApiOrderNoDuplicate             = NewAPIError("6051", "客户订单号重复,请确认订单信息")
	ErrApiApiNotSupported              = NewAPIError("6052", "该客户不支持 API 接口")
	ErrApiFeeRateNotConfigured         = NewAPIError("6053", "该客户费率未配置")
	ErrApiRechargeOrderNoDuplicate     = NewAPIError("6054", "充值订单号重复，请确认充值信息")
	ErrApiRechargeAmountNotFound       = NewAPIError("6055", "未查询到可充值金额")
	ErrApiRechargeAccountMismatch      = NewAPIError("6056", "充值账号与平台不一致")
	ErrApiRechargeAmountExceeded       = NewAPIError("6057", "充值金额大于可充值金额")
	ErrApiMultiPaymentMethods          = NewAPIError("6058", "批量付款只能选择一种代付方式")
	ErrApiManagementFeeModeMismatch    = NewAPIError("6059", "客户管理费扣费方式与服务商不一致")
	ErrApiManagementFeeRateMismatch    = NewAPIError("6060", "客户管理费费率方式与服务商不一致")
	ErrApiSignConfigNotFound           = NewAPIError("6062", "未查询到签约要素配置")
	ErrApiProviderConfigIncomplete     = NewAPIError("6063", "服务商配置未完成，请联系运营")
	ErrApiEnterpriseApiNotSupported    = NewAPIError("6064", "该企业不支持API,请联系运营")
	ErrApiChannelNotSupported          = NewAPIError("6065", "暂不支持该通道余额查询和分账")
	ErrApiMerchantPublicKeyError       = NewAPIError("6067", "商户公钥格式错误")
	ErrApiOneClickPaymentNotEnabled    = NewAPIError("6093", "未开通一键下发功能，请联系运营")
	ErrApiManualConfirmRequired        = NewAPIError("6100", "个人需手动确认收款，请在app或小程序发起")
	ErrApiVerifySignOrTaskFailed       = NewAPIError("6101", "校验签约，任务领取单等信息失败")
	ErrApiRequestTimeout               = NewAPIError("6102", "请求超时，请重试")
	ErrApiOrderCannotBeCancelled       = NewAPIError("6103", "非待确认订单不可撤销")
	ErrApiSettleTimeError              = NewAPIError("6104", "当前时间不可结算,请稍后重试")
	ErrApiSignTimeError                = NewAPIError("6105", "当前时间不可签约,请稍后重试")
	ErrApiNoElectronicReceipt          = NewAPIError("6220", "暂无电子回单")
)
