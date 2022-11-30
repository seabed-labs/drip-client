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

// OrcaWhirlpoolConfig struct for OrcaWhirlpoolConfig
type OrcaWhirlpoolConfig struct {
	Vault string `json:"vault"`
	VaultProtoConfig string `json:"vaultProtoConfig"`
	VaultTokenAAccount string `json:"vaultTokenAAccount"`
	VaultTokenBAccount string `json:"vaultTokenBAccount"`
	TokenAMint string `json:"tokenAMint"`
	TokenBMint string `json:"tokenBMint"`
	Whirlpool string `json:"whirlpool"`
	TokenVaultA string `json:"tokenVaultA"`
	TokenVaultB string `json:"tokenVaultB"`
	Oracle string `json:"oracle"`
}

// NewOrcaWhirlpoolConfig instantiates a new OrcaWhirlpoolConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrcaWhirlpoolConfig(vault string, vaultProtoConfig string, vaultTokenAAccount string, vaultTokenBAccount string, tokenAMint string, tokenBMint string, whirlpool string, tokenVaultA string, tokenVaultB string, oracle string) *OrcaWhirlpoolConfig {
	this := OrcaWhirlpoolConfig{}
	this.Vault = vault
	this.VaultProtoConfig = vaultProtoConfig
	this.VaultTokenAAccount = vaultTokenAAccount
	this.VaultTokenBAccount = vaultTokenBAccount
	this.TokenAMint = tokenAMint
	this.TokenBMint = tokenBMint
	this.Whirlpool = whirlpool
	this.TokenVaultA = tokenVaultA
	this.TokenVaultB = tokenVaultB
	this.Oracle = oracle
	return &this
}

// NewOrcaWhirlpoolConfigWithDefaults instantiates a new OrcaWhirlpoolConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrcaWhirlpoolConfigWithDefaults() *OrcaWhirlpoolConfig {
	this := OrcaWhirlpoolConfig{}
	return &this
}

// GetVault returns the Vault field value
func (o *OrcaWhirlpoolConfig) GetVault() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vault
}

// GetVaultOk returns a tuple with the Vault field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetVaultOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Vault, true
}

// SetVault sets field value
func (o *OrcaWhirlpoolConfig) SetVault(v string) {
	o.Vault = v
}

// GetVaultProtoConfig returns the VaultProtoConfig field value
func (o *OrcaWhirlpoolConfig) GetVaultProtoConfig() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultProtoConfig
}

// GetVaultProtoConfigOk returns a tuple with the VaultProtoConfig field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetVaultProtoConfigOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultProtoConfig, true
}

// SetVaultProtoConfig sets field value
func (o *OrcaWhirlpoolConfig) SetVaultProtoConfig(v string) {
	o.VaultProtoConfig = v
}

// GetVaultTokenAAccount returns the VaultTokenAAccount field value
func (o *OrcaWhirlpoolConfig) GetVaultTokenAAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultTokenAAccount
}

// GetVaultTokenAAccountOk returns a tuple with the VaultTokenAAccount field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetVaultTokenAAccountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultTokenAAccount, true
}

// SetVaultTokenAAccount sets field value
func (o *OrcaWhirlpoolConfig) SetVaultTokenAAccount(v string) {
	o.VaultTokenAAccount = v
}

// GetVaultTokenBAccount returns the VaultTokenBAccount field value
func (o *OrcaWhirlpoolConfig) GetVaultTokenBAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VaultTokenBAccount
}

// GetVaultTokenBAccountOk returns a tuple with the VaultTokenBAccount field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetVaultTokenBAccountOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VaultTokenBAccount, true
}

// SetVaultTokenBAccount sets field value
func (o *OrcaWhirlpoolConfig) SetVaultTokenBAccount(v string) {
	o.VaultTokenBAccount = v
}

// GetTokenAMint returns the TokenAMint field value
func (o *OrcaWhirlpoolConfig) GetTokenAMint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenAMint
}

// GetTokenAMintOk returns a tuple with the TokenAMint field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetTokenAMintOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TokenAMint, true
}

// SetTokenAMint sets field value
func (o *OrcaWhirlpoolConfig) SetTokenAMint(v string) {
	o.TokenAMint = v
}

// GetTokenBMint returns the TokenBMint field value
func (o *OrcaWhirlpoolConfig) GetTokenBMint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenBMint
}

// GetTokenBMintOk returns a tuple with the TokenBMint field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetTokenBMintOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TokenBMint, true
}

// SetTokenBMint sets field value
func (o *OrcaWhirlpoolConfig) SetTokenBMint(v string) {
	o.TokenBMint = v
}

// GetWhirlpool returns the Whirlpool field value
func (o *OrcaWhirlpoolConfig) GetWhirlpool() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Whirlpool
}

// GetWhirlpoolOk returns a tuple with the Whirlpool field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetWhirlpoolOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Whirlpool, true
}

// SetWhirlpool sets field value
func (o *OrcaWhirlpoolConfig) SetWhirlpool(v string) {
	o.Whirlpool = v
}

// GetTokenVaultA returns the TokenVaultA field value
func (o *OrcaWhirlpoolConfig) GetTokenVaultA() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenVaultA
}

// GetTokenVaultAOk returns a tuple with the TokenVaultA field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetTokenVaultAOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TokenVaultA, true
}

// SetTokenVaultA sets field value
func (o *OrcaWhirlpoolConfig) SetTokenVaultA(v string) {
	o.TokenVaultA = v
}

// GetTokenVaultB returns the TokenVaultB field value
func (o *OrcaWhirlpoolConfig) GetTokenVaultB() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenVaultB
}

// GetTokenVaultBOk returns a tuple with the TokenVaultB field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetTokenVaultBOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.TokenVaultB, true
}

// SetTokenVaultB sets field value
func (o *OrcaWhirlpoolConfig) SetTokenVaultB(v string) {
	o.TokenVaultB = v
}

// GetOracle returns the Oracle field value
func (o *OrcaWhirlpoolConfig) GetOracle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Oracle
}

// GetOracleOk returns a tuple with the Oracle field value
// and a boolean to check if the value has been set.
func (o *OrcaWhirlpoolConfig) GetOracleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Oracle, true
}

// SetOracle sets field value
func (o *OrcaWhirlpoolConfig) SetOracle(v string) {
	o.Oracle = v
}

func (o OrcaWhirlpoolConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vault"] = o.Vault
	}
	if true {
		toSerialize["vaultProtoConfig"] = o.VaultProtoConfig
	}
	if true {
		toSerialize["vaultTokenAAccount"] = o.VaultTokenAAccount
	}
	if true {
		toSerialize["vaultTokenBAccount"] = o.VaultTokenBAccount
	}
	if true {
		toSerialize["tokenAMint"] = o.TokenAMint
	}
	if true {
		toSerialize["tokenBMint"] = o.TokenBMint
	}
	if true {
		toSerialize["whirlpool"] = o.Whirlpool
	}
	if true {
		toSerialize["tokenVaultA"] = o.TokenVaultA
	}
	if true {
		toSerialize["tokenVaultB"] = o.TokenVaultB
	}
	if true {
		toSerialize["oracle"] = o.Oracle
	}
	return json.Marshal(toSerialize)
}

type NullableOrcaWhirlpoolConfig struct {
	value *OrcaWhirlpoolConfig
	isSet bool
}

func (v NullableOrcaWhirlpoolConfig) Get() *OrcaWhirlpoolConfig {
	return v.value
}

func (v *NullableOrcaWhirlpoolConfig) Set(val *OrcaWhirlpoolConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableOrcaWhirlpoolConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableOrcaWhirlpoolConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrcaWhirlpoolConfig(val *OrcaWhirlpoolConfig) *NullableOrcaWhirlpoolConfig {
	return &NullableOrcaWhirlpoolConfig{value: val, isSet: true}
}

func (v NullableOrcaWhirlpoolConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrcaWhirlpoolConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


