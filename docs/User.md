# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [optional] 
**Name** | **string** |  | [optional] [default to ]
**FamilyName** | **string** |  | [optional] [default to ]
**CreatedOn** | [**time.Time**](time.Time.md) |  | [optional] 
**ModifiedOn** | [**time.Time**](time.Time.md) |  | [optional] 
**PasswordReset** | [**UserPasswordReset**](user_passwordReset.md) |  | [optional] 
**Verified** | **bool** |  | [optional] [default to false]
**Lang** | **string** |  | [optional] [default to LANG_EN]
**Phone** | **string** |  | [optional] [default to ]
**LastLogin** | [**[]UserLastLogin**](user_lastLogin.md) |  | [optional] 
**Limit** | [**UserLimit**](user_limit.md) |  | [optional] 
**Credential** | [**UserCredential**](user_credential.md) |  | [optional] 
**NetworkAcl** | [**[]UserNetworkAcl**](user_networkAcl.md) |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


