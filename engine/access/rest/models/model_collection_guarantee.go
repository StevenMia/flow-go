/*
 * Access API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models

type CollectionGuarantee struct {
	CollectionId  string `json:"collection_id"`
	SignerIndices string `json:"signer_indices"`
	Signature     string `json:"signature"`
}
