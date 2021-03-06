// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user/v1/register.proto

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

// Validate checks the field values on RegisterAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterAccountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterAccountRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterAccountRequestMultiError, or nil if none found.
func (m *RegisterAccountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterAccountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUsername()) < 1 {
		err := RegisterAccountRequestValidationError{
			field:  "Username",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		err := RegisterAccountRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetNickname()) < 1 {
		err := RegisterAccountRequestValidationError{
			field:  "Nickname",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterAccountRequestMultiError(errors)
	}

	return nil
}

// RegisterAccountRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterAccountRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterAccountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterAccountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterAccountRequestMultiError) AllErrors() []error { return m }

// RegisterAccountRequestValidationError is the validation error returned by
// RegisterAccountRequest.Validate if the designated constraints aren't met.
type RegisterAccountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterAccountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterAccountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterAccountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterAccountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterAccountRequestValidationError) ErrorName() string {
	return "RegisterAccountRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterAccountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterAccountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterAccountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterAccountRequestValidationError{}

// Validate checks the field values on RegisterReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReplyMultiError, or
// nil if none found.
func (m *RegisterReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return RegisterReplyMultiError(errors)
	}

	return nil
}

// RegisterReplyMultiError is an error wrapping multiple validation errors
// returned by RegisterReply.ValidateAll() if the designated constraints
// aren't met.
type RegisterReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReplyMultiError) AllErrors() []error { return m }

// RegisterReplyValidationError is the validation error returned by
// RegisterReply.Validate if the designated constraints aren't met.
type RegisterReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReplyValidationError) ErrorName() string { return "RegisterReplyValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReplyValidationError{}

// Validate checks the field values on RegisterPreEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterPreEmailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterPreEmailRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterPreEmailRequestMultiError, or nil if none found.
func (m *RegisterPreEmailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterPreEmailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetEmail()) < 1 {
		err := RegisterPreEmailRequestValidationError{
			field:  "Email",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterPreEmailRequestMultiError(errors)
	}

	return nil
}

// RegisterPreEmailRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterPreEmailRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterPreEmailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterPreEmailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterPreEmailRequestMultiError) AllErrors() []error { return m }

// RegisterPreEmailRequestValidationError is the validation error returned by
// RegisterPreEmailRequest.Validate if the designated constraints aren't met.
type RegisterPreEmailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterPreEmailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterPreEmailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterPreEmailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterPreEmailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterPreEmailRequestValidationError) ErrorName() string {
	return "RegisterPreEmailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterPreEmailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterPreEmailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterPreEmailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterPreEmailRequestValidationError{}

// Validate checks the field values on RegisterEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterEmailRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterEmailRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterEmailRequestMultiError, or nil if none found.
func (m *RegisterEmailRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterEmailRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateEmail(m.GetEmail()); err != nil {
		err = RegisterEmailRequestValidationError{
			field:  "Email",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		err := RegisterEmailRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetNickname()) < 1 {
		err := RegisterEmailRequestValidationError{
			field:  "Nickname",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCode()) < 1 {
		err := RegisterEmailRequestValidationError{
			field:  "Code",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterEmailRequestMultiError(errors)
	}

	return nil
}

func (m *RegisterEmailRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *RegisterEmailRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// RegisterEmailRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterEmailRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterEmailRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterEmailRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterEmailRequestMultiError) AllErrors() []error { return m }

// RegisterEmailRequestValidationError is the validation error returned by
// RegisterEmailRequest.Validate if the designated constraints aren't met.
type RegisterEmailRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterEmailRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterEmailRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterEmailRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterEmailRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterEmailRequestValidationError) ErrorName() string {
	return "RegisterEmailRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterEmailRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterEmailRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterEmailRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterEmailRequestValidationError{}

// Validate checks the field values on RegisterPrePhoneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterPrePhoneRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterPrePhoneRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterPrePhoneRequestMultiError, or nil if none found.
func (m *RegisterPrePhoneRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterPrePhoneRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if !_RegisterPrePhoneRequest_Phone_Pattern.MatchString(m.GetPhone()) {
		err := RegisterPrePhoneRequestValidationError{
			field:  "Phone",
			reason: "value does not match regex pattern \"^\\\\+[1-9]\\\\d{1,14}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterPrePhoneRequestMultiError(errors)
	}

	return nil
}

// RegisterPrePhoneRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterPrePhoneRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterPrePhoneRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterPrePhoneRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterPrePhoneRequestMultiError) AllErrors() []error { return m }

// RegisterPrePhoneRequestValidationError is the validation error returned by
// RegisterPrePhoneRequest.Validate if the designated constraints aren't met.
type RegisterPrePhoneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterPrePhoneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterPrePhoneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterPrePhoneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterPrePhoneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterPrePhoneRequestValidationError) ErrorName() string {
	return "RegisterPrePhoneRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterPrePhoneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterPrePhoneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterPrePhoneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterPrePhoneRequestValidationError{}

var _RegisterPrePhoneRequest_Phone_Pattern = regexp.MustCompile("^\\+[1-9]\\d{1,14}$")

// Validate checks the field values on RegisterPhoneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegisterPhoneRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterPhoneRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterPhoneRequestMultiError, or nil if none found.
func (m *RegisterPhoneRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterPhoneRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetPhone()) < 1 {
		err := RegisterPhoneRequestValidationError{
			field:  "Phone",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPassword()) < 1 {
		err := RegisterPhoneRequestValidationError{
			field:  "Password",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetNickname()) < 1 {
		err := RegisterPhoneRequestValidationError{
			field:  "Nickname",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetCode()) < 1 {
		err := RegisterPhoneRequestValidationError{
			field:  "Code",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return RegisterPhoneRequestMultiError(errors)
	}

	return nil
}

// RegisterPhoneRequestMultiError is an error wrapping multiple validation
// errors returned by RegisterPhoneRequest.ValidateAll() if the designated
// constraints aren't met.
type RegisterPhoneRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterPhoneRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterPhoneRequestMultiError) AllErrors() []error { return m }

// RegisterPhoneRequestValidationError is the validation error returned by
// RegisterPhoneRequest.Validate if the designated constraints aren't met.
type RegisterPhoneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterPhoneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterPhoneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterPhoneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterPhoneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterPhoneRequestValidationError) ErrorName() string {
	return "RegisterPhoneRequestValidationError"
}

// Error satisfies the builtin error interface
func (e RegisterPhoneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterPhoneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterPhoneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterPhoneRequestValidationError{}
