// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/sample-service/sample-service.proto

package sample_service

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

// Validate checks the field values on Template with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Template) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	return nil
}

// TemplateValidationError is the validation error returned by
// Template.Validate if the designated constraints aren't met.
type TemplateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TemplateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TemplateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TemplateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TemplateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TemplateValidationError) ErrorName() string { return "TemplateValidationError" }

// Error satisfies the builtin error interface
func (e TemplateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTemplate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TemplateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TemplateValidationError{}

// Validate checks the field values on SampleMethodV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SampleMethodV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetId() <= 0 {
		return SampleMethodV1RequestValidationError{
			field:  "Id",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// SampleMethodV1RequestValidationError is the validation error returned by
// SampleMethodV1Request.Validate if the designated constraints aren't met.
type SampleMethodV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SampleMethodV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SampleMethodV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SampleMethodV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SampleMethodV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SampleMethodV1RequestValidationError) ErrorName() string {
	return "SampleMethodV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e SampleMethodV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSampleMethodV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SampleMethodV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SampleMethodV1RequestValidationError{}

// Validate checks the field values on SampleMethodV1Response with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *SampleMethodV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SampleMethodV1ResponseValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// SampleMethodV1ResponseValidationError is the validation error returned by
// SampleMethodV1Response.Validate if the designated constraints aren't met.
type SampleMethodV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SampleMethodV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SampleMethodV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SampleMethodV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SampleMethodV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SampleMethodV1ResponseValidationError) ErrorName() string {
	return "SampleMethodV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SampleMethodV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSampleMethodV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SampleMethodV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SampleMethodV1ResponseValidationError{}
