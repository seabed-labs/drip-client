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

// Position struct for Position
type Position struct {
	Pubkey string `json:"pubkey"`
	Vault string `json:"vault"`
	Authority string `json:"authority"`
	DepositedTokenAAmount int32 `json:"depositedTokenAAmount"`
	WithdrawnTokenBAmount int32 `json:"withdrawnTokenBAmount"`
	DepositTimestamp string `json:"depositTimestamp"`
	DcaPeriodIdBeforeDeposit int32 `json:"dcaPeriodIdBeforeDeposit"`
	NumberOfSwaps int32 `json:"numberOfSwaps"`
	PeriodicDripAmount int32 `json:"periodicDripAmount"`
	IsClosed bool `json:"isClosed"`
}

// NewPosition instantiates a new Position object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPosition(pubkey string, vault string, authority string, depositedTokenAAmount int32, withdrawnTokenBAmount int32, depositTimestamp string, dcaPeriodIdBeforeDeposit int32, numberOfSwaps int32, periodicDripAmount int32, isClosed bool) *Position {
	this := Position{}
	this.Pubkey = pubkey
	this.Vault = vault
	this.Authority = authority
	this.DepositedTokenAAmount = depositedTokenAAmount
	this.WithdrawnTokenBAmount = withdrawnTokenBAmount
	this.DepositTimestamp = depositTimestamp
	this.DcaPeriodIdBeforeDeposit = dcaPeriodIdBeforeDeposit
	this.NumberOfSwaps = numberOfSwaps
	this.PeriodicDripAmount = periodicDripAmount
	this.IsClosed = isClosed
	return &this
}

// NewPositionWithDefaults instantiates a new Position object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositionWithDefaults() *Position {
	this := Position{}
	return &this
}

// GetPubkey returns the Pubkey field value
func (o *Position) GetPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value
// and a boolean to check if the value has been set.
func (o *Position) GetPubkeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pubkey, true
}

// SetPubkey sets field value
func (o *Position) SetPubkey(v string) {
	o.Pubkey = v
}

// GetVault returns the Vault field value
func (o *Position) GetVault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vault
}

// GetVaultOk returns a tuple with the Vault field value
// and a boolean to check if the value has been set.
func (o *Position) GetVaultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vault, true
}

// SetVault sets field value
func (o *Position) SetVault(v string) {
	o.Vault = v
}

// GetAuthority returns the Authority field value
func (o *Position) GetAuthority() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Authority
}

// GetAuthorityOk returns a tuple with the Authority field value
// and a boolean to check if the value has been set.
func (o *Position) GetAuthorityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Authority, true
}

// SetAuthority sets field value
func (o *Position) SetAuthority(v string) {
	o.Authority = v
}

// GetDepositedTokenAAmount returns the DepositedTokenAAmount field value
func (o *Position) GetDepositedTokenAAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DepositedTokenAAmount
}

// GetDepositedTokenAAmountOk returns a tuple with the DepositedTokenAAmount field value
// and a boolean to check if the value has been set.
func (o *Position) GetDepositedTokenAAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositedTokenAAmount, true
}

// SetDepositedTokenAAmount sets field value
func (o *Position) SetDepositedTokenAAmount(v int32) {
	o.DepositedTokenAAmount = v
}

// GetWithdrawnTokenBAmount returns the WithdrawnTokenBAmount field value
func (o *Position) GetWithdrawnTokenBAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WithdrawnTokenBAmount
}

// GetWithdrawnTokenBAmountOk returns a tuple with the WithdrawnTokenBAmount field value
// and a boolean to check if the value has been set.
func (o *Position) GetWithdrawnTokenBAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WithdrawnTokenBAmount, true
}

// SetWithdrawnTokenBAmount sets field value
func (o *Position) SetWithdrawnTokenBAmount(v int32) {
	o.WithdrawnTokenBAmount = v
}

// GetDepositTimestamp returns the DepositTimestamp field value
func (o *Position) GetDepositTimestamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepositTimestamp
}

// GetDepositTimestampOk returns a tuple with the DepositTimestamp field value
// and a boolean to check if the value has been set.
func (o *Position) GetDepositTimestampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepositTimestamp, true
}

// SetDepositTimestamp sets field value
func (o *Position) SetDepositTimestamp(v string) {
	o.DepositTimestamp = v
}

// GetDcaPeriodIdBeforeDeposit returns the DcaPeriodIdBeforeDeposit field value
func (o *Position) GetDcaPeriodIdBeforeDeposit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DcaPeriodIdBeforeDeposit
}

// GetDcaPeriodIdBeforeDepositOk returns a tuple with the DcaPeriodIdBeforeDeposit field value
// and a boolean to check if the value has been set.
func (o *Position) GetDcaPeriodIdBeforeDepositOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DcaPeriodIdBeforeDeposit, true
}

// SetDcaPeriodIdBeforeDeposit sets field value
func (o *Position) SetDcaPeriodIdBeforeDeposit(v int32) {
	o.DcaPeriodIdBeforeDeposit = v
}

// GetNumberOfSwaps returns the NumberOfSwaps field value
func (o *Position) GetNumberOfSwaps() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfSwaps
}

// GetNumberOfSwapsOk returns a tuple with the NumberOfSwaps field value
// and a boolean to check if the value has been set.
func (o *Position) GetNumberOfSwapsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfSwaps, true
}

// SetNumberOfSwaps sets field value
func (o *Position) SetNumberOfSwaps(v int32) {
	o.NumberOfSwaps = v
}

// GetPeriodicDripAmount returns the PeriodicDripAmount field value
func (o *Position) GetPeriodicDripAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PeriodicDripAmount
}

// GetPeriodicDripAmountOk returns a tuple with the PeriodicDripAmount field value
// and a boolean to check if the value has been set.
func (o *Position) GetPeriodicDripAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PeriodicDripAmount, true
}

// SetPeriodicDripAmount sets field value
func (o *Position) SetPeriodicDripAmount(v int32) {
	o.PeriodicDripAmount = v
}

// GetIsClosed returns the IsClosed field value
func (o *Position) GetIsClosed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsClosed
}

// GetIsClosedOk returns a tuple with the IsClosed field value
// and a boolean to check if the value has been set.
func (o *Position) GetIsClosedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsClosed, true
}

// SetIsClosed sets field value
func (o *Position) SetIsClosed(v bool) {
	o.IsClosed = v
}

func (o Position) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pubkey"] = o.Pubkey
	}
	if true {
		toSerialize["vault"] = o.Vault
	}
	if true {
		toSerialize["authority"] = o.Authority
	}
	if true {
		toSerialize["depositedTokenAAmount"] = o.DepositedTokenAAmount
	}
	if true {
		toSerialize["withdrawnTokenBAmount"] = o.WithdrawnTokenBAmount
	}
	if true {
		toSerialize["depositTimestamp"] = o.DepositTimestamp
	}
	if true {
		toSerialize["dcaPeriodIdBeforeDeposit"] = o.DcaPeriodIdBeforeDeposit
	}
	if true {
		toSerialize["numberOfSwaps"] = o.NumberOfSwaps
	}
	if true {
		toSerialize["periodicDripAmount"] = o.PeriodicDripAmount
	}
	if true {
		toSerialize["isClosed"] = o.IsClosed
	}
	return json.Marshal(toSerialize)
}

type NullablePosition struct {
	value *Position
	isSet bool
}

func (v NullablePosition) Get() *Position {
	return v.value
}

func (v *NullablePosition) Set(val *Position) {
	v.value = val
	v.isSet = true
}

func (v NullablePosition) IsSet() bool {
	return v.isSet
}

func (v *NullablePosition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePosition(val *Position) *NullablePosition {
	return &NullablePosition{value: val, isSet: true}
}

func (v NullablePosition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePosition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


