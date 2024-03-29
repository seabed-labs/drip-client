/* tslint:disable */
/* eslint-disable */
/**
 * Drip Backend
 * Drip backend service. All API\'s have a IP rate limit of 10 requests per second. 
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: mocha@dcaf.so
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface Vault
 */
export interface Vault {
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    pubkey: string;
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    protoConfig: string;
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    tokenAAccount: string;
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    tokenBAccount: string;
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    treasuryTokenBAccount: string;
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    tokenAMint: string;
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    tokenBMint: string;
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    lastDcaPeriod: string;
    /**
     * 
     * @type {string}
     * @memberof Vault
     */
    dripAmount: string;
    /**
     * unix timestamp
     * @type {string}
     * @memberof Vault
     */
    dcaActivationTimestamp: string;
    /**
     * 
     * @type {boolean}
     * @memberof Vault
     */
    enabled: boolean;
    /**
     * 
     * @type {number}
     * @memberof Vault
     */
    maxSlippageBps: number;
}

export function VaultFromJSON(json: any): Vault {
    return VaultFromJSONTyped(json, false);
}

export function VaultFromJSONTyped(json: any, ignoreDiscriminator: boolean): Vault {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pubkey': json['pubkey'],
        'protoConfig': json['protoConfig'],
        'tokenAAccount': json['tokenAAccount'],
        'tokenBAccount': json['tokenBAccount'],
        'treasuryTokenBAccount': json['treasuryTokenBAccount'],
        'tokenAMint': json['tokenAMint'],
        'tokenBMint': json['tokenBMint'],
        'lastDcaPeriod': json['lastDcaPeriod'],
        'dripAmount': json['dripAmount'],
        'dcaActivationTimestamp': json['dcaActivationTimestamp'],
        'enabled': json['enabled'],
        'maxSlippageBps': json['maxSlippageBps'],
    };
}

export function VaultToJSON(value?: Vault | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pubkey': value.pubkey,
        'protoConfig': value.protoConfig,
        'tokenAAccount': value.tokenAAccount,
        'tokenBAccount': value.tokenBAccount,
        'treasuryTokenBAccount': value.treasuryTokenBAccount,
        'tokenAMint': value.tokenAMint,
        'tokenBMint': value.tokenBMint,
        'lastDcaPeriod': value.lastDcaPeriod,
        'dripAmount': value.dripAmount,
        'dcaActivationTimestamp': value.dcaActivationTimestamp,
        'enabled': value.enabled,
        'maxSlippageBps': value.maxSlippageBps,
    };
}

