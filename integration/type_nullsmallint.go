// SPDX-FileCopyrightText: 2023 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "gen_type NullSmallInt github.com/ncotds/go-dblib/asetypes.NullInt16 -columndef smallint null"; DO NOT EDIT.

package integration

import (
	"database/sql"

	"github.com/ncotds/go-dblib/asetypes"

	"testing"
)

// DoTestNullSmallInt tests the handling of the NullSmallInt.
func DoTestNullSmallInt(t *testing.T) {
	TestForEachDB("TestNullSmallInt", t, testNullSmallInt)
	//
}

func testNullSmallInt(t *testing.T, db *sql.DB, tableName string) {
	// insert is the amount of insertions (see fn SetupTableInsert)
	insert := 2

	pass := make([]interface{}, len(samplesNullSmallInt))
	mySamples := make([]asetypes.NullInt16, len(samplesNullSmallInt)*insert)

	for i, sample := range samplesNullSmallInt {

		mySample := sample

		pass[i] = mySample

		// Add passed sample for the later validation (for every
		// insert)
		for j := 0; j < insert; j++ {
			mySamples[i+(len(samplesNullSmallInt)*j)] = mySample
		}
	}

	rows, teardownFn, err := SetupTableInsert(db, tableName, "smallint null", pass...)
	if err != nil {
		t.Errorf("Error preparing table: %v", err)
		return
	}
	defer rows.Close()
	defer teardownFn()

	i := 0
	var recv asetypes.NullInt16
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
