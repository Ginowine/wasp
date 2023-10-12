/**
 * Wasp API
 * REST API for the Wasp node
 *
 * OpenAPI spec version: 0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { HttpFile } from '../http/http';

export class NFTDataResponse {
    'id': string;
    'issuer': string;
    'metadata': string;
    'owner': string;

    static readonly discriminator: string | undefined = undefined;

    static readonly attributeTypeMap: Array<{name: string, baseName: string, type: string, format: string}> = [
        {
            "name": "id",
            "baseName": "id",
            "type": "string",
            "format": "string"
        },
        {
            "name": "issuer",
            "baseName": "issuer",
            "type": "string",
            "format": "string"
        },
        {
            "name": "metadata",
            "baseName": "metadata",
            "type": "string",
            "format": "string"
        },
        {
            "name": "owner",
            "baseName": "owner",
            "type": "string",
            "format": "string"
        }    ];

    static getAttributeTypeMap() {
        return NFTDataResponse.attributeTypeMap;
    }

    public constructor() {
    }
}
