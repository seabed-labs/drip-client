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
 * @interface TokenAccount
 */
export interface TokenAccount {
    /**
     * 
     * @type {string}
     * @memberof TokenAccount
     */
    pubkey: string;
    /**
     * 
     * @type {string}
     * @memberof TokenAccount
     */
    mint: string;
    /**
     * 
     * @type {string}
     * @memberof TokenAccount
     */
    owner: string;
    /**
     * 
     * @type {string}
     * @memberof TokenAccount
     */
    amount: string;
    /**
     * 
     * @type {string}
     * @memberof TokenAccount
     */
    state: string;
}

export function TokenAccountFromJSON(json: any): TokenAccount {
    return TokenAccountFromJSONTyped(json, false);
}

export function TokenAccountFromJSONTyped(json: any, ignoreDiscriminator: boolean): TokenAccount {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pubkey': json['pubkey'],
        'mint': json['mint'],
        'owner': json['owner'],
        'amount': json['amount'],
        'state': json['state'],
    };
}

export function TokenAccountToJSON(value?: TokenAccount | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pubkey': value.pubkey,
        'mint': value.mint,
        'owner': value.owner,
        'amount': value.amount,
        'state': value.state,
    };
}

