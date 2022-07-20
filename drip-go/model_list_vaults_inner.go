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

// ListVaultsInner struct for ListVaultsInner
type ListVaultsInner struct {
	Pubkey string `json:"pubkey"`
	ProtoConfig string `json:"protoConfig"`
	TokenAAccount string `json:"tokenAAccount"`
	TokenBAccount string `json:"tokenBAccount"`
	TreasuryTokenBAccount string `json:"treasuryTokenBAccount"`
	TokenAMint string `json:"tokenAMint"`
	TokenBMint string `json:"tokenBMint"`
	LastDcaPeriod string `json:"lastDcaPeriod"`
	DripAmount string `json:"dripAmount"`
	// unix timestamp
	DcaActivationTimestamp string `json:"dcaActivationTimestamp"`
}

// NewListVaultsInner instantiates a new ListVaultsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVaultsInner(pubkey string, protoConfig string, tokenAAccount string, tokenBAccount string, treasuryTokenBAccount string, tokenAMint string, tokenBMint string, lastDcaPeriod string, dripAmount string, dcaActivationTimestamp string) *ListVaultsInner {
	this := ListVaultsInner{}
	this.Pubkey = pubkey
	this.ProtoConfig = protoConfig
	this.TokenAAccount = tokenAAccount
	this.TokenBAccount = tokenBAccount
	this.TreasuryTokenBAccount = treasuryTokenBAccount
	this.TokenAMint = tokenAMint
	this.TokenBMint = tokenBMint
	this.LastDcaPeriod = lastDcaPeriod
	this.DripAmount = dripAmount
	this.DcaActivationTimestamp = dcaActivationTimestamp
	return &this
}

// NewListVaultsInnerWithDefaults instantiates a new ListVaultsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVaultsInnerWithDefaults() *ListVaultsInner {
	this := ListVaultsInner{}
	return &this
}

// GetPubkey returns the Pubkey field value
func (o *ListVaultsInner) GetPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pubkey, true
}

// SetPubkey sets field value
func (o *ListVaultsInner) SetPubkey(v string) {
	o.Pubkey = v
}

// GetProtoConfig returns the ProtoConfig field value
func (o *ListVaultsInner) GetProtoConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProtoConfig
}

// GetProtoConfigOk returns a tuple with the ProtoConfig field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetProtoConfigOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProtoConfig, true
}

// SetProtoConfig sets field value
func (o *ListVaultsInner) SetProtoConfig(v string) {
	o.ProtoConfig = v
}

// GetTokenAAccount returns the TokenAAccount field value
func (o *ListVaultsInner) GetTokenAAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAAccount
}

// GetTokenAAccountOk returns a tuple with the TokenAAccount field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetTokenAAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAAccount, true
}

// SetTokenAAccount sets field value
func (o *ListVaultsInner) SetTokenAAccount(v string) {
	o.TokenAAccount = v
}

// GetTokenBAccount returns the TokenBAccount field value
func (o *ListVaultsInner) GetTokenBAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenBAccount
}

// GetTokenBAccountOk returns a tuple with the TokenBAccount field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetTokenBAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenBAccount, true
}

// SetTokenBAccount sets field value
func (o *ListVaultsInner) SetTokenBAccount(v string) {
	o.TokenBAccount = v
}

// GetTreasuryTokenBAccount returns the TreasuryTokenBAccount field value
func (o *ListVaultsInner) GetTreasuryTokenBAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TreasuryTokenBAccount
}

// GetTreasuryTokenBAccountOk returns a tuple with the TreasuryTokenBAccount field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetTreasuryTokenBAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TreasuryTokenBAccount, true
}

// SetTreasuryTokenBAccount sets field value
func (o *ListVaultsInner) SetTreasuryTokenBAccount(v string) {
	o.TreasuryTokenBAccount = v
}

// GetTokenAMint returns the TokenAMint field value
func (o *ListVaultsInner) GetTokenAMint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAMint
}

// GetTokenAMintOk returns a tuple with the TokenAMint field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetTokenAMintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenAMint, true
}

// SetTokenAMint sets field value
func (o *ListVaultsInner) SetTokenAMint(v string) {
	o.TokenAMint = v
}

// GetTokenBMint returns the TokenBMint field value
func (o *ListVaultsInner) GetTokenBMint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenBMint
}

// GetTokenBMintOk returns a tuple with the TokenBMint field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetTokenBMintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenBMint, true
}

// SetTokenBMint sets field value
func (o *ListVaultsInner) SetTokenBMint(v string) {
	o.TokenBMint = v
}

// GetLastDcaPeriod returns the LastDcaPeriod field value
func (o *ListVaultsInner) GetLastDcaPeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastDcaPeriod
}

// GetLastDcaPeriodOk returns a tuple with the LastDcaPeriod field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetLastDcaPeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastDcaPeriod, true
}

// SetLastDcaPeriod sets field value
func (o *ListVaultsInner) SetLastDcaPeriod(v string) {
	o.LastDcaPeriod = v
}

// GetDripAmount returns the DripAmount field value
func (o *ListVaultsInner) GetDripAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DripAmount
}

// GetDripAmountOk returns a tuple with the DripAmount field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetDripAmountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DripAmount, true
}

// SetDripAmount sets field value
func (o *ListVaultsInner) SetDripAmount(v string) {
	o.DripAmount = v
}

// GetDcaActivationTimestamp returns the DcaActivationTimestamp field value
func (o *ListVaultsInner) GetDcaActivationTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DcaActivationTimestamp
}

// GetDcaActivationTimestampOk returns a tuple with the DcaActivationTimestamp field value
// and a boolean to check if the value has been set.
func (o *ListVaultsInner) GetDcaActivationTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcaActivationTimestamp, true
}

// SetDcaActivationTimestamp sets field value
func (o *ListVaultsInner) SetDcaActivationTimestamp(v string) {
	o.DcaActivationTimestamp = v
}

func (o ListVaultsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pubkey"] = o.Pubkey
	}
	if true {
		toSerialize["protoConfig"] = o.ProtoConfig
	}
	if true {
		toSerialize["tokenAAccount"] = o.TokenAAccount
	}
	if true {
		toSerialize["tokenBAccount"] = o.TokenBAccount
	}
	if true {
		toSerialize["treasuryTokenBAccount"] = o.TreasuryTokenBAccount
	}
	if true {
		toSerialize["tokenAMint"] = o.TokenAMint
	}
	if true {
		toSerialize["tokenBMint"] = o.TokenBMint
	}
	if true {
		toSerialize["lastDcaPeriod"] = o.LastDcaPeriod
	}
	if true {
		toSerialize["dripAmount"] = o.DripAmount
	}
	if true {
		toSerialize["dcaActivationTimestamp"] = o.DcaActivationTimestamp
	}
	return json.Marshal(toSerialize)
}

type NullableListVaultsInner struct {
	value *ListVaultsInner
	isSet bool
}

func (v NullableListVaultsInner) Get() *ListVaultsInner {
	return v.value
}

func (v *NullableListVaultsInner) Set(val *ListVaultsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListVaultsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListVaultsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVaultsInner(val *ListVaultsInner) *NullableListVaultsInner {
	return &NullableListVaultsInner{value: val, isSet: true}
}

func (v NullableListVaultsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVaultsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


