package accountevent

import (
	"github.com/function61/pi-security-module/util/eventbase"
	"log"
)

const (
	SecretUsedTypeSshSigning      = "SshSigning"
	SecretUsedTypePasswordExposed = "PasswordExposed"
)

type SecretUsed struct {
	eventbase.Event
	Account string
	Type    string
}

func (e *SecretUsed) Apply() {
	log.Printf("Account %s was used, type = %s", e.Account, e.Type)
}