# UserAgentXMLDataArray

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

### NewUserAgentXMLDataArray

`func NewUserAgentXMLDataArray() *UserAgentXMLDataArray`

NewUserAgentXMLDataArray instantiates a new UserAgentXMLDataArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserAgentXMLDataArrayWithDefaults

`func NewUserAgentXMLDataArrayWithDefaults() *UserAgentXMLDataArray`

NewUserAgentXMLDataArrayWithDefaults instantiates a new UserAgentXMLDataArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserAgentString

`func (o *UserAgentXMLDataArray) GetUserAgentString() string`

GetUserAgentString returns the UserAgentString field if non-nil, zero value otherwise.

### GetUserAgentStringOk

`func (o *UserAgentXMLDataArray) GetUserAgentStringOk() (*string, bool)`

GetUserAgentStringOk returns a tuple with the UserAgentString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentString

`func (o *UserAgentXMLDataArray) SetUserAgentString(v string)`

SetUserAgentString sets UserAgentString field to given value.

### HasUserAgentString

`func (o *UserAgentXMLDataArray) HasUserAgentString() bool`

HasUserAgentString returns a boolean if a field has been set.

### GetName

`func (o *UserAgentXMLDataArray) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserAgentXMLDataArray) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserAgentXMLDataArray) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserAgentXMLDataArray) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UserAgentXMLDataArray) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserAgentXMLDataArray) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserAgentXMLDataArray) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserAgentXMLDataArray) HasType() bool`

HasType returns a boolean if a field has been set.

### GetVersion

`func (o *UserAgentXMLDataArray) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UserAgentXMLDataArray) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UserAgentXMLDataArray) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UserAgentXMLDataArray) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVersionMajor

`func (o *UserAgentXMLDataArray) GetVersionMajor() string`

GetVersionMajor returns the VersionMajor field if non-nil, zero value otherwise.

### GetVersionMajorOk

`func (o *UserAgentXMLDataArray) GetVersionMajorOk() (*string, bool)`

GetVersionMajorOk returns a tuple with the VersionMajor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionMajor

`func (o *UserAgentXMLDataArray) SetVersionMajor(v string)`

SetVersionMajor sets VersionMajor field to given value.

### HasVersionMajor

`func (o *UserAgentXMLDataArray) HasVersionMajor() bool`

HasVersionMajor returns a boolean if a field has been set.

### GetDevice

`func (o *UserAgentXMLDataArray) GetDevice() UserAgentDataDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *UserAgentXMLDataArray) GetDeviceOk() (*UserAgentDataDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *UserAgentXMLDataArray) SetDevice(v UserAgentDataDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *UserAgentXMLDataArray) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetEngine

`func (o *UserAgentXMLDataArray) GetEngine() UserAgentDataEngine`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *UserAgentXMLDataArray) GetEngineOk() (*UserAgentDataEngine, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *UserAgentXMLDataArray) SetEngine(v UserAgentDataEngine)`

SetEngine sets Engine field to given value.

### HasEngine

`func (o *UserAgentXMLDataArray) HasEngine() bool`

HasEngine returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *UserAgentXMLDataArray) GetOperatingSystem() UserAgentDataOperatingSystem`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *UserAgentXMLDataArray) GetOperatingSystemOk() (*UserAgentDataOperatingSystem, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *UserAgentXMLDataArray) SetOperatingSystem(v UserAgentDataOperatingSystem)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *UserAgentXMLDataArray) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


