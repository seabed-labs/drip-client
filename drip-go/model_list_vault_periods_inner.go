/*
Drip Backend

Drip backend service.

API version: 1.0.0
Contact: dcafmocha@protonmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package drip

import (
	"encoding/json"
)

// ListVaultPeriodsInner struct for ListVaultPeriodsInner
type ListVaultPeriodsInner struct {
	Pubkey string `json:"pubkey"`
	Vault string `json:"vault"`
	PeriodId string `json:"periodId"`
	Twap string `json:"twap"`
	Dar string `json:"dar"`
}

// NewListVaultPeriodsInner instantiates a new ListVaultPeriodsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVaultPeriodsInner(pubkey string, vault string, periodId string, twap string, dar string) *ListVaultPeriodsInner {
	this := ListVaultPeriodsInner{}
	this.Pubkey = pubkey
	this.Vault = vault
	this.PeriodId = periodId
	this.Twap = twap
	this.Dar = dar
	return &this
}

// NewListVaultPeriodsInnerWithDefaults instantiates a new ListVaultPeriodsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVaultPeriodsInnerWithDefaults() *ListVaultPeriodsInner {
	this := ListVaultPeriodsInner{}
	return &this
}

// GetPubkey returns the Pubkey field value
func (o *ListVaultPeriodsInner) GetPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value
// and a boolean to check if the value has been set.
func (o *ListVaultPeriodsInner) GetPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pubkey, true
}

// SetPubkey sets field value
func (o *ListVaultPeriodsInner) SetPubkey(v string) {
	o.Pubkey = v
}

// GetVault returns the Vault field value
func (o *ListVaultPeriodsInner) GetVault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vault
}

// GetVaultOk returns a tuple with the Vault field value
// and a boolean to check if the value has been set.
func (o *ListVaultPeriodsInner) GetVaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vault, true
}

// SetVault sets field value
func (o *ListVaultPeriodsInner) SetVault(v string) {
	o.Vault = v
}

// GetPeriodId returns the PeriodId field value
func (o *ListVaultPeriodsInner) GetPeriodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodId
}

// GetPeriodIdOk returns a tuple with the PeriodId field value
// and a boolean to check if the value has been set.
func (o *ListVaultPeriodsInner) GetPeriodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodId, true
}

// SetPeriodId sets field value
func (o *ListVaultPeriodsInner) SetPeriodId(v string) {
	o.PeriodId = v
}

// GetTwap returns the Twap field value
func (o *ListVaultPeriodsInner) GetTwap() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Twap
}

// GetTwapOk returns a tuple with the Twap field value
// and a boolean to check if the value has been set.
func (o *ListVaultPeriodsInner) GetTwapOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Twap, true
}

// SetTwap sets field value
func (o *ListVaultPeriodsInner) SetTwap(v string) {
	o.Twap = v
}

// GetDar returns the Dar field value
func (o *ListVaultPeriodsInner) GetDar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dar
}

// GetDarOk returns a tuple with the Dar field value
// and a boolean to check if the value has been set.
func (o *ListVaultPeriodsInner) GetDarOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dar, true
}

// SetDar sets field value
func (o *ListVaultPeriodsInner) SetDar(v string) {
	o.Dar = v
}

func (o ListVaultPeriodsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pubkey"] = o.Pubkey
	}
	if true {
		toSerialize["vault"] = o.Vault
	}
	if true {
		toSerialize["periodId"] = o.PeriodId
	}
	if true {
		toSerialize["twap"] = o.Twap
	}
	if true {
		toSerialize["dar"] = o.Dar
	}
	return json.Marshal(toSerialize)
}

type NullableListVaultPeriodsInner struct {
	value *ListVaultPeriodsInner
	isSet bool
}

func (v NullableListVaultPeriodsInner) Get() *ListVaultPeriodsInner {
	return v.value
}

func (v *NullableListVaultPeriodsInner) Set(val *ListVaultPeriodsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListVaultPeriodsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListVaultPeriodsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVaultPeriodsInner(val *ListVaultPeriodsInner) *NullableListVaultPeriodsInner {
	return &NullableListVaultPeriodsInner{value: val, isSet: true}
}

func (v NullableListVaultPeriodsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVaultPeriodsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

