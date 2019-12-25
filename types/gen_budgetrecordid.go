// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package types

import (
	"fmt"

	"github.com/gkany/graphSDK/logging"
	"github.com/gkany/graphSDK/util"
	"github.com/juju/errors"
)

type BudgetRecordID struct {
	ObjectID
}

func (p BudgetRecordID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *BudgetRecordID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeBudgetRecord) << 48) | instance)
	return nil
}

type BudgetRecordIDs []BudgetRecordID

func (p BudgetRecordIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode BudgetRecordID")
		}
	}

	return nil
}

func BudgetRecordIDFromObject(ob GrapheneObject) BudgetRecordID {
	id, ok := ob.(*BudgetRecordID)
	if ok {
		return *id
	}

	p := BudgetRecordID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeBudgetRecord {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeBudgetRecord'", p.ID()))
	}

	return p
}

//NewBudgetRecordID creates an new BudgetRecordID object
func NewBudgetRecordID(id string) GrapheneObject {
	gid := new(BudgetRecordID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"BudgetRecordID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeBudgetRecord {
		logging.Errorf(
			"BudgetRecordID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeBudgetRecord'", id),
		)
		return nil
	}

	return gid
}
