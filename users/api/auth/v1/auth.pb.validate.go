// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/auth/v1/auth.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on LoginUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginUsersRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginUsersRequestMultiError, or nil if none found.
func (m *LoginUsersRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginUsersRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetAccount()) < 2 {
		err := LoginUsersRequestValidationError{
			field:  "Account",
			reason: "value length must be at least 2 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginUsersRequestMultiError(errors)
	}

	return nil
}

// LoginUsersRequestMultiError is an error wrapping multiple validation errors
// returned by LoginUsersRequest.ValidateAll() if the designated constraints
// aren't met.
type LoginUsersRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginUsersRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginUsersRequestMultiError) AllErrors() []error { return m }

// LoginUsersRequestValidationError is the validation error returned by
// LoginUsersRequest.Validate if the designated constraints aren't met.
type LoginUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUsersRequestValidationError) ErrorName() string {
	return "LoginUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LoginUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUsersRequestValidationError{}

// Validate checks the field values on LoginUsersReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginUsersReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginUsersReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginUsersReplyMultiError, or nil if none found.
func (m *LoginUsersReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginUsersReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Message

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginUsersReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginUsersReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginUsersReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return LoginUsersReplyMultiError(errors)
	}

	return nil
}

// LoginUsersReplyMultiError is an error wrapping multiple validation errors
// returned by LoginUsersReply.ValidateAll() if the designated constraints
// aren't met.
type LoginUsersReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginUsersReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginUsersReplyMultiError) AllErrors() []error { return m }

// LoginUsersReplyValidationError is the validation error returned by
// LoginUsersReply.Validate if the designated constraints aren't met.
type LoginUsersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUsersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUsersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUsersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUsersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUsersReplyValidationError) ErrorName() string { return "LoginUsersReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginUsersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUsersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUsersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUsersReplyValidationError{}

// Validate checks the field values on LoginData with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginData with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginDataMultiError, or nil
// if none found.
func (m *LoginData) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	// no validation rules for RefreshToken

	// no validation rules for Uid

	if len(errors) > 0 {
		return LoginDataMultiError(errors)
	}

	return nil
}

// LoginDataMultiError is an error wrapping multiple validation errors returned
// by LoginData.ValidateAll() if the designated constraints aren't met.
type LoginDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginDataMultiError) AllErrors() []error { return m }

// LoginDataValidationError is the validation error returned by
// LoginData.Validate if the designated constraints aren't met.
type LoginDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginDataValidationError) ErrorName() string { return "LoginDataValidationError" }

// Error satisfies the builtin error interface
func (e LoginDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginDataValidationError{}

// Validate checks the field values on ExitUsersLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExitUsersLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExitUsersLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExitUsersLoginRequestMultiError, or nil if none found.
func (m *ExitUsersLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExitUsersLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	if len(errors) > 0 {
		return ExitUsersLoginRequestMultiError(errors)
	}

	return nil
}

// ExitUsersLoginRequestMultiError is an error wrapping multiple validation
// errors returned by ExitUsersLoginRequest.ValidateAll() if the designated
// constraints aren't met.
type ExitUsersLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExitUsersLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExitUsersLoginRequestMultiError) AllErrors() []error { return m }

// ExitUsersLoginRequestValidationError is the validation error returned by
// ExitUsersLoginRequest.Validate if the designated constraints aren't met.
type ExitUsersLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExitUsersLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExitUsersLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExitUsersLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExitUsersLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExitUsersLoginRequestValidationError) ErrorName() string {
	return "ExitUsersLoginRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExitUsersLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExitUsersLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExitUsersLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExitUsersLoginRequestValidationError{}

// Validate checks the field values on ExitUsersLoginReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ExitUsersLoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExitUsersLoginReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExitUsersLoginReplyMultiError, or nil if none found.
func (m *ExitUsersLoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ExitUsersLoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Message

	if len(errors) > 0 {
		return ExitUsersLoginReplyMultiError(errors)
	}

	return nil
}

// ExitUsersLoginReplyMultiError is an error wrapping multiple validation
// errors returned by ExitUsersLoginReply.ValidateAll() if the designated
// constraints aren't met.
type ExitUsersLoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExitUsersLoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExitUsersLoginReplyMultiError) AllErrors() []error { return m }

// ExitUsersLoginReplyValidationError is the validation error returned by
// ExitUsersLoginReply.Validate if the designated constraints aren't met.
type ExitUsersLoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExitUsersLoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExitUsersLoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExitUsersLoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExitUsersLoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExitUsersLoginReplyValidationError) ErrorName() string {
	return "ExitUsersLoginReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ExitUsersLoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExitUsersLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExitUsersLoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExitUsersLoginReplyValidationError{}

// Validate checks the field values on PatchUsersLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PatchUsersLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PatchUsersLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PatchUsersLoginRequestMultiError, or nil if none found.
func (m *PatchUsersLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PatchUsersLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for RefreshToken

	if len(errors) > 0 {
		return PatchUsersLoginRequestMultiError(errors)
	}

	return nil
}

// PatchUsersLoginRequestMultiError is an error wrapping multiple validation
// errors returned by PatchUsersLoginRequest.ValidateAll() if the designated
// constraints aren't met.
type PatchUsersLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PatchUsersLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PatchUsersLoginRequestMultiError) AllErrors() []error { return m }

// PatchUsersLoginRequestValidationError is the validation error returned by
// PatchUsersLoginRequest.Validate if the designated constraints aren't met.
type PatchUsersLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchUsersLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchUsersLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchUsersLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchUsersLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchUsersLoginRequestValidationError) ErrorName() string {
	return "PatchUsersLoginRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PatchUsersLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchUsersLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchUsersLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchUsersLoginRequestValidationError{}

// Validate checks the field values on PatchUsersLoginReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PatchUsersLoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PatchUsersLoginReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PatchUsersLoginReplyMultiError, or nil if none found.
func (m *PatchUsersLoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PatchUsersLoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Message

	if len(errors) > 0 {
		return PatchUsersLoginReplyMultiError(errors)
	}

	return nil
}

// PatchUsersLoginReplyMultiError is an error wrapping multiple validation
// errors returned by PatchUsersLoginReply.ValidateAll() if the designated
// constraints aren't met.
type PatchUsersLoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PatchUsersLoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PatchUsersLoginReplyMultiError) AllErrors() []error { return m }

// PatchUsersLoginReplyValidationError is the validation error returned by
// PatchUsersLoginReply.Validate if the designated constraints aren't met.
type PatchUsersLoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchUsersLoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchUsersLoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchUsersLoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchUsersLoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchUsersLoginReplyValidationError) ErrorName() string {
	return "PatchUsersLoginReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PatchUsersLoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchUsersLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchUsersLoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchUsersLoginReplyValidationError{}

// Validate checks the field values on PatchPasswordReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PatchPasswordReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PatchPasswordReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PatchPasswordReplyMultiError, or nil if none found.
func (m *PatchPasswordReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PatchPasswordReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Message

	if len(errors) > 0 {
		return PatchPasswordReplyMultiError(errors)
	}

	return nil
}

// PatchPasswordReplyMultiError is an error wrapping multiple validation errors
// returned by PatchPasswordReply.ValidateAll() if the designated constraints
// aren't met.
type PatchPasswordReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PatchPasswordReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PatchPasswordReplyMultiError) AllErrors() []error { return m }

// PatchPasswordReplyValidationError is the validation error returned by
// PatchPasswordReply.Validate if the designated constraints aren't met.
type PatchPasswordReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchPasswordReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchPasswordReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchPasswordReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchPasswordReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchPasswordReplyValidationError) ErrorName() string {
	return "PatchPasswordReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PatchPasswordReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchPasswordReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchPasswordReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchPasswordReplyValidationError{}

// Validate checks the field values on AuthLoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *AuthLoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthLoginRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthLoginRequestMultiError, or nil if none found.
func (m *AuthLoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthLoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return AuthLoginRequestMultiError(errors)
	}

	return nil
}

// AuthLoginRequestMultiError is an error wrapping multiple validation errors
// returned by AuthLoginRequest.ValidateAll() if the designated constraints
// aren't met.
type AuthLoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthLoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthLoginRequestMultiError) AllErrors() []error { return m }

// AuthLoginRequestValidationError is the validation error returned by
// AuthLoginRequest.Validate if the designated constraints aren't met.
type AuthLoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthLoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthLoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthLoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthLoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthLoginRequestValidationError) ErrorName() string { return "AuthLoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e AuthLoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthLoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthLoginRequestValidationError{}

// Validate checks the field values on AuthLoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AuthLoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthLoginReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AuthLoginReplyMultiError,
// or nil if none found.
func (m *AuthLoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthLoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Message

	if len(errors) > 0 {
		return AuthLoginReplyMultiError(errors)
	}

	return nil
}

// AuthLoginReplyMultiError is an error wrapping multiple validation errors
// returned by AuthLoginReply.ValidateAll() if the designated constraints
// aren't met.
type AuthLoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthLoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthLoginReplyMultiError) AllErrors() []error { return m }

// AuthLoginReplyValidationError is the validation error returned by
// AuthLoginReply.Validate if the designated constraints aren't met.
type AuthLoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthLoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthLoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthLoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthLoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthLoginReplyValidationError) ErrorName() string { return "AuthLoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e AuthLoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthLoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthLoginReplyValidationError{}

// Validate checks the field values on PatchUsersLoginReplyData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PatchUsersLoginReplyData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PatchUsersLoginReplyData with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PatchUsersLoginReplyDataMultiError, or nil if none found.
func (m *PatchUsersLoginReplyData) ValidateAll() error {
	return m.validate(true)
}

func (m *PatchUsersLoginReplyData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uid

	// no validation rules for Token

	if len(errors) > 0 {
		return PatchUsersLoginReplyDataMultiError(errors)
	}

	return nil
}

// PatchUsersLoginReplyDataMultiError is an error wrapping multiple validation
// errors returned by PatchUsersLoginReplyData.ValidateAll() if the designated
// constraints aren't met.
type PatchUsersLoginReplyDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PatchUsersLoginReplyDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PatchUsersLoginReplyDataMultiError) AllErrors() []error { return m }

// PatchUsersLoginReplyDataValidationError is the validation error returned by
// PatchUsersLoginReplyData.Validate if the designated constraints aren't met.
type PatchUsersLoginReplyDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchUsersLoginReplyDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchUsersLoginReplyDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchUsersLoginReplyDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchUsersLoginReplyDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchUsersLoginReplyDataValidationError) ErrorName() string {
	return "PatchUsersLoginReplyDataValidationError"
}

// Error satisfies the builtin error interface
func (e PatchUsersLoginReplyDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchUsersLoginReplyData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchUsersLoginReplyDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchUsersLoginReplyDataValidationError{}