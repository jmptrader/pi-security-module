package accountevent

import (
	"github.com/function61/pi-security-module/state"
	"github.com/function61/pi-security-module/util/eventbase"
)

type SshKeyAdded struct {
	eventbase.Event
	Account                string
	Id                     string
	SshPrivateKey          string
	SshPublicKeyAuthorized string
}

func (e *SshKeyAdded) Apply() {
	for idx, account := range state.Inst.State.Accounts {
		if account.Id == e.Account {
			secret := state.Secret{
				Id:                     e.Id,
				Kind:                   state.SecretKindSshKey,
				Created:                e.Timestamp,
				SshPrivateKey:          e.SshPrivateKey,
				SshPublicKeyAuthorized: e.SshPublicKeyAuthorized,
			}

			account.Secrets = append(account.Secrets, secret)
			state.Inst.State.Accounts[idx] = account
			return
		}
	}
}
