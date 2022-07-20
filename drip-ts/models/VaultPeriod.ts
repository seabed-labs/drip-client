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
 * @interface VaultPeriod
 */
export interface VaultPeriod {
    /**
     * 
     * @type {string}
     * @memberof VaultPeriod
     */
    pubkey: string;
    /**
     * 
     * @type {string}
     * @memberof VaultPeriod
     */
    vault: string;
    /**
     * 
     * @type {string}
     * @memberof VaultPeriod
     */
    periodId: string;
    /**
     * 
     * @type {string}
     * @memberof VaultPeriod
     */
    twap: string;
    /**
     * 
     * @type {string}
     * @memberof VaultPeriod
     */
    dar: string;
}

export function VaultPeriodFromJSON(json: any): VaultPeriod {
    return VaultPeriodFromJSONTyped(json, false);
}

export function VaultPeriodFromJSONTyped(json: any, ignoreDiscriminator: boolean): VaultPeriod {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pubkey': json['pubkey'],
        'vault': json['vault'],
        'periodId': json['periodId'],
        'twap': json['twap'],
        'dar': json['dar'],
    };
}

export function VaultPeriodToJSON(value?: VaultPeriod | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pubkey': value.pubkey,
        'vault': value.vault,
        'periodId': value.periodId,
        'twap': value.twap,
        'dar': value.dar,
    };
}
