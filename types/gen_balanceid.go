// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/denkhaus/bitshares/util"
	"github.com/denkhaus/logging"
	"github.com/juju/errors"
)

type BalanceID struct {
	ObjectID
}

func (p BalanceID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *BalanceID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeBalance) << 48) | instance)
	return nil
}

type BalanceIDs []BalanceID

func (p BalanceIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode BalanceID")
		}
	}

	return nil
}

func BalanceIDFromObject(ob GrapheneObject) BalanceID {
	id, ok := ob.(*BalanceID)
	if ok {
		return *id
	}

	p := BalanceID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeBalance {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeBalance'", p.ID()))
	}

	return p
}

//NewBalanceID creates an new BalanceID object
func NewBalanceID(id string) GrapheneObject {
	gid := new(BalanceID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"BalanceID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeBalance {
		logging.Errorf(
			"BalanceID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeBalance'", id),
		)
		return nil
	}

	return gid
}
