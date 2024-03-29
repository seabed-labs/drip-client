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
 * @interface ListTokensInner
 */
export interface ListTokensInner {
    /**
     * 
     * @type {string}
     * @memberof ListTokensInner
     */
    pubkey: string;
    /**
     * 
     * @type {string}
     * @memberof ListTokensInner
     */
    symbol: string;
    /**
     * 
     * @type {number}
     * @memberof ListTokensInner
     */
    decimals: number;
}

export function ListTokensInnerFromJSON(json: any): ListTokensInner {
    return ListTokensInnerFromJSONTyped(json, false);
}

export function ListTokensInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListTokensInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pubkey': json['pubkey'],
        'symbol': json['symbol'],
        'decimals': json['decimals'],
    };
}

export function ListTokensInnerToJSON(value?: ListTokensInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pubkey': value.pubkey,
        'symbol': value.symbol,
        'decimals': value.decimals,
    };
}

