/*
Drip Backend

Drip backend service. All API's have a IP rate limit of 10 requests per second. 

API version: 1.0.0
Contact: mocha@dcaf.so
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package drip

import (
	"encoding/json"
)

// ExpandedAdminVaultAllOf struct for ExpandedAdminVaultAllOf
type ExpandedAdminVaultAllOf struct {
	ProtoConfigValue *ProtoConfig `json:"protoConfigValue,omitempty"`
	TokenAMintValue *Token `json:"tokenAMintValue,omitempty"`
	TokenBMintValue *Token `json:"tokenBMintValue,omitempty"`
	TokenAAccountValue *TokenAccount `json:"tokenAAccountValue,omitempty"`
	TokenBAccountValue *TokenAccount `json:"tokenBAccountValue,omitempty"`
	TreasuryTokenBAccountValue *TokenAccount `json:"treasuryTokenBAccountValue,omitempty"`
}

// NewExpandedAdminVaultAllOf instantiates a new ExpandedAdminVaultAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExpandedAdminVaultAllOf() *ExpandedAdminVaultAllOf {
	this := ExpandedAdminVaultAllOf{}
	return &this
}

// NewExpandedAdminVaultAllOfWithDefaults instantiates a new ExpandedAdminVaultAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExpandedAdminVaultAllOfWithDefaults() *ExpandedAdminVaultAllOf {
	this := ExpandedAdminVaultAllOf{}
	return &this
}

// GetProtoConfigValue returns the ProtoConfigValue field value if set, zero value otherwise.
func (o *ExpandedAdminVaultAllOf) GetProtoConfigValue() ProtoConfig {
	if o == nil || isNil(o.ProtoConfigValue) {
		var ret ProtoConfig
		return ret
	}
	return *o.ProtoConfigValue
}

// GetProtoConfigValueOk returns a tuple with the ProtoConfigValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminVaultAllOf) GetProtoConfigValueOk() (*ProtoConfig, bool) {
	if o == nil || isNil(o.ProtoConfigValue) {
    return nil, false
	}
	return o.ProtoConfigValue, true
}

// HasProtoConfigValue returns a boolean if a field has been set.
func (o *ExpandedAdminVaultAllOf) HasProtoConfigValue() bool {
	if o != nil && !isNil(o.ProtoConfigValue) {
		return true
	}

	return false
}

// SetProtoConfigValue gets a reference to the given ProtoConfig and assigns it to the ProtoConfigValue field.
func (o *ExpandedAdminVaultAllOf) SetProtoConfigValue(v ProtoConfig) {
	o.ProtoConfigValue = &v
}

// GetTokenAMintValue returns the TokenAMintValue field value if set, zero value otherwise.
func (o *ExpandedAdminVaultAllOf) GetTokenAMintValue() Token {
	if o == nil || isNil(o.TokenAMintValue) {
		var ret Token
		return ret
	}
	return *o.TokenAMintValue
}

// GetTokenAMintValueOk returns a tuple with the TokenAMintValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminVaultAllOf) GetTokenAMintValueOk() (*Token, bool) {
	if o == nil || isNil(o.TokenAMintValue) {
    return nil, false
	}
	return o.TokenAMintValue, true
}

// HasTokenAMintValue returns a boolean if a field has been set.
func (o *ExpandedAdminVaultAllOf) HasTokenAMintValue() bool {
	if o != nil && !isNil(o.TokenAMintValue) {
		return true
	}

	return false
}

// SetTokenAMintValue gets a reference to the given Token and assigns it to the TokenAMintValue field.
func (o *ExpandedAdminVaultAllOf) SetTokenAMintValue(v Token) {
	o.TokenAMintValue = &v
}

// GetTokenBMintValue returns the TokenBMintValue field value if set, zero value otherwise.
func (o *ExpandedAdminVaultAllOf) GetTokenBMintValue() Token {
	if o == nil || isNil(o.TokenBMintValue) {
		var ret Token
		return ret
	}
	return *o.TokenBMintValue
}

// GetTokenBMintValueOk returns a tuple with the TokenBMintValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminVaultAllOf) GetTokenBMintValueOk() (*Token, bool) {
	if o == nil || isNil(o.TokenBMintValue) {
    return nil, false
	}
	return o.TokenBMintValue, true
}

// HasTokenBMintValue returns a boolean if a field has been set.
func (o *ExpandedAdminVaultAllOf) HasTokenBMintValue() bool {
	if o != nil && !isNil(o.TokenBMintValue) {
		return true
	}

	return false
}

// SetTokenBMintValue gets a reference to the given Token and assigns it to the TokenBMintValue field.
func (o *ExpandedAdminVaultAllOf) SetTokenBMintValue(v Token) {
	o.TokenBMintValue = &v
}

// GetTokenAAccountValue returns the TokenAAccountValue field value if set, zero value otherwise.
func (o *ExpandedAdminVaultAllOf) GetTokenAAccountValue() TokenAccount {
	if o == nil || isNil(o.TokenAAccountValue) {
		var ret TokenAccount
		return ret
	}
	return *o.TokenAAccountValue
}

// GetTokenAAccountValueOk returns a tuple with the TokenAAccountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminVaultAllOf) GetTokenAAccountValueOk() (*TokenAccount, bool) {
	if o == nil || isNil(o.TokenAAccountValue) {
    return nil, false
	}
	return o.TokenAAccountValue, true
}

// HasTokenAAccountValue returns a boolean if a field has been set.
func (o *ExpandedAdminVaultAllOf) HasTokenAAccountValue() bool {
	if o != nil && !isNil(o.TokenAAccountValue) {
		return true
	}

	return false
}

// SetTokenAAccountValue gets a reference to the given TokenAccount and assigns it to the TokenAAccountValue field.
func (o *ExpandedAdminVaultAllOf) SetTokenAAccountValue(v TokenAccount) {
	o.TokenAAccountValue = &v
}

// GetTokenBAccountValue returns the TokenBAccountValue field value if set, zero value otherwise.
func (o *ExpandedAdminVaultAllOf) GetTokenBAccountValue() TokenAccount {
	if o == nil || isNil(o.TokenBAccountValue) {
		var ret TokenAccount
		return ret
	}
	return *o.TokenBAccountValue
}

// GetTokenBAccountValueOk returns a tuple with the TokenBAccountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminVaultAllOf) GetTokenBAccountValueOk() (*TokenAccount, bool) {
	if o == nil || isNil(o.TokenBAccountValue) {
    return nil, false
	}
	return o.TokenBAccountValue, true
}

// HasTokenBAccountValue returns a boolean if a field has been set.
func (o *ExpandedAdminVaultAllOf) HasTokenBAccountValue() bool {
	if o != nil && !isNil(o.TokenBAccountValue) {
		return true
	}

	return false
}

// SetTokenBAccountValue gets a reference to the given TokenAccount and assigns it to the TokenBAccountValue field.
func (o *ExpandedAdminVaultAllOf) SetTokenBAccountValue(v TokenAccount) {
	o.TokenBAccountValue = &v
}

// GetTreasuryTokenBAccountValue returns the TreasuryTokenBAccountValue field value if set, zero value otherwise.
func (o *ExpandedAdminVaultAllOf) GetTreasuryTokenBAccountValue() TokenAccount {
	if o == nil || isNil(o.TreasuryTokenBAccountValue) {
		var ret TokenAccount
		return ret
	}
	return *o.TreasuryTokenBAccountValue
}

// GetTreasuryTokenBAccountValueOk returns a tuple with the TreasuryTokenBAccountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExpandedAdminVaultAllOf) GetTreasuryTokenBAccountValueOk() (*TokenAccount, bool) {
	if o == nil || isNil(o.TreasuryTokenBAccountValue) {
    return nil, false
	}
	return o.TreasuryTokenBAccountValue, true
}

// HasTreasuryTokenBAccountValue returns a boolean if a field has been set.
func (o *ExpandedAdminVaultAllOf) HasTreasuryTokenBAccountValue() bool {
	if o != nil && !isNil(o.TreasuryTokenBAccountValue) {
		return true
	}

	return false
}

// SetTreasuryTokenBAccountValue gets a reference to the given TokenAccount and assigns it to the TreasuryTokenBAccountValue field.
func (o *ExpandedAdminVaultAllOf) SetTreasuryTokenBAccountValue(v TokenAccount) {
	o.TreasuryTokenBAccountValue = &v
}

func (o ExpandedAdminVaultAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ProtoConfigValue) {
		toSerialize["protoConfigValue"] = o.ProtoConfigValue
	}
	if !isNil(o.TokenAMintValue) {
		toSerialize["tokenAMintValue"] = o.TokenAMintValue
	}
	if !isNil(o.TokenBMintValue) {
		toSerialize["tokenBMintValue"] = o.TokenBMintValue
	}
	if !isNil(o.TokenAAccountValue) {
		toSerialize["tokenAAccountValue"] = o.TokenAAccountValue
	}
	if !isNil(o.TokenBAccountValue) {
		toSerialize["tokenBAccountValue"] = o.TokenBAccountValue
	}
	if !isNil(o.TreasuryTokenBAccountValue) {
		toSerialize["treasuryTokenBAccountValue"] = o.TreasuryTokenBAccountValue
	}
	return json.Marshal(toSerialize)
}

type NullableExpandedAdminVaultAllOf struct {
	value *ExpandedAdminVaultAllOf
	isSet bool
}

func (v NullableExpandedAdminVaultAllOf) Get() *ExpandedAdminVaultAllOf {
	return v.value
}

func (v *NullableExpandedAdminVaultAllOf) Set(val *ExpandedAdminVaultAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExpandedAdminVaultAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExpandedAdminVaultAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpandedAdminVaultAllOf(val *ExpandedAdminVaultAllOf) *NullableExpandedAdminVaultAllOf {
	return &NullableExpandedAdminVaultAllOf{value: val, isSet: true}
}

func (v NullableExpandedAdminVaultAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpandedAdminVaultAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


