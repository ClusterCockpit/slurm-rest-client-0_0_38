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

// checks if the V0038ReservationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V0038ReservationsResponse{}

// V0038ReservationsResponse struct for V0038ReservationsResponse
type V0038ReservationsResponse struct {
	Meta *V0038Meta `json:"meta,omitempty"`
	// slurm errors
	Errors []V0038Error `json:"errors,omitempty"`
	// reservation info
	Reservations []V0038Reservation `json:"reservations,omitempty"`
}

// NewV0038ReservationsResponse instantiates a new V0038ReservationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV0038ReservationsResponse() *V0038ReservationsResponse {
	this := V0038ReservationsResponse{}
	return &this
}

// NewV0038ReservationsResponseWithDefaults instantiates a new V0038ReservationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV0038ReservationsResponseWithDefaults() *V0038ReservationsResponse {
	this := V0038ReservationsResponse{}
	return &this
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *V0038ReservationsResponse) GetMeta() V0038Meta {
	if o == nil || IsNil(o.Meta) {
		var ret V0038Meta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038ReservationsResponse) GetMetaOk() (*V0038Meta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *V0038ReservationsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given V0038Meta and assigns it to the Meta field.
func (o *V0038ReservationsResponse) SetMeta(v V0038Meta) {
	o.Meta = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *V0038ReservationsResponse) GetErrors() []V0038Error {
	if o == nil || IsNil(o.Errors) {
		var ret []V0038Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038ReservationsResponse) GetErrorsOk() ([]V0038Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *V0038ReservationsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []V0038Error and assigns it to the Errors field.
func (o *V0038ReservationsResponse) SetErrors(v []V0038Error) {
	o.Errors = v
}

// GetReservations returns the Reservations field value if set, zero value otherwise.
func (o *V0038ReservationsResponse) GetReservations() []V0038Reservation {
	if o == nil || IsNil(o.Reservations) {
		var ret []V0038Reservation
		return ret
	}
	return o.Reservations
}

// GetReservationsOk returns a tuple with the Reservations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V0038ReservationsResponse) GetReservationsOk() ([]V0038Reservation, bool) {
	if o == nil || IsNil(o.Reservations) {
		return nil, false
	}
	return o.Reservations, true
}

// HasReservations returns a boolean if a field has been set.
func (o *V0038ReservationsResponse) HasReservations() bool {
	if o != nil && !IsNil(o.Reservations) {
		return true
	}

	return false
}

// SetReservations gets a reference to the given []V0038Reservation and assigns it to the Reservations field.
func (o *V0038ReservationsResponse) SetReservations(v []V0038Reservation) {
	o.Reservations = v
}

func (o V0038ReservationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V0038ReservationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.Reservations) {
		toSerialize["reservations"] = o.Reservations
	}
	return toSerialize, nil
}

type NullableV0038ReservationsResponse struct {
	value *V0038ReservationsResponse
	isSet bool
}

func (v NullableV0038ReservationsResponse) Get() *V0038ReservationsResponse {
	return v.value
}

func (v *NullableV0038ReservationsResponse) Set(val *V0038ReservationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV0038ReservationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV0038ReservationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV0038ReservationsResponse(val *V0038ReservationsResponse) *NullableV0038ReservationsResponse {
	return &NullableV0038ReservationsResponse{value: val, isSet: true}
}

func (v NullableV0038ReservationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV0038ReservationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


