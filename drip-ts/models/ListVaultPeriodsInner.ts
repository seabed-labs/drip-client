/* tslint:disable */
/* eslint-disable */
/**
 * Drip Backend
 * Drip backend service.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: dcafmocha@protonmail.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface ListVaultPeriodsInner
 */
export interface ListVaultPeriodsInner {
    /**
     * 
     * @type {string}
     * @memberof ListVaultPeriodsInner
     */
    pubkey: string;
    /**
     * 
     * @type {string}
     * @memberof ListVaultPeriodsInner
     */
    vault: string;
    /**
     * 
     * @type {string}
     * @memberof ListVaultPeriodsInner
     */
    periodId: string;
    /**
     * 
     * @type {string}
     * @memberof ListVaultPeriodsInner
     */
    twap: string;
    /**
     * 
     * @type {string}
     * @memberof ListVaultPeriodsInner
     */
    dar: string;
}

export function ListVaultPeriodsInnerFromJSON(json: any): ListVaultPeriodsInner {
    return ListVaultPeriodsInnerFromJSONTyped(json, false);
}

export function ListVaultPeriodsInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListVaultPeriodsInner {
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

export function ListVaultPeriodsInnerToJSON(value?: ListVaultPeriodsInner | null): any {
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
