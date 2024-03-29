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

// VaultPeriod struct for VaultPeriod
type VaultPeriod struct {
	Pubkey string `json:"pubkey"`
	Vault string `json:"vault"`
	PeriodId string `json:"periodId"`
	Twap string `json:"twap"`
	Dar string `json:"dar"`
	PriceBOverA *string `json:"priceBOverA,omitempty"`
}

// NewVaultPeriod instantiates a new VaultPeriod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVaultPeriod(pubkey string, vault string, periodId string, twap string, dar string) *VaultPeriod {
	this := VaultPeriod{}
	this.Pubkey = pubkey
	this.Vault = vault
	this.PeriodId = periodId
	this.Twap = twap
	this.Dar = dar
	return &this
}

// NewVaultPeriodWithDefaults instantiates a new VaultPeriod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVaultPeriodWithDefaults() *VaultPeriod {
	this := VaultPeriod{}
	return &this
}

// GetPubkey returns the Pubkey field value
func (o *VaultPeriod) GetPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value
// and a boolean to check if the value has been set.
func (o *VaultPeriod) GetPubkeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pubkey, true
}

// SetPubkey sets field value
func (o *VaultPeriod) SetPubkey(v string) {
	o.Pubkey = v
}

// GetVault returns the Vault field value
func (o *VaultPeriod) GetVault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vault
}

// GetVaultOk returns a tuple with the Vault field value
// and a boolean to check if the value has been set.
func (o *VaultPeriod) GetVaultOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vault, true
}

// SetVault sets field value
func (o *VaultPeriod) SetVault(v string) {
	o.Vault = v
}

// GetPeriodId returns the PeriodId field value
func (o *VaultPeriod) GetPeriodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PeriodId
}

// GetPeriodIdOk returns a tuple with the PeriodId field value
// and a boolean to check if the value has been set.
func (o *VaultPeriod) GetPeriodIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.PeriodId, true
}

// SetPeriodId sets field value
func (o *VaultPeriod) SetPeriodId(v string) {
	o.PeriodId = v
}

// GetTwap returns the Twap field value
func (o *VaultPeriod) GetTwap() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Twap
}

// GetTwapOk returns a tuple with the Twap field value
// and a boolean to check if the value has been set.
func (o *VaultPeriod) GetTwapOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Twap, true
}

// SetTwap sets field value
func (o *VaultPeriod) SetTwap(v string) {
	o.Twap = v
}

// GetDar returns the Dar field value
func (o *VaultPeriod) GetDar() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dar
}

// GetDarOk returns a tuple with the Dar field value
// and a boolean to check if the value has been set.
func (o *VaultPeriod) GetDarOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Dar, true
}

// SetDar sets field value
func (o *VaultPeriod) SetDar(v string) {
	o.Dar = v
}

// GetPriceBOverA returns the PriceBOverA field value if set, zero value otherwise.
func (o *VaultPeriod) GetPriceBOverA() string {
	if o == nil || isNil(o.PriceBOverA) {
		var ret string
		return ret
	}
	return *o.PriceBOverA
}

// GetPriceBOverAOk returns a tuple with the PriceBOverA field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VaultPeriod) GetPriceBOverAOk() (*string, bool) {
	if o == nil || isNil(o.PriceBOverA) {
    return nil, false
	}
	return o.PriceBOverA, true
}

// HasPriceBOverA returns a boolean if a field has been set.
func (o *VaultPeriod) HasPriceBOverA() bool {
	if o != nil && !isNil(o.PriceBOverA) {
		return true
	}

	return false
}

// SetPriceBOverA gets a reference to the given string and assigns it to the PriceBOverA field.
func (o *VaultPeriod) SetPriceBOverA(v string) {
	o.PriceBOverA = &v
}

func (o VaultPeriod) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.PriceBOverA) {
		toSerialize["priceBOverA"] = o.PriceBOverA
	}
	return json.Marshal(toSerialize)
}

type NullableVaultPeriod struct {
	value *VaultPeriod
	isSet bool
}

func (v NullableVaultPeriod) Get() *VaultPeriod {
	return v.value
}

func (v *NullableVaultPeriod) Set(val *VaultPeriod) {
	v.value = val
	v.isSet = true
}

func (v NullableVaultPeriod) IsSet() bool {
	return v.isSet
}

func (v *NullableVaultPeriod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVaultPeriod(val *VaultPeriod) *NullableVaultPeriod {
	return &NullableVaultPeriod{value: val, isSet: true}
}

func (v NullableVaultPeriod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVaultPeriod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


