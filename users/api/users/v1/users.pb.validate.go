// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/users/v1/users.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on CreateUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	if l := utf8.RuneCountInString(m.GetAccount()); l < 1 || l > 20 {
		return CreateUsersRequestValidationError{
			field:  "Account",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
	}

	// no validation rules for Phone

	// no validation rules for Name

	if utf8.RuneCountInString(m.GetPassword()) != 32 {
		return CreateUsersRequestValidationError{
			field:  "Password",
			reason: "value length must be 32 runes",
		}

	}

	// no validation rules for Age

	// no validation rules for Sex

	// no validation rules for Email

	// no validation rules for School

	// no validation rules for Address

	// no validation rules for CardId

	// no validation rules for CardName

	return nil
}

// CreateUsersRequestValidationError is the validation error returned by
// CreateUsersRequest.Validate if the designated constraints aren't met.
type CreateUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUsersRequestValidationError) ErrorName() string {
	return "CreateUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUsersRequestValidationError{}

// Validate checks the field values on CreateUsersReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateUsersReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	// no validation rules for Message

	return nil
}

// CreateUsersReplyValidationError is the validation error returned by
// CreateUsersReply.Validate if the designated constraints aren't met.
type CreateUsersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUsersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUsersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUsersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUsersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUsersReplyValidationError) ErrorName() string { return "CreateUsersReplyValidationError" }

// Error satisfies the builtin error interface
func (e CreateUsersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUsersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUsersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUsersReplyValidationError{}

// Validate checks the field values on UpdateUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateUsersRequestValidationError is the validation error returned by
// UpdateUsersRequest.Validate if the designated constraints aren't met.
type UpdateUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUsersRequestValidationError) ErrorName() string {
	return "UpdateUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUsersRequestValidationError{}

// Validate checks the field values on UpdateUsersReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateUsersReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateUsersReplyValidationError is the validation error returned by
// UpdateUsersReply.Validate if the designated constraints aren't met.
type UpdateUsersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUsersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUsersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUsersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUsersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUsersReplyValidationError) ErrorName() string { return "UpdateUsersReplyValidationError" }

// Error satisfies the builtin error interface
func (e UpdateUsersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUsersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUsersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUsersReplyValidationError{}

// Validate checks the field values on DeleteUsersRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Uuid

	return nil
}

// DeleteUsersRequestValidationError is the validation error returned by
// DeleteUsersRequest.Validate if the designated constraints aren't met.
type DeleteUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUsersRequestValidationError) ErrorName() string {
	return "DeleteUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUsersRequestValidationError{}

// Validate checks the field values on DeleteUsersReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *DeleteUsersReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	// no validation rules for Msg

	return nil
}

// DeleteUsersReplyValidationError is the validation error returned by
// DeleteUsersReply.Validate if the designated constraints aren't met.
type DeleteUsersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteUsersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteUsersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteUsersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteUsersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteUsersReplyValidationError) ErrorName() string { return "DeleteUsersReplyValidationError" }

// Error satisfies the builtin error interface
func (e DeleteUsersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteUsersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteUsersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteUsersReplyValidationError{}

// Validate checks the field values on GetUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetUuid()) < 1 {
		return GetUsersRequestValidationError{
			field:  "Uuid",
			reason: "value length must be at least 1 runes",
		}
	}

	return nil
}

// GetUsersRequestValidationError is the validation error returned by
// GetUsersRequest.Validate if the designated constraints aren't met.
type GetUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersRequestValidationError) ErrorName() string { return "GetUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersRequestValidationError{}

// Validate checks the field values on GetUsersReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUsersReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	// no validation rules for Message

	if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUsersReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// GetUsersReplyValidationError is the validation error returned by
// GetUsersReply.Validate if the designated constraints aren't met.
type GetUsersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersReplyValidationError) ErrorName() string { return "GetUsersReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUsersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersReplyValidationError{}

// Validate checks the field values on GetUsersData with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GetUsersData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Age

	// no validation rules for Phone

	// no validation rules for Sex

	// no validation rules for Account

	// no validation rules for Email

	// no validation rules for Address

	// no validation rules for Avatar

	// no validation rules for Status

	// no validation rules for Uid

	// no validation rules for School

	return nil
}

// GetUsersDataValidationError is the validation error returned by
// GetUsersData.Validate if the designated constraints aren't met.
type GetUsersDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUsersDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUsersDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUsersDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUsersDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUsersDataValidationError) ErrorName() string { return "GetUsersDataValidationError" }

// Error satisfies the builtin error interface
func (e GetUsersDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUsersData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUsersDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUsersDataValidationError{}

// Validate checks the field values on PatchUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PatchUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Age

	// no validation rules for Address

	// no validation rules for Email

	// no validation rules for Avatar

	// no validation rules for School

	// no validation rules for Phone

	// no validation rules for Sex

	// no validation rules for Uuid

	return nil
}

// PatchUsersRequestValidationError is the validation error returned by
// PatchUsersRequest.Validate if the designated constraints aren't met.
type PatchUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchUsersRequestValidationError) ErrorName() string {
	return "PatchUsersRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PatchUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchUsersRequestValidationError{}

// Validate checks the field values on PatchUsersReply with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PatchUsersReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	// no validation rules for Message

	return nil
}

// PatchUsersReplyValidationError is the validation error returned by
// PatchUsersReply.Validate if the designated constraints aren't met.
type PatchUsersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchUsersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchUsersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchUsersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchUsersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchUsersReplyValidationError) ErrorName() string { return "PatchUsersReplyValidationError" }

// Error satisfies the builtin error interface
func (e PatchUsersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchUsersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchUsersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchUsersReplyValidationError{}

// Validate checks the field values on ListUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListUsersRequestValidationError is the validation error returned by
// ListUsersRequest.Validate if the designated constraints aren't met.
type ListUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersRequestValidationError) ErrorName() string { return "ListUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersRequestValidationError{}

// Validate checks the field values on ListUsersReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListUsersReply) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListUsersReplyValidationError is the validation error returned by
// ListUsersReply.Validate if the designated constraints aren't met.
type ListUsersReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersReplyValidationError) ErrorName() string { return "ListUsersReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListUsersReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersReplyValidationError{}

// Validate checks the field values on LoginUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *LoginUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Account

	// no validation rules for Password

	return nil
}

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

// Validate checks the field values on PatchPasswordRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PatchPasswordRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Account

	// no validation rules for Password

	// no validation rules for OldPassword

	return nil
}

// PatchPasswordRequestValidationError is the validation error returned by
// PatchPasswordRequest.Validate if the designated constraints aren't met.
type PatchPasswordRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchPasswordRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchPasswordRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchPasswordRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchPasswordRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchPasswordRequestValidationError) ErrorName() string {
	return "PatchPasswordRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PatchPasswordRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchPasswordRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchPasswordRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchPasswordRequestValidationError{}

// Validate checks the field values on PatchPasswordReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PatchPasswordReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Success

	// no validation rules for Message

	return nil
}

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

// Validate checks the field values on PatchUsersReplyData with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *PatchUsersReplyData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Age

	// no validation rules for Phone

	// no validation rules for Sex

	// no validation rules for Account

	// no validation rules for Email

	// no validation rules for Address

	// no validation rules for Avatar

	// no validation rules for Status

	// no validation rules for Uuid

	// no validation rules for School

	return nil
}

// PatchUsersReplyDataValidationError is the validation error returned by
// PatchUsersReplyData.Validate if the designated constraints aren't met.
type PatchUsersReplyDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PatchUsersReplyDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PatchUsersReplyDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PatchUsersReplyDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PatchUsersReplyDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PatchUsersReplyDataValidationError) ErrorName() string {
	return "PatchUsersReplyDataValidationError"
}

// Error satisfies the builtin error interface
func (e PatchUsersReplyDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPatchUsersReplyData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PatchUsersReplyDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PatchUsersReplyDataValidationError{}
