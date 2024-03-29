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
 * @interface AnalyticsOverviewResponse
 */
export interface AnalyticsOverviewResponse {
    /**
     * A float 64 string for the current tvl in USD.
     * @type {string}
     * @memberof AnalyticsOverviewResponse
     */
    usdTvl: string;
    /**
     * A float 64 string for life time deposits (normalized to the current usd price of assets).
     * @type {string}
     * @memberof AnalyticsOverviewResponse
     */
    lifeTimeUsdDeposit: string;
    /**
     * A float 64 string for life time withdrawals (normalized to the current usd price of assets).
     * @type {string}
     * @memberof AnalyticsOverviewResponse
     */
    lifeTimeUsdWithdrawal: string;
    /**
     * A float 64 string for life time volume (normalized to the current usd price of assets).
     * @type {string}
     * @memberof AnalyticsOverviewResponse
     */
    lifeTimeUsdVolume: string;
    /**
     * Number of wallets with active postions.
     * @type {number}
     * @memberof AnalyticsOverviewResponse
     */
    activeWallets: number;
    /**
     * Total number of unique depositors.
     * @type {number}
     * @memberof AnalyticsOverviewResponse
     */
    uniqueDepositors: number;
}

export function AnalyticsOverviewResponseFromJSON(json: any): AnalyticsOverviewResponse {
    return AnalyticsOverviewResponseFromJSONTyped(json, false);
}

export function AnalyticsOverviewResponseFromJSONTyped(json: any, ignoreDiscriminator: boolean): AnalyticsOverviewResponse {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'usdTvl': json['usdTvl'],
        'lifeTimeUsdDeposit': json['lifeTimeUsdDeposit'],
        'lifeTimeUsdWithdrawal': json['lifeTimeUsdWithdrawal'],
        'lifeTimeUsdVolume': json['lifeTimeUsdVolume'],
        'activeWallets': json['activeWallets'],
        'uniqueDepositors': json['uniqueDepositors'],
    };
}

export function AnalyticsOverviewResponseToJSON(value?: AnalyticsOverviewResponse | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'usdTvl': value.usdTvl,
        'lifeTimeUsdDeposit': value.lifeTimeUsdDeposit,
        'lifeTimeUsdWithdrawal': value.lifeTimeUsdWithdrawal,
        'lifeTimeUsdVolume': value.lifeTimeUsdVolume,
        'activeWallets': value.activeWallets,
        'uniqueDepositors': value.uniqueDepositors,
    };
}

