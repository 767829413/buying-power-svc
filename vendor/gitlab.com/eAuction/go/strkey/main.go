package strkey

import (
	"bytes"
	"encoding/base32"
	"encoding/binary"
	"fmt"

	"github.com/pkg/errors"
	"gitlab.com/eAuction/go/crc16"
)

// ErrInvalidVersionByte is returned when the version byte from a provided
// strkey-encoded string is not one of the valid values.
var ErrInvalidVersionByte = errors.New("invalid version byte")

// VersionByte represents one of the possible prefix values for a StrKey base
// string--the string the when encoded using base32 yields a final StrKey.
type VersionByte byte

const (
	//VersionByteAccountID is the version byte used for encoded stellar addresses
	VersionByteAccountID VersionByte = 6 << 3 // Base32-encodes to 'G...'
	//VersionByteSeed is the version byte used for encoded stellar seed
	VersionByteSeed      = 18 << 3 // Base32-encodes to 'S...'
	VersionByteBalanceID = 1 << 3  // Base32-encodes to 'B...'
)

// Decode decodes the provided StrKey into a raw value, checking the checksum
// and ensuring the expected VersionByte (the version parameter) is the value
// actually encoded into the provided src string.
func Decode(expected VersionByte, src string) ([]byte, error) {
	if err := checkValidVersionByte(expected); err != nil {
		return nil, err
	}

	raw, err := base32.StdEncoding.DecodeString(src)
	if err != nil {
		return nil, err
	}

	if len(raw) < 3 {
		return nil, fmt.Errorf("encoded value is %d bytes; minimum valid length is 3", len(raw))
	}

	// decode into components
	version := VersionByte(raw[0])
	vp := raw[0 : len(raw)-2]
	payload := raw[1 : len(raw)-2]
	checksum := raw[len(raw)-2:]

	// ensure version byte is expected
	if version != expected {
		return nil, ErrInvalidVersionByte
	}

	// ensure checksum is valid
	if err := crc16.Validate(vp, checksum); err != nil {
		return nil, err
	}

	// if we made it through the gaunlet, return the decoded value
	return payload, nil
}

// MustDecode is like Decode, but panics on error
func MustDecode(expected VersionByte, src string) []byte {
	d, err := Decode(expected, src)
	if err != nil {
		panic(err)
	}
	return d
}

// Encode encodes the provided data to a StrKey, using the provided version
// byte.
func Encode(version VersionByte, src []byte) (string, error) {
	if err := checkValidVersionByte(version); err != nil {
		return "", err
	}

	var raw bytes.Buffer

	// write version byte
	if err := binary.Write(&raw, binary.LittleEndian, version); err != nil {
		return "", err
	}

	// write payload
	if _, err := raw.Write(src); err != nil {
		return "", err
	}

	// calculate and write checksum
	checksum := crc16.Checksum(raw.Bytes())
	if _, err := raw.Write(checksum); err != nil {
		return "", err
	}

	result := base32.StdEncoding.EncodeToString(raw.Bytes())
	return result, nil
}

// MustEncode is like Encode, but panics on error
func MustEncode(version VersionByte, src []byte) string {
	e, err := Encode(version, src)
	if err != nil {
		panic(err)
	}
	return e
}

// checkValidVersionByte returns an error if the provided value
// is not one of the defined valid version byte constants.
func checkValidVersionByte(version VersionByte) error {
	if version == VersionByteAccountID {
		return nil
	}
	if version == VersionByteSeed {
		return nil
	}

	if version == VersionByteBalanceID {
		return nil
	}

	return ErrInvalidVersionByte
}
