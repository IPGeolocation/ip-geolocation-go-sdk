# UserAgentData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserAgentString** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**VersionMajor** | Pointer to **string** |  | [optional] 
**Device** | Pointer to [**UserAgentDataDevice**](UserAgentDataDevice.md) |  | [optional] 
**Engine** | Pointer to [**UserAgentDataEngine**](UserAgentDataEngine.md) |  | [optional] 
**OperatingSystem** | Pointer to [**UserAgentDataOperatingSystem**](UserAgentDataOperatingSystem.md) |  | [optional] 

## Methods

### NewUserAgentData

`func NewUserAgentData() *UserAgentData`

NewUserAgentData instantiates a new UserAgentData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAgentDataWithDefaults

`func NewUserAgentDataWithDefaults() *UserAgentData`

NewUserAgentDataWithDefaults instantiates a new UserAgentData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAgentString

`func (o *UserAgentData) GetUserAgentString() string`

GetUserAgentString returns the UserAgentString field if non-nil, zero value otherwise.

### GetUserAgentStringOk

`func (o *UserAgentData) GetUserAgentStringOk() (*string, bool)`

GetUserAgentStringOk returns a tuple with the UserAgentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentString

`func (o *UserAgentData) SetUserAgentString(v string)`

SetUserAgentString sets UserAgentString field to given value.

### HasUserAgentString

`func (o *UserAgentData) HasUserAgentString() bool`

HasUserAgentString returns a boolean if a field has been set.

### GetName

`func (o *UserAgentData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAgentData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAgentData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserAgentData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UserAgentData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAgentData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAgentData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserAgentData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *UserAgentData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserAgentData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserAgentData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserAgentData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionMajor

`func (o *UserAgentData) GetVersionMajor() string`

GetVersionMajor returns the VersionMajor field if non-nil, zero value otherwise.

### GetVersionMajorOk

`func (o *UserAgentData) GetVersionMajorOk() (*string, bool)`

GetVersionMajorOk returns a tuple with the VersionMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMajor

`func (o *UserAgentData) SetVersionMajor(v string)`

SetVersionMajor sets VersionMajor field to given value.

### HasVersionMajor

`func (o *UserAgentData) HasVersionMajor() bool`

HasVersionMajor returns a boolean if a field has been set.

### GetDevice

`func (o *UserAgentData) GetDevice() UserAgentDataDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *UserAgentData) GetDeviceOk() (*UserAgentDataDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *UserAgentData) SetDevice(v UserAgentDataDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *UserAgentData) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetEngine

`func (o *UserAgentData) GetEngine() UserAgentDataEngine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *UserAgentData) GetEngineOk() (*UserAgentDataEngine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *UserAgentData) SetEngine(v UserAgentDataEngine)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *UserAgentData) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *UserAgentData) GetOperatingSystem() UserAgentDataOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *UserAgentData) GetOperatingSystemOk() (*UserAgentDataOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *UserAgentData) SetOperatingSystem(v UserAgentDataOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *UserAgentData) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


