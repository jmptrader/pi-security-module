package accountevent

import (
	"encoding/json"
	"github.com/function61/pi-security-module/state"
	"github.com/function61/pi-security-module/util/eventbase"
)

type DescriptionChanged struct {
	eventbase.Event
	Id          string
	Description string
}

func (e DescriptionChanged) Serialize() string {
	asJson, _ := json.Marshal(e)

	return "DescriptionChanged " + string(asJson)
}

func (e DescriptionChanged) Apply() {
	for idx, s := range state.Inst.State.Accounts {
		if s.Id == e.Id {
			s.Description = e.Description
			state.Inst.State.Accounts[idx] = s
			return
		}
	}
}
