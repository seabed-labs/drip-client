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
 * @interface ListProtoConfigsInner
 */
export interface ListProtoConfigsInner {
    /**
     * 
     * @type {string}
     * @memberof ListProtoConfigsInner
     */
    pubkey: string;
    /**
     * 
     * @type {number}
     * @memberof ListProtoConfigsInner
     */
    granularity: number;
    /**
     * 
     * @type {number}
     * @memberof ListProtoConfigsInner
     */
    triggerDcaSpread: number;
    /**
     * 
     * @type {number}
     * @memberof ListProtoConfigsInner
     */
    baseWithdrawalSpread: number;
}

export function ListProtoConfigsInnerFromJSON(json: any): ListProtoConfigsInner {
    return ListProtoConfigsInnerFromJSONTyped(json, false);
}

export function ListProtoConfigsInnerFromJSONTyped(json: any, ignoreDiscriminator: boolean): ListProtoConfigsInner {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'pubkey': json['pubkey'],
        'granularity': json['granularity'],
        'triggerDcaSpread': json['triggerDcaSpread'],
        'baseWithdrawalSpread': json['baseWithdrawalSpread'],
    };
}

export function ListProtoConfigsInnerToJSON(value?: ListProtoConfigsInner | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'pubkey': value.pubkey,
        'granularity': value.granularity,
        'triggerDcaSpread': value.triggerDcaSpread,
        'baseWithdrawalSpread': value.baseWithdrawalSpread,
    };
}
