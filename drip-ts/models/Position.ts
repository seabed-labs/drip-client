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
 * @interface Position
 */
export interface Position {
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    pubkey: string;
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    vault: string;
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    authority: string;
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    depositedTokenAAmount: string;
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    withdrawnTokenBAmount: string;
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    depositTimestamp: string;
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    dcaPeriodIdBeforeDeposit: string;
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    numberOfSwaps: string;
    /**
     * 
     * @type {string}
     * @memberof Position
     */
    periodicDripAmount: string;
    /**
     * 
     * @type {boolean}
     * @memberof Position
     */
    isClosed: boolean;
}

export function PositionFromJSON(json: any): Position {
    return PositionFromJSONTyped(json, false);
}

export function PositionFromJSONTyped(json: any, ignoreDiscriminator: boolean): Position {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pubkey': json['pubkey'],
        'vault': json['vault'],
        'authority': json['authority'],
        'depositedTokenAAmount': json['depositedTokenAAmount'],
        'withdrawnTokenBAmount': json['withdrawnTokenBAmount'],
        'depositTimestamp': json['depositTimestamp'],
        'dcaPeriodIdBeforeDeposit': json['dcaPeriodIdBeforeDeposit'],
        'numberOfSwaps': json['numberOfSwaps'],
        'periodicDripAmount': json['periodicDripAmount'],
        'isClosed': json['isClosed'],
    };
}

export function PositionToJSON(value?: Position | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pubkey': value.pubkey,
        'vault': value.vault,
        'authority': value.authority,
        'depositedTokenAAmount': value.depositedTokenAAmount,
        'withdrawnTokenBAmount': value.withdrawnTokenBAmount,
        'depositTimestamp': value.depositTimestamp,
        'dcaPeriodIdBeforeDeposit': value.dcaPeriodIdBeforeDeposit,
        'numberOfSwaps': value.numberOfSwaps,
        'periodicDripAmount': value.periodicDripAmount,
        'isClosed': value.isClosed,
    };
}

