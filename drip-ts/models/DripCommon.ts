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
 * @interface DripCommon
 */
export interface DripCommon {
    /**
     * 
     * @type {string}
     * @memberof DripCommon
     */
    vault: string;
    /**
     * 
     * @type {string}
     * @memberof DripCommon
     */
    vaultProtoConfig: string;
    /**
     * 
     * @type {string}
     * @memberof DripCommon
     */
    vaultTokenAAccount: string;
    /**
     * 
     * @type {string}
     * @memberof DripCommon
     */
    vaultTokenBAccount: string;
    /**
     * 
     * @type {string}
     * @memberof DripCommon
     */
    tokenAMint: string;
    /**
     * 
     * @type {string}
     * @memberof DripCommon
     */
    tokenBMint: string;
    /**
     * 
     * @type {string}
     * @memberof DripCommon
     */
    oracleConfig?: string;
}

/**
 * Check if a given object implements the DripCommon interface.
 */
export function instanceOfDripCommon(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "vault" in value;
    isInstance = isInstance && "vaultProtoConfig" in value;
    isInstance = isInstance && "vaultTokenAAccount" in value;
    isInstance = isInstance && "vaultTokenBAccount" in value;
    isInstance = isInstance && "tokenAMint" in value;
    isInstance = isInstance && "tokenBMint" in value;

    return isInstance;
}

export function DripCommonFromJSON(json: any): DripCommon {
    return DripCommonFromJSONTyped(json, false);
}

export function DripCommonFromJSONTyped(json: any, ignoreDiscriminator: boolean): DripCommon {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'vault': json['vault'],
        'vaultProtoConfig': json['vaultProtoConfig'],
        'vaultTokenAAccount': json['vaultTokenAAccount'],
        'vaultTokenBAccount': json['vaultTokenBAccount'],
        'tokenAMint': json['tokenAMint'],
        'tokenBMint': json['tokenBMint'],
        'oracleConfig': !exists(json, 'oracleConfig') ? undefined : json['oracleConfig'],
    };
}

export function DripCommonToJSON(value?: DripCommon | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'vault': value.vault,
        'vaultProtoConfig': value.vaultProtoConfig,
        'vaultTokenAAccount': value.vaultTokenAAccount,
        'vaultTokenBAccount': value.vaultTokenBAccount,
        'tokenAMint': value.tokenAMint,
        'tokenBMint': value.tokenBMint,
        'oracleConfig': value.oracleConfig,
    };
}

