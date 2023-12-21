// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/formancehq/formance-sdk-go/pkg/utils"
)

type StageType string

const (
	StageTypeStageSend      StageType = "StageSend"
	StageTypeStageDelay     StageType = "StageDelay"
	StageTypeStageWaitEvent StageType = "StageWaitEvent"
)

type Stage struct {
	StageSend      *StageSend
	StageDelay     *StageDelay
	StageWaitEvent *StageWaitEvent

	Type StageType
}

func CreateStageStageSend(stageSend StageSend) Stage {
	typ := StageTypeStageSend

	return Stage{
		StageSend: &stageSend,
		Type:      typ,
	}
}

func CreateStageStageDelay(stageDelay StageDelay) Stage {
	typ := StageTypeStageDelay

	return Stage{
		StageDelay: &stageDelay,
		Type:       typ,
	}
}

func CreateStageStageWaitEvent(stageWaitEvent StageWaitEvent) Stage {
	typ := StageTypeStageWaitEvent

	return Stage{
		StageWaitEvent: &stageWaitEvent,
		Type:           typ,
	}
}

func (u *Stage) UnmarshalJSON(data []byte) error {

	stageWaitEvent := StageWaitEvent{}
	if err := utils.UnmarshalJSON(data, &stageWaitEvent, "", true, true); err == nil {
		u.StageWaitEvent = &stageWaitEvent
		u.Type = StageTypeStageWaitEvent
		return nil
	}

	stageDelay := StageDelay{}
	if err := utils.UnmarshalJSON(data, &stageDelay, "", true, true); err == nil {
		u.StageDelay = &stageDelay
		u.Type = StageTypeStageDelay
		return nil
	}

	stageSend := StageSend{}
	if err := utils.UnmarshalJSON(data, &stageSend, "", true, true); err == nil {
		u.StageSend = &stageSend
		u.Type = StageTypeStageSend
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Stage) MarshalJSON() ([]byte, error) {
	if u.StageSend != nil {
		return utils.MarshalJSON(u.StageSend, "", true)
	}

	if u.StageDelay != nil {
		return utils.MarshalJSON(u.StageDelay, "", true)
	}

	if u.StageWaitEvent != nil {
		return utils.MarshalJSON(u.StageWaitEvent, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
