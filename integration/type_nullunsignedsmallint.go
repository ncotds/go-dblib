// SPDX-FileCopyrightText: 2023 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "gen_type NullUnsignedSmallInt github.com/ncotds/go-dblib/asetypes.NullUint16 -columndef unsigned smallint null"; DO NOT EDIT.

package integration

import (
	"database/sql"

	"github.com/ncotds/go-dblib/asetypes"

	"testing"
)

// DoTestNullUnsignedSmallInt tests the handling of the NullUnsignedSmallInt.
func DoTestNullUnsignedSmallInt(t *testing.T) {
	TestForEachDB("TestNullUnsignedSmallInt", t, testNullUnsignedSmallInt)
	//
}

func testNullUnsignedSmallInt(t *testing.T, db *sql.DB, tableName string) {
	// insert is the amount of insertions (see fn SetupTableInsert)
	insert := 2

	pass := make([]interface{}, len(samplesNullUnsignedSmallInt))
	mySamples := make([]asetypes.NullUint16, len(samplesNullUnsignedSmallInt)*insert)

	for i, sample := range samplesNullUnsignedSmallInt {

		mySample := sample

		pass[i] = mySample

		// Add passed sample for the later validation (for every
		// insert)
		for j := 0; j < insert; j++ {
			mySamples[i+(len(samplesNullUnsignedSmallInt)*j)] = mySample
		}
	}

	rows, teardownFn, err := SetupTableInsert(db, tableName, "unsigned smallint null", pass...)
	if err != nil {
		t.Errorf("Error preparing table: %v", err)
		return
	}
	defer rows.Close()
	defer teardownFn()

	i := 0
	var recv asetypes.NullUint16
	for rows.Next() {
		if err := rows.Scan(&recv); err != nil {
			t.Errorf("Scan failed on %dth scan: %v", i, err)
			continue
		}

		if recv != mySamples[i] {

			t.Errorf("Received value does not match passed parameter")
			t.Errorf("Expected: %v", mySamples[i])
			t.Errorf("Received: %v", recv)
		}

		i++
	}

	if err := rows.Err(); err != nil {
		t.Errorf("Error preparing rows: %v", err)
	}

	if i != len(pass)*insert {
		t.Errorf("Only read %d values from database, expected to read %d", i, len(pass)*insert)
	}
}
