/*
Formance Stack API

Open, modular foundation for unique payments flows  # Introduction This API is documented in **OpenAPI format**.  # Authentication Formance Stack offers one forms of authentication:   - OAuth2 OAuth2 - an open protocol to allow secure authorization in a simple and standard method from web, mobile and desktop applications. <SecurityDefinitions /> 

API version: v1.0.0-rc.2
Contact: support@formance.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package formance

import (
	"encoding/json"
	"fmt"
)

// ListConnectorTasks200ResponseInner - struct for ListConnectorTasks200ResponseInner
type ListConnectorTasks200ResponseInner struct {
	TaskDescriptorBankingCircle *TaskDescriptorBankingCircle
	TaskDescriptorCurrencyCloud *TaskDescriptorCurrencyCloud
	TaskDescriptorDummyPay *TaskDescriptorDummyPay
	TaskDescriptorModulr *TaskDescriptorModulr
	TaskDescriptorStripe *TaskDescriptorStripe
	TaskDescriptorWise *TaskDescriptorWise
}

// TaskDescriptorBankingCircleAsListConnectorTasks200ResponseInner is a convenience function that returns TaskDescriptorBankingCircle wrapped in ListConnectorTasks200ResponseInner
func TaskDescriptorBankingCircleAsListConnectorTasks200ResponseInner(v *TaskDescriptorBankingCircle) ListConnectorTasks200ResponseInner {
	return ListConnectorTasks200ResponseInner{
		TaskDescriptorBankingCircle: v,
	}
}

// TaskDescriptorCurrencyCloudAsListConnectorTasks200ResponseInner is a convenience function that returns TaskDescriptorCurrencyCloud wrapped in ListConnectorTasks200ResponseInner
func TaskDescriptorCurrencyCloudAsListConnectorTasks200ResponseInner(v *TaskDescriptorCurrencyCloud) ListConnectorTasks200ResponseInner {
	return ListConnectorTasks200ResponseInner{
		TaskDescriptorCurrencyCloud: v,
	}
}

// TaskDescriptorDummyPayAsListConnectorTasks200ResponseInner is a convenience function that returns TaskDescriptorDummyPay wrapped in ListConnectorTasks200ResponseInner
func TaskDescriptorDummyPayAsListConnectorTasks200ResponseInner(v *TaskDescriptorDummyPay) ListConnectorTasks200ResponseInner {
	return ListConnectorTasks200ResponseInner{
		TaskDescriptorDummyPay: v,
	}
}

// TaskDescriptorModulrAsListConnectorTasks200ResponseInner is a convenience function that returns TaskDescriptorModulr wrapped in ListConnectorTasks200ResponseInner
func TaskDescriptorModulrAsListConnectorTasks200ResponseInner(v *TaskDescriptorModulr) ListConnectorTasks200ResponseInner {
	return ListConnectorTasks200ResponseInner{
		TaskDescriptorModulr: v,
	}
}

// TaskDescriptorStripeAsListConnectorTasks200ResponseInner is a convenience function that returns TaskDescriptorStripe wrapped in ListConnectorTasks200ResponseInner
func TaskDescriptorStripeAsListConnectorTasks200ResponseInner(v *TaskDescriptorStripe) ListConnectorTasks200ResponseInner {
	return ListConnectorTasks200ResponseInner{
		TaskDescriptorStripe: v,
	}
}

// TaskDescriptorWiseAsListConnectorTasks200ResponseInner is a convenience function that returns TaskDescriptorWise wrapped in ListConnectorTasks200ResponseInner
func TaskDescriptorWiseAsListConnectorTasks200ResponseInner(v *TaskDescriptorWise) ListConnectorTasks200ResponseInner {
	return ListConnectorTasks200ResponseInner{
		TaskDescriptorWise: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ListConnectorTasks200ResponseInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into TaskDescriptorBankingCircle
	err = newStrictDecoder(data).Decode(&dst.TaskDescriptorBankingCircle)
	if err == nil {
		jsonTaskDescriptorBankingCircle, _ := json.Marshal(dst.TaskDescriptorBankingCircle)
		if string(jsonTaskDescriptorBankingCircle) == "{}" { // empty struct
			dst.TaskDescriptorBankingCircle = nil
		} else {
			match++
		}
	} else {
		dst.TaskDescriptorBankingCircle = nil
	}

	// try to unmarshal data into TaskDescriptorCurrencyCloud
	err = newStrictDecoder(data).Decode(&dst.TaskDescriptorCurrencyCloud)
	if err == nil {
		jsonTaskDescriptorCurrencyCloud, _ := json.Marshal(dst.TaskDescriptorCurrencyCloud)
		if string(jsonTaskDescriptorCurrencyCloud) == "{}" { // empty struct
			dst.TaskDescriptorCurrencyCloud = nil
		} else {
			match++
		}
	} else {
		dst.TaskDescriptorCurrencyCloud = nil
	}

	// try to unmarshal data into TaskDescriptorDummyPay
	err = newStrictDecoder(data).Decode(&dst.TaskDescriptorDummyPay)
	if err == nil {
		jsonTaskDescriptorDummyPay, _ := json.Marshal(dst.TaskDescriptorDummyPay)
		if string(jsonTaskDescriptorDummyPay) == "{}" { // empty struct
			dst.TaskDescriptorDummyPay = nil
		} else {
			match++
		}
	} else {
		dst.TaskDescriptorDummyPay = nil
	}

	// try to unmarshal data into TaskDescriptorModulr
	err = newStrictDecoder(data).Decode(&dst.TaskDescriptorModulr)
	if err == nil {
		jsonTaskDescriptorModulr, _ := json.Marshal(dst.TaskDescriptorModulr)
		if string(jsonTaskDescriptorModulr) == "{}" { // empty struct
			dst.TaskDescriptorModulr = nil
		} else {
			match++
		}
	} else {
		dst.TaskDescriptorModulr = nil
	}

	// try to unmarshal data into TaskDescriptorStripe
	err = newStrictDecoder(data).Decode(&dst.TaskDescriptorStripe)
	if err == nil {
		jsonTaskDescriptorStripe, _ := json.Marshal(dst.TaskDescriptorStripe)
		if string(jsonTaskDescriptorStripe) == "{}" { // empty struct
			dst.TaskDescriptorStripe = nil
		} else {
			match++
		}
	} else {
		dst.TaskDescriptorStripe = nil
	}

	// try to unmarshal data into TaskDescriptorWise
	err = newStrictDecoder(data).Decode(&dst.TaskDescriptorWise)
	if err == nil {
		jsonTaskDescriptorWise, _ := json.Marshal(dst.TaskDescriptorWise)
		if string(jsonTaskDescriptorWise) == "{}" { // empty struct
			dst.TaskDescriptorWise = nil
		} else {
			match++
		}
	} else {
		dst.TaskDescriptorWise = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.TaskDescriptorBankingCircle = nil
		dst.TaskDescriptorCurrencyCloud = nil
		dst.TaskDescriptorDummyPay = nil
		dst.TaskDescriptorModulr = nil
		dst.TaskDescriptorStripe = nil
		dst.TaskDescriptorWise = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ListConnectorTasks200ResponseInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ListConnectorTasks200ResponseInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ListConnectorTasks200ResponseInner) MarshalJSON() ([]byte, error) {
	if src.TaskDescriptorBankingCircle != nil {
		return json.Marshal(&src.TaskDescriptorBankingCircle)
	}

	if src.TaskDescriptorCurrencyCloud != nil {
		return json.Marshal(&src.TaskDescriptorCurrencyCloud)
	}

	if src.TaskDescriptorDummyPay != nil {
		return json.Marshal(&src.TaskDescriptorDummyPay)
	}

	if src.TaskDescriptorModulr != nil {
		return json.Marshal(&src.TaskDescriptorModulr)
	}

	if src.TaskDescriptorStripe != nil {
		return json.Marshal(&src.TaskDescriptorStripe)
	}

	if src.TaskDescriptorWise != nil {
		return json.Marshal(&src.TaskDescriptorWise)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ListConnectorTasks200ResponseInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.TaskDescriptorBankingCircle != nil {
		return obj.TaskDescriptorBankingCircle
	}

	if obj.TaskDescriptorCurrencyCloud != nil {
		return obj.TaskDescriptorCurrencyCloud
	}

	if obj.TaskDescriptorDummyPay != nil {
		return obj.TaskDescriptorDummyPay
	}

	if obj.TaskDescriptorModulr != nil {
		return obj.TaskDescriptorModulr
	}

	if obj.TaskDescriptorStripe != nil {
		return obj.TaskDescriptorStripe
	}

	if obj.TaskDescriptorWise != nil {
		return obj.TaskDescriptorWise
	}

	// all schemas are nil
	return nil
}

type NullableListConnectorTasks200ResponseInner struct {
	value *ListConnectorTasks200ResponseInner
	isSet bool
}

func (v NullableListConnectorTasks200ResponseInner) Get() *ListConnectorTasks200ResponseInner {
	return v.value
}

func (v *NullableListConnectorTasks200ResponseInner) Set(val *ListConnectorTasks200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorTasks200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorTasks200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorTasks200ResponseInner(val *ListConnectorTasks200ResponseInner) *NullableListConnectorTasks200ResponseInner {
	return &NullableListConnectorTasks200ResponseInner{value: val, isSet: true}
}

func (v NullableListConnectorTasks200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorTasks200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


