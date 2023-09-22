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


import * as runtime from '../runtime';
import {
    ActiveWallet,
    ActiveWalletFromJSON,
    ActiveWalletToJSON,
    ErrorResponse,
    ErrorResponseFromJSON,
    ErrorResponseToJSON,
    ExpandedAdminPosition,
    ExpandedAdminPositionFromJSON,
    ExpandedAdminPositionToJSON,
    ExpandedAdminVault,
    ExpandedAdminVaultFromJSON,
    ExpandedAdminVaultToJSON,
    Vault,
    VaultFromJSON,
    VaultToJSON,
} from '../models';

export interface V1AdminPositionsGetRequest {
    tokenId: string;
    expand?: Array<V1AdminPositionsGetExpandEnum>;
    enabled?: boolean;
    isClosed?: boolean;
    offset?: number;
    limit?: number;
}

export interface V1AdminSummaryActivewalletsGetRequest {
    tokenId: string;
    vault?: string;
    isClosed?: boolean;
    owner?: string;
}

export interface V1AdminVaultPubkeyPathEnablePutRequest {
    pubkeyPath: string;
    tokenId: string;
}

export interface V1AdminVaultsGetRequest {
    tokenId: string;
    expand?: Array<V1AdminVaultsGetExpandEnum>;
    vaultProtoConfig?: string;
    vault?: string;
    tokenA?: string;
    tokenB?: string;
    enabled?: boolean;
    offset?: number;
    limit?: number;
}

/**
 * 
 */
export class AdminApi extends runtime.BaseAPI {

    /**
     * Get all positions with pagination.
     * Get All Positions
     */
    async v1AdminPositionsGetRaw(requestParameters: V1AdminPositionsGetRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<Array<ExpandedAdminPosition>>> {
        if (requestParameters.tokenId === null || requestParameters.tokenId === undefined) {
            throw new runtime.RequiredError('tokenId','Required parameter requestParameters.tokenId was null or undefined when calling v1AdminPositionsGet.');
        }

        const queryParameters: any = {};

        if (requestParameters.expand) {
            queryParameters['expand'] = requestParameters.expand;
        }

        if (requestParameters.enabled !== undefined) {
            queryParameters['enabled'] = requestParameters.enabled;
        }

        if (requestParameters.isClosed !== undefined) {
            queryParameters['isClosed'] = requestParameters.isClosed;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.tokenId !== undefined && requestParameters.tokenId !== null) {
            headerParameters['token-id'] = String(requestParameters.tokenId);
        }

        const response = await this.request({
            path: `/v1/admin/positions`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(ExpandedAdminPositionFromJSON));
    }

    /**
     * Get all positions with pagination.
     * Get All Positions
     */
    async v1AdminPositionsGet(requestParameters: V1AdminPositionsGetRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<Array<ExpandedAdminPosition>> {
        const response = await this.v1AdminPositionsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get all wallet addresses with open positions.
     * Get All Active Wallet Addresses
     */
    async v1AdminSummaryActivewalletsGetRaw(requestParameters: V1AdminSummaryActivewalletsGetRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<Array<ActiveWallet>>> {
        if (requestParameters.tokenId === null || requestParameters.tokenId === undefined) {
            throw new runtime.RequiredError('tokenId','Required parameter requestParameters.tokenId was null or undefined when calling v1AdminSummaryActivewalletsGet.');
        }

        const queryParameters: any = {};

        if (requestParameters.vault !== undefined) {
            queryParameters['vault'] = requestParameters.vault;
        }

        if (requestParameters.isClosed !== undefined) {
            queryParameters['isClosed'] = requestParameters.isClosed;
        }

        if (requestParameters.owner !== undefined) {
            queryParameters['owner'] = requestParameters.owner;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.tokenId !== undefined && requestParameters.tokenId !== null) {
            headerParameters['token-id'] = String(requestParameters.tokenId);
        }

        const response = await this.request({
            path: `/v1/admin/summary/activewallets`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(ActiveWalletFromJSON));
    }

    /**
     * Get all wallet addresses with open positions.
     * Get All Active Wallet Addresses
     */
    async v1AdminSummaryActivewalletsGet(requestParameters: V1AdminSummaryActivewalletsGetRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<Array<ActiveWallet>> {
        const response = await this.v1AdminSummaryActivewalletsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Enable/disable the specified vault
     * Toggle the \'enabled\' flag on a vault
     */
    async v1AdminVaultPubkeyPathEnablePutRaw(requestParameters: V1AdminVaultPubkeyPathEnablePutRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<Vault>> {
        if (requestParameters.pubkeyPath === null || requestParameters.pubkeyPath === undefined) {
            throw new runtime.RequiredError('pubkeyPath','Required parameter requestParameters.pubkeyPath was null or undefined when calling v1AdminVaultPubkeyPathEnablePut.');
        }

        if (requestParameters.tokenId === null || requestParameters.tokenId === undefined) {
            throw new runtime.RequiredError('tokenId','Required parameter requestParameters.tokenId was null or undefined when calling v1AdminVaultPubkeyPathEnablePut.');
        }

        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.tokenId !== undefined && requestParameters.tokenId !== null) {
            headerParameters['token-id'] = String(requestParameters.tokenId);
        }

        const response = await this.request({
            path: `/v1/admin/vault/{pubkeyPath}/enable`.replace(`{${"pubkeyPath"}}`, encodeURIComponent(String(requestParameters.pubkeyPath))),
            method: 'PUT',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => VaultFromJSON(jsonValue));
    }

    /**
     * Enable/disable the specified vault
     * Toggle the \'enabled\' flag on a vault
     */
    async v1AdminVaultPubkeyPathEnablePut(requestParameters: V1AdminVaultPubkeyPathEnablePutRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<Vault> {
        const response = await this.v1AdminVaultPubkeyPathEnablePutRaw(requestParameters, initOverrides);
        return await response.value();
    }

    /**
     * Get all vaults with filters and expanded properties.
     * Get All Vaults
     */
    async v1AdminVaultsGetRaw(requestParameters: V1AdminVaultsGetRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<runtime.ApiResponse<Array<ExpandedAdminVault>>> {
        if (requestParameters.tokenId === null || requestParameters.tokenId === undefined) {
            throw new runtime.RequiredError('tokenId','Required parameter requestParameters.tokenId was null or undefined when calling v1AdminVaultsGet.');
        }

        const queryParameters: any = {};

        if (requestParameters.expand) {
            queryParameters['expand'] = requestParameters.expand;
        }

        if (requestParameters.vaultProtoConfig !== undefined) {
            queryParameters['vaultProtoConfig'] = requestParameters.vaultProtoConfig;
        }

        if (requestParameters.vault !== undefined) {
            queryParameters['vault'] = requestParameters.vault;
        }

        if (requestParameters.tokenA !== undefined) {
            queryParameters['tokenA'] = requestParameters.tokenA;
        }

        if (requestParameters.tokenB !== undefined) {
            queryParameters['tokenB'] = requestParameters.tokenB;
        }

        if (requestParameters.enabled !== undefined) {
            queryParameters['enabled'] = requestParameters.enabled;
        }

        if (requestParameters.offset !== undefined) {
            queryParameters['offset'] = requestParameters.offset;
        }

        if (requestParameters.limit !== undefined) {
            queryParameters['limit'] = requestParameters.limit;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        if (requestParameters.tokenId !== undefined && requestParameters.tokenId !== null) {
            headerParameters['token-id'] = String(requestParameters.tokenId);
        }

        const response = await this.request({
            path: `/v1/admin/vaults`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(ExpandedAdminVaultFromJSON));
    }

    /**
     * Get all vaults with filters and expanded properties.
     * Get All Vaults
     */
    async v1AdminVaultsGet(requestParameters: V1AdminVaultsGetRequest, initOverrides?: RequestInit | runtime.InitOverideFunction): Promise<Array<ExpandedAdminVault>> {
        const response = await this.v1AdminVaultsGetRaw(requestParameters, initOverrides);
        return await response.value();
    }

}

/**
 * @export
 */
export const V1AdminPositionsGetExpandEnum = {
    All: 'all',
    Vault: 'vault',
    ProtoConfig: 'protoConfig',
    TokenA: 'tokenA',
    TokenB: 'tokenB'
} as const;
export type V1AdminPositionsGetExpandEnum = typeof V1AdminPositionsGetExpandEnum[keyof typeof V1AdminPositionsGetExpandEnum];
/**
 * @export
 */
export const V1AdminVaultsGetExpandEnum = {
    All: 'all',
    ProtoConfigValue: 'protoConfigValue',
    TokenAMintValue: 'tokenAMintValue',
    TokenBMintValue: 'tokenBMintValue',
    TokenAAccountValue: 'tokenAAccountValue',
    TokenBAccountValue: 'tokenBAccountValue',
    TreasuryTokenBAccountValue: 'treasuryTokenBAccountValue'
} as const;
export type V1AdminVaultsGetExpandEnum = typeof V1AdminVaultsGetExpandEnum[keyof typeof V1AdminVaultsGetExpandEnum];
