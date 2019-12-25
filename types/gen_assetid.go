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

type AssetID struct {
	ObjectID
}

func (p AssetID) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(p.Instance())); err != nil {
		return errors.Annotate(err, "encode instance")
	}

	return nil
}

func (p *AssetID) Unmarshal(dec *util.TypeDecoder) error {
	var instance uint64
	if err := dec.DecodeUVarint(&instance); err != nil {
		return errors.Annotate(err, "decode instance")
	}

	p.number = UInt64((uint64(SpaceTypeProtocol) << 56) | (uint64(ObjectTypeAsset) << 48) | instance)
	return nil
}

type AssetIDs []AssetID

func (p AssetIDs) Marshal(enc *util.TypeEncoder) error {
	if err := enc.EncodeUVarint(uint64(len(p))); err != nil {
		return errors.Annotate(err, "encode length")
	}

	for _, ex := range p {
		if err := enc.Encode(ex); err != nil {
			return errors.Annotate(err, "encode AssetID")
		}
	}

	return nil
}

func AssetIDFromObject(ob GrapheneObject) AssetID {
	id, ok := ob.(*AssetID)
	if ok {
		return *id
	}

	p := AssetID{}
	p.MustFromObject(ob)
	if p.ObjectType() != ObjectTypeAsset {
		panic(fmt.Sprintf("invalid ObjectType: %q has no ObjectType 'ObjectTypeAsset'", p.ID()))
	}

	return p
}

//NewAssetID creates an new AssetID object
func NewAssetID(id string) GrapheneObject {
	gid := new(AssetID)
	if err := gid.Parse(id); err != nil {
		logging.Errorf(
			"AssetID parser error %v",
			errors.Annotate(err, "Parse"),
		)
		return nil
	}

	if gid.ObjectType() != ObjectTypeAsset {
		logging.Errorf(
			"AssetID parser error %s",
			fmt.Sprintf("%q has no ObjectType 'ObjectTypeAsset'", id),
		)
		return nil
	}

	return gid
}
