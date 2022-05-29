package validator

import (
	"net/mail"
	"neural_storage/cube/core/entities/user"
	"time"
	"unicode"

	"github.com/google/uuid"
)

func (v *Validator) validateId(id string) bool {
	_, err := uuid.Parse(id)

	return err == nil
}

func (v *Validator) validateEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	return err == nil
}

func (v *Validator) validateUsername(username string) bool {
	if len(username) < v.conf.MinUnameLen() {
		return false
	}
	if len(username) > v.conf.MaxUnameLen() {
		return false
	}

	isAlphaNumercialASCII := func(value []rune) bool {
		for i := 0; i < len(value); i++ {
			if value[i] > unicode.MaxASCII {
				return false
			}

			if !unicode.IsDigit(value[i]) && !unicode.IsLetter(value[i]) {
				return false
			}
		}
		return true
	}

	return isAlphaNumercialASCII([]rune(username))
}

func (v *Validator) validatePassword(password string) bool {
	if len(password) < v.conf.MinPwdLen() {
		return false
	}
	if len(password) > v.conf.MaxPwdLen() {
		return false
	}

	isPrintableASCII := func(value []rune) bool {
		for i := 0; i < len(value); i++ {
			if value[i] > unicode.MaxASCII || !unicode.IsPrint(value[i]) {
				return false
			}
		}
		return true
	}

	return isPrintableASCII([]rune(password))
}

func (v *Validator) validateBlockTime(t time.Time) bool {
	return t.After(time.Now())
}

func (v *Validator) ValidateUserInfo(info *user.Info) bool {
	if info.ID() != "" && !v.validateId(info.ID()) {
		return false
	}
	if info.Email() != "" && !v.validateEmail(info.Email()) {
		return false
	}
	if info.Username() != "" && !v.validateUsername(info.Username()) {
		return false
	}
	if info.Pwd() != "" && !v.validatePassword(info.Pwd()) {
		return false
	}
	if !info.BlockedUntil().IsZero() && !v.validateBlockTime(info.BlockedUntil()) {
		return false
	}

	return true
}
