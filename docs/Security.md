# Security

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThreatScore** | Pointer to **int32** |  | [optional] 
**IsTor** | Pointer to **bool** |  | [optional] 
**IsProxy** | Pointer to **bool** |  | [optional] 
**ProxyType** | Pointer to **string** |  | [optional] 
**ProxyProvider** | Pointer to **string** |  | [optional] 
**IsAnonymous** | Pointer to **bool** |  | [optional] 
**IsKnownAttacker** | Pointer to **bool** |  | [optional] 
**IsSpam** | Pointer to **bool** |  | [optional] 
**IsBot** | Pointer to **bool** |  | [optional] 
**IsCloudProvider** | Pointer to **bool** |  | [optional] 
**CloudProvider** | Pointer to **string** |  | [optional] 

## Methods

### NewSecurity

`func NewSecurity() *Security`

NewSecurity instantiates a new Security object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityWithDefaults

`func NewSecurityWithDefaults() *Security`

NewSecurityWithDefaults instantiates a new Security object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThreatScore

`func (o *Security) GetThreatScore() int32`

GetThreatScore returns the ThreatScore field if non-nil, zero value otherwise.

### GetThreatScoreOk

`func (o *Security) GetThreatScoreOk() (*int32, bool)`

GetThreatScoreOk returns a tuple with the ThreatScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreatScore

`func (o *Security) SetThreatScore(v int32)`

SetThreatScore sets ThreatScore field to given value.

### HasThreatScore

`func (o *Security) HasThreatScore() bool`

HasThreatScore returns a boolean if a field has been set.

### GetIsTor

`func (o *Security) GetIsTor() bool`

GetIsTor returns the IsTor field if non-nil, zero value otherwise.

### GetIsTorOk

`func (o *Security) GetIsTorOk() (*bool, bool)`

GetIsTorOk returns a tuple with the IsTor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTor

`func (o *Security) SetIsTor(v bool)`

SetIsTor sets IsTor field to given value.

### HasIsTor

`func (o *Security) HasIsTor() bool`

HasIsTor returns a boolean if a field has been set.

### GetIsProxy

`func (o *Security) GetIsProxy() bool`

GetIsProxy returns the IsProxy field if non-nil, zero value otherwise.

### GetIsProxyOk

`func (o *Security) GetIsProxyOk() (*bool, bool)`

GetIsProxyOk returns a tuple with the IsProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsProxy

`func (o *Security) SetIsProxy(v bool)`

SetIsProxy sets IsProxy field to given value.

### HasIsProxy

`func (o *Security) HasIsProxy() bool`

HasIsProxy returns a boolean if a field has been set.

### GetProxyType

`func (o *Security) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *Security) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *Security) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *Security) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetProxyProvider

`func (o *Security) GetProxyProvider() string`

GetProxyProvider returns the ProxyProvider field if non-nil, zero value otherwise.

### GetProxyProviderOk

`func (o *Security) GetProxyProviderOk() (*string, bool)`

GetProxyProviderOk returns a tuple with the ProxyProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProvider

`func (o *Security) SetProxyProvider(v string)`

SetProxyProvider sets ProxyProvider field to given value.

### HasProxyProvider

`func (o *Security) HasProxyProvider() bool`

HasProxyProvider returns a boolean if a field has been set.

### GetIsAnonymous

`func (o *Security) GetIsAnonymous() bool`

GetIsAnonymous returns the IsAnonymous field if non-nil, zero value otherwise.

### GetIsAnonymousOk

`func (o *Security) GetIsAnonymousOk() (*bool, bool)`

GetIsAnonymousOk returns a tuple with the IsAnonymous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnonymous

`func (o *Security) SetIsAnonymous(v bool)`

SetIsAnonymous sets IsAnonymous field to given value.

### HasIsAnonymous

`func (o *Security) HasIsAnonymous() bool`

HasIsAnonymous returns a boolean if a field has been set.

### GetIsKnownAttacker

`func (o *Security) GetIsKnownAttacker() bool`

GetIsKnownAttacker returns the IsKnownAttacker field if non-nil, zero value otherwise.

### GetIsKnownAttackerOk

`func (o *Security) GetIsKnownAttackerOk() (*bool, bool)`

GetIsKnownAttackerOk returns a tuple with the IsKnownAttacker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsKnownAttacker

`func (o *Security) SetIsKnownAttacker(v bool)`

SetIsKnownAttacker sets IsKnownAttacker field to given value.

### HasIsKnownAttacker

`func (o *Security) HasIsKnownAttacker() bool`

HasIsKnownAttacker returns a boolean if a field has been set.

### GetIsSpam

`func (o *Security) GetIsSpam() bool`

GetIsSpam returns the IsSpam field if non-nil, zero value otherwise.

### GetIsSpamOk

`func (o *Security) GetIsSpamOk() (*bool, bool)`

GetIsSpamOk returns a tuple with the IsSpam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSpam

`func (o *Security) SetIsSpam(v bool)`

SetIsSpam sets IsSpam field to given value.

### HasIsSpam

`func (o *Security) HasIsSpam() bool`

HasIsSpam returns a boolean if a field has been set.

### GetIsBot

`func (o *Security) GetIsBot() bool`

GetIsBot returns the IsBot field if non-nil, zero value otherwise.

### GetIsBotOk

`func (o *Security) GetIsBotOk() (*bool, bool)`

GetIsBotOk returns a tuple with the IsBot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBot

`func (o *Security) SetIsBot(v bool)`

SetIsBot sets IsBot field to given value.

### HasIsBot

`func (o *Security) HasIsBot() bool`

HasIsBot returns a boolean if a field has been set.

### GetIsCloudProvider

`func (o *Security) GetIsCloudProvider() bool`

GetIsCloudProvider returns the IsCloudProvider field if non-nil, zero value otherwise.

### GetIsCloudProviderOk

`func (o *Security) GetIsCloudProviderOk() (*bool, bool)`

GetIsCloudProviderOk returns a tuple with the IsCloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCloudProvider

`func (o *Security) SetIsCloudProvider(v bool)`

SetIsCloudProvider sets IsCloudProvider field to given value.

### HasIsCloudProvider

`func (o *Security) HasIsCloudProvider() bool`

HasIsCloudProvider returns a boolean if a field has been set.

### GetCloudProvider

`func (o *Security) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *Security) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *Security) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *Security) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


