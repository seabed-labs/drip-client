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
 * @interface Token
 */
export interface Token {
    /**
     * 
     * @type {string}
     * @memberof Token
     */
    pubkey: string;
    /**
     * 
     * @type {number}
     * @memberof Token
     */
    decimals: number;
    /**
     * 
     * @type {string}
     * @memberof Token
     */
    symbol?: string;
    /**
     * 
     * @type {string}
     * @memberof Token
     */
    iconUrl?: string;
    /**
     * 
     * @type {string}
     * @memberof Token
     */
    coinGeckoId?: string;
}

export function TokenFromJSON(json: any): Token {
    return TokenFromJSONTyped(json, false);
}

export function TokenFromJSONTyped(json: any, ignoreDiscriminator: boolean): Token {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pubkey': json['pubkey'],
        'decimals': json['decimals'],
        'symbol': !exists(json, 'symbol') ? undefined : json['symbol'],
        'iconUrl': !exists(json, 'iconUrl') ? undefined : json['iconUrl'],
        'coinGeckoId': !exists(json, 'coinGeckoId') ? undefined : json['coinGeckoId'],
    };
}

export function TokenToJSON(value?: Token | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pubkey': value.pubkey,
        'decimals': value.decimals,
        'symbol': value.symbol,
        'iconUrl': value.iconUrl,
        'coinGeckoId': value.coinGeckoId,
    };
}

