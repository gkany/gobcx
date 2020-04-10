// Code generated by ffjson <https://github.com/pquerna/ffjson>. DO NOT EDIT.
// source: createnhassetoperation.go

package operations

import (
	"bytes"
	"fmt"
	"github.com/gkany/cocos-go/types"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

// MarshalJSON marshal bytes to json - template
func (j *CreateNhAssetOperation) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if j == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := j.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalJSONBuf marshal buff to json - template
func (j *CreateNhAssetOperation) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if j == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{ "fee_paying_account":`)

	{

		obj, err = j.FeePayingAccount.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"owner":`)

	{

		obj, err = j.Owner.MarshalJSON()
		if err != nil {
			return err
		}
		buf.Write(obj)

	}
	buf.WriteString(`,"asset_id":`)
	fflib.WriteJsonString(buf, string(j.AssetID))
	buf.WriteString(`,"world_view":`)
	fflib.WriteJsonString(buf, string(j.WorldView))
	buf.WriteString(`,"base_describe":`)
	fflib.WriteJsonString(buf, string(j.BaseDescribe))
	buf.WriteString(`,"collateral":`)
	fflib.FormatBits2(buf, uint64(j.Collateral), 10, j.Collateral < 0)
	buf.WriteByte(',')
	if j.Fee != nil {
		if true {
			buf.WriteString(`"fee":`)

			{

				err = j.Fee.MarshalJSONBuf(buf)
				if err != nil {
					return err
				}

			}
			buf.WriteByte(',')
		}
	}
	buf.Rewind(1)
	buf.WriteByte('}')
	return nil
}

const (
	ffjtCreateNhAssetOperationbase = iota
	ffjtCreateNhAssetOperationnosuchkey

	ffjtCreateNhAssetOperationFeePayingAccount

	ffjtCreateNhAssetOperationOwner

	ffjtCreateNhAssetOperationAssetID

	ffjtCreateNhAssetOperationWorldView

	ffjtCreateNhAssetOperationBaseDescribe

	ffjtCreateNhAssetOperationCollateral

	ffjtCreateNhAssetOperationFee
)

var ffjKeyCreateNhAssetOperationFeePayingAccount = []byte("fee_paying_account")

var ffjKeyCreateNhAssetOperationOwner = []byte("owner")

var ffjKeyCreateNhAssetOperationAssetID = []byte("asset_id")

var ffjKeyCreateNhAssetOperationWorldView = []byte("world_view")

var ffjKeyCreateNhAssetOperationBaseDescribe = []byte("base_describe")

var ffjKeyCreateNhAssetOperationCollateral = []byte("collateral")

var ffjKeyCreateNhAssetOperationFee = []byte("fee")

// UnmarshalJSON umarshall json - template of ffjson
func (j *CreateNhAssetOperation) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return j.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

// UnmarshalJSONFFLexer fast json unmarshall - template ffjson
func (j *CreateNhAssetOperation) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error
	currentKey := ffjtCreateNhAssetOperationbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffjtCreateNhAssetOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'a':

					if bytes.Equal(ffjKeyCreateNhAssetOperationAssetID, kn) {
						currentKey = ffjtCreateNhAssetOperationAssetID
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'b':

					if bytes.Equal(ffjKeyCreateNhAssetOperationBaseDescribe, kn) {
						currentKey = ffjtCreateNhAssetOperationBaseDescribe
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'c':

					if bytes.Equal(ffjKeyCreateNhAssetOperationCollateral, kn) {
						currentKey = ffjtCreateNhAssetOperationCollateral
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'f':

					if bytes.Equal(ffjKeyCreateNhAssetOperationFeePayingAccount, kn) {
						currentKey = ffjtCreateNhAssetOperationFeePayingAccount
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffjKeyCreateNhAssetOperationFee, kn) {
						currentKey = ffjtCreateNhAssetOperationFee
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'o':

					if bytes.Equal(ffjKeyCreateNhAssetOperationOwner, kn) {
						currentKey = ffjtCreateNhAssetOperationOwner
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'w':

					if bytes.Equal(ffjKeyCreateNhAssetOperationWorldView, kn) {
						currentKey = ffjtCreateNhAssetOperationWorldView
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.SimpleLetterEqualFold(ffjKeyCreateNhAssetOperationFee, kn) {
					currentKey = ffjtCreateNhAssetOperationFee
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCreateNhAssetOperationCollateral, kn) {
					currentKey = ffjtCreateNhAssetOperationCollateral
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCreateNhAssetOperationBaseDescribe, kn) {
					currentKey = ffjtCreateNhAssetOperationBaseDescribe
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCreateNhAssetOperationWorldView, kn) {
					currentKey = ffjtCreateNhAssetOperationWorldView
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffjKeyCreateNhAssetOperationAssetID, kn) {
					currentKey = ffjtCreateNhAssetOperationAssetID
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffjKeyCreateNhAssetOperationOwner, kn) {
					currentKey = ffjtCreateNhAssetOperationOwner
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.AsciiEqualFold(ffjKeyCreateNhAssetOperationFeePayingAccount, kn) {
					currentKey = ffjtCreateNhAssetOperationFeePayingAccount
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffjtCreateNhAssetOperationnosuchkey
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffjtCreateNhAssetOperationFeePayingAccount:
					goto handle_FeePayingAccount

				case ffjtCreateNhAssetOperationOwner:
					goto handle_Owner

				case ffjtCreateNhAssetOperationAssetID:
					goto handle_AssetID

				case ffjtCreateNhAssetOperationWorldView:
					goto handle_WorldView

				case ffjtCreateNhAssetOperationBaseDescribe:
					goto handle_BaseDescribe

				case ffjtCreateNhAssetOperationCollateral:
					goto handle_Collateral

				case ffjtCreateNhAssetOperationFee:
					goto handle_Fee

				case ffjtCreateNhAssetOperationnosuchkey:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_FeePayingAccount:

	/* handler: j.FeePayingAccount type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.FeePayingAccount.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Owner:

	/* handler: j.Owner type=types.AccountID kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Owner.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_AssetID:

	/* handler: j.AssetID type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.AssetID = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_WorldView:

	/* handler: j.WorldView type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.WorldView = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_BaseDescribe:

	/* handler: j.BaseDescribe type=string kind=string quoted=false*/

	{

		{
			if tok != fflib.FFTok_string && tok != fflib.FFTok_null {
				return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for string", tok))
			}
		}

		if tok == fflib.FFTok_null {

		} else {

			outBuf := fs.Output.Bytes()

			j.BaseDescribe = string(string(outBuf))

		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Collateral:

	/* handler: j.Collateral type=types.Int64 kind=int64 quoted=false*/

	{
		if tok == fflib.FFTok_null {

		} else {

			tbuf, err := fs.CaptureField(tok)
			if err != nil {
				return fs.WrapErr(err)
			}

			err = j.Collateral.UnmarshalJSON(tbuf)
			if err != nil {
				return fs.WrapErr(err)
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_Fee:

	/* handler: j.Fee type=types.AssetAmount kind=struct quoted=false*/

	{
		if tok == fflib.FFTok_null {

			j.Fee = nil

		} else {

			if j.Fee == nil {
				j.Fee = new(types.AssetAmount)
			}

			err = j.Fee.UnmarshalJSONFFLexer(fs, fflib.FFParse_want_key)
			if err != nil {
				return err
			}
		}
		state = fflib.FFParse_after_value
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:

	return nil
}
