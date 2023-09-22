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
 * @interface SplTokenSwapConfigAllOf
 */
export interface SplTokenSwapConfigAllOf {
    /**
     * 
     * @type {string}
     * @memberof SplTokenSwapConfigAllOf
     */
    swapTokenMint: string;
    /**
     * 
     * @type {string}
     * @memberof SplTokenSwapConfigAllOf
     */
    swapTokenAAccount: string;
    /**
     * 
     * @type {string}
     * @memberof SplTokenSwapConfigAllOf
     */
    swapTokenBAccount: string;
    /**
     * 
     * @type {string}
     * @memberof SplTokenSwapConfigAllOf
     */
    swapFeeAccount: string;
    /**
     * 
     * @type {string}
     * @memberof SplTokenSwapConfigAllOf
     */
    swapAuthority: string;
    /**
     * 
     * @type {string}
     * @memberof SplTokenSwapConfigAllOf
     */
    swap: string;
}

export function SplTokenSwapConfigAllOfFromJSON(json: any): SplTokenSwapConfigAllOf {
    return SplTokenSwapConfigAllOfFromJSONTyped(json, false);
}

export function SplTokenSwapConfigAllOfFromJSONTyped(json: any, ignoreDiscriminator: boolean): SplTokenSwapConfigAllOf {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'swapTokenMint': json['swapTokenMint'],
        'swapTokenAAccount': json['swapTokenAAccount'],
        'swapTokenBAccount': json['swapTokenBAccount'],
        'swapFeeAccount': json['swapFeeAccount'],
        'swapAuthority': json['swapAuthority'],
        'swap': json['swap'],
    };
}

export function SplTokenSwapConfigAllOfToJSON(value?: SplTokenSwapConfigAllOf | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'swapTokenMint': value.swapTokenMint,
        'swapTokenAAccount': value.swapTokenAAccount,
        'swapTokenBAccount': value.swapTokenBAccount,
        'swapFeeAccount': value.swapFeeAccount,
        'swapAuthority': value.swapAuthority,
        'swap': value.swap,
    };
}

