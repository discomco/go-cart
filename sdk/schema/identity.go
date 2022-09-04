package schema

import (
	"fmt"
	"github.com/discomco/go-cart/sdk/core/constants"
	"github.com/pkg/errors"
	"regexp"
	"strings"

	"github.com/discomco/go-cart/sdk/test"
	"github.com/google/uuid"
)

type IIdentity interface {
	Id() string
}

func ImplementsIIdentity(id IIdentity) bool {
	return true
}

type Identity struct {
	Prefix string `json:"prefix"`
	Value  string `json:"value"`
}

func (i *Identity) Id() string {
	return fmt.Sprintf("%+v-%+v", i.Prefix, i.Value)
}

func checkPrefix(prefix string) (string, error) {
	if prefix == "" {
		prefix = test.DEFAULT_PREFIX
	}
	regex := regexp.MustCompile("^[a-z]+$")
	match := regex.Match([]byte(prefix))
	if !match {
		err := fmt.Errorf("Prefix cmd_must only contain characters a-z (lowercase)")
		return "", err
	}
	return prefix, nil
}

func checkValue(value string) (string, error) {
	if value == "" {
		value = test.CLEAN_NULL_UUID
	}
	value = strings.Replace(value, "-", "", -1)
	value = strings.ToLower(value)
	regex := regexp.MustCompile("^[a-z\\d]+$")
	match := regex.Match([]byte(value))
	if !match {
		err := fmt.Errorf("Value cmd_must only contain characters 0-9 and a-z (lowercase)")
		err = errors.Wrap(err, "(checkValue).Regex")
		return "", err
	}
	return value, nil
}

func NilIdentity() (*Identity, error) {
	return NewIdentityFrom(constants.Nil, uuid.Nil.String())
}

func NewIdentityFrom(prefix string, id string) (*Identity, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		err = errors.Wrapf(err, "(NewIdentityFrom.UUID")
		return nil, err
	}
	prefix, err = checkPrefix(prefix)
	if err != nil {
		err = errors.Wrap(err, "(NewIdentityFrom.checkPrefix")
		return nil, err
	}
	value, err := checkValue(uid.String())
	if err != nil {
		err = errors.Wrap(err, "(NewIdentityFrom.checkValue")
		return nil, err
	}
	return &Identity{Prefix: prefix, Value: value}, err
}

func NewIdentity(prefix string) (*Identity, error) {
	prefix, err := checkPrefix(prefix)
	if err != nil {
		err = errors.Wrap(err, "(NewIdentity.checkPrefix")
		return nil, err
	}
	val, err := CleanUuid()
	if err != nil {
		err = errors.Wrap(err, "(NewIdentity.CleanUuid")
		return nil, err
	}
	return &Identity{Prefix: prefix, Value: val}, err
}

func IdentityFromPrefixedId(prefixedId string) (*Identity, error) {
	parts := strings.Split(prefixedId, "-")
	prefix, err := checkPrefix(parts[0])
	if err != nil {
		err = errors.Wrap(err, "(IdentityFromPrefixedId.checkPrefix")
		return nil, err
	}
	value, err := checkValue(parts[1])
	if err != nil {
		err = errors.Wrap(err, "(IdentityFromPrefixedId.checkValue")
		return nil, err
	}
	return &Identity{
		Prefix: prefix,
		Value:  value,
	}, err
}
