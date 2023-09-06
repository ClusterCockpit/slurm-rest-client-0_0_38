/*
Slurm Rest API

API to access and control Slurm.

API version: 0.0.38
Contact: sales@schedmd.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the V0038MetaPlugin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038MetaPlugin{}

// V0038MetaPlugin struct for V0038MetaPlugin
type V0038MetaPlugin struct {
	// 
	Type *string `json:"type,omitempty"`
	// 
	Name *string `json:"name,omitempty"`
}

// NewV0038MetaPlugin instantiates a new V0038MetaPlugin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038MetaPlugin() *V0038MetaPlugin {
	this := V0038MetaPlugin{}
	return &this
}

// NewV0038MetaPluginWithDefaults instantiates a new V0038MetaPlugin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038MetaPluginWithDefaults() *V0038MetaPlugin {
	this := V0038MetaPlugin{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V0038MetaPlugin) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038MetaPlugin) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V0038MetaPlugin) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V0038MetaPlugin) SetType(v string) {
	o.Type = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V0038MetaPlugin) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038MetaPlugin) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V0038MetaPlugin) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V0038MetaPlugin) SetName(v string) {
	o.Name = &v
}

func (o V0038MetaPlugin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038MetaPlugin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableV0038MetaPlugin struct {
	value *V0038MetaPlugin
	isSet bool
}

func (v NullableV0038MetaPlugin) Get() *V0038MetaPlugin {
	return v.value
}

func (v *NullableV0038MetaPlugin) Set(val *V0038MetaPlugin) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038MetaPlugin) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038MetaPlugin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038MetaPlugin(val *V0038MetaPlugin) *NullableV0038MetaPlugin {
	return &NullableV0038MetaPlugin{value: val, isSet: true}
}

func (v NullableV0038MetaPlugin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038MetaPlugin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


