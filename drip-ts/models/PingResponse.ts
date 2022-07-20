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
 * @interface PingResponse
 */
export interface PingResponse {
    /**
     * 
     * @type {string}
     * @memberof PingResponse
     */
    message: string;
}

export function PingResponseFromJSON(json: any): PingResponse {
    return PingResponseFromJSONTyped(json, false);
}

export function PingResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): PingResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'message': json['message'],
    };
}

export function PingResponseToJSON(value?: PingResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'message': value.message,
    };
}

