// SPDX-FileCopyrightText: 2020 SAP SE
// SPDX-FileCopyrightText: 2021 SAP SE
// SPDX-FileCopyrightText: 2022 SAP SE
// SPDX-FileCopyrightText: 2023 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

package tds

import "fmt"

var _ Package = (*ControlPackage)(nil)

type ControlPackage struct {
	// Reference the previous RowFmt
	rowFmt        *RowFmtPackage
	ColumnControl []byte
}

// LastPkg implements the tds.LastPkgAcceptor interface.
func (pkg *ControlPackage) LastPkg(other Package) error {
	if rowFmt, ok := other.(*RowFmtPackage); ok {
		pkg.rowFmt = rowFmt
		return nil
	}
	return fmt.Errorf("received package other than RowFmtPackage: %T", other)
}

// ReadFrom implements the tds.Package interface.
func (pkg *ControlPackage) ReadFrom(ch BytesChannel) error {
	length, err := ch.Uint16()
	if err != nil {
		return ErrNotEnoughBytes
	}
	pkg.ColumnControl, err = ch.Bytes(int(length))
	if err != nil {
		return ErrNotEnoughBytes
	}
	return nil
}

// WriteTo implements the tds.Package interface.
func (pkg ControlPackage) WriteTo(ch BytesChannel) error {
	var err error
	err = ch.WriteByte(byte(TDS_CONTROL))
	if err != nil {
		return fmt.Errorf("error occurred writing TDS Token %s: %w", TDS_CONTROL, err)
	}
	err = ch.WriteUint16(uint16(len(pkg.ColumnControl)))
	if err != nil {
		return fmt.Errorf("error occurred writing package length: %w", err)
	}
	err = ch.WriteBytes(pkg.ColumnControl)
	if err != nil {
		return fmt.Errorf("error occurred writing column bytes: %w", err)
	}
	return nil
}

func (pkg ControlPackage) String() string {
	s := fmt.Sprintf("%T(%d): %v", pkg, len(pkg.ColumnControl), pkg.ColumnControl)
	return s
}
