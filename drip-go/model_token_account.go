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

// TokenAccount struct for TokenAccount
type TokenAccount struct {
	Pubkey string `json:"pubkey"`
	Mint string `json:"mint"`
	Owner string `json:"owner"`
	Amount string `json:"amount"`
	State string `json:"state"`
}

// NewTokenAccount instantiates a new TokenAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTokenAccount(pubkey string, mint string, owner string, amount string, state string) *TokenAccount {
	this := TokenAccount{}
	this.Pubkey = pubkey
	this.Mint = mint
	this.Owner = owner
	this.Amount = amount
	this.State = state
	return &this
}

// NewTokenAccountWithDefaults instantiates a new TokenAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTokenAccountWithDefaults() *TokenAccount {
	this := TokenAccount{}
	return &this
}

// GetPubkey returns the Pubkey field value
func (o *TokenAccount) GetPubkey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pubkey
}

// GetPubkeyOk returns a tuple with the Pubkey field value
// and a boolean to check if the value has been set.
func (o *TokenAccount) GetPubkeyOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Pubkey, true
}

// SetPubkey sets field value
func (o *TokenAccount) SetPubkey(v string) {
	o.Pubkey = v
}

// GetMint returns the Mint field value
func (o *TokenAccount) GetMint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mint
}

// GetMintOk returns a tuple with the Mint field value
// and a boolean to check if the value has been set.
func (o *TokenAccount) GetMintOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Mint, true
}

// SetMint sets field value
func (o *TokenAccount) SetMint(v string) {
	o.Mint = v
}

// GetOwner returns the Owner field value
func (o *TokenAccount) GetOwner() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value
// and a boolean to check if the value has been set.
func (o *TokenAccount) GetOwnerOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Owner, true
}

// SetOwner sets field value
func (o *TokenAccount) SetOwner(v string) {
	o.Owner = v
}

// GetAmount returns the Amount field value
func (o *TokenAccount) GetAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *TokenAccount) GetAmountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *TokenAccount) SetAmount(v string) {
	o.Amount = v
}

// GetState returns the State field value
func (o *TokenAccount) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *TokenAccount) GetStateOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *TokenAccount) SetState(v string) {
	o.State = v
}

func (o TokenAccount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pubkey"] = o.Pubkey
	}
	if true {
		toSerialize["mint"] = o.Mint
	}
	if true {
		toSerialize["owner"] = o.Owner
	}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableTokenAccount struct {
	value *TokenAccount
	isSet bool
}

func (v NullableTokenAccount) Get() *TokenAccount {
	return v.value
}

func (v *NullableTokenAccount) Set(val *TokenAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableTokenAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableTokenAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTokenAccount(val *TokenAccount) *NullableTokenAccount {
	return &NullableTokenAccount{value: val, isSet: true}
}

func (v NullableTokenAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTokenAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


