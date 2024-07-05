// SPDX-FileCopyrightText: 2023 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "gen_type Money4 github.com/ncotds/go-dblib/*asetypes.Decimal -columndef smallmoney -convert convertSmallMoney -compare compareDecimal"; DO NOT EDIT.

package integration

import (
	"database/sql"

	"github.com/ncotds/go-dblib/asetypes"

	"testing"
)

// DoTestMoney4 tests the handling of the Money4.
func DoTestMoney4(t *testing.T) {
	TestForEachDB("TestMoney4", t, testMoney4)
	//
}

func testMoney4(t *testing.T, db *sql.DB, tableName string) {
	// insert is the amount of insertions (see fn SetupTableInsert)
	insert := 2

	pass := make([]interface{}, len(samplesMoney4))
	mySamples := make([]*asetypes.Decimal, len(samplesMoney4)*insert)

	for i, sample := range samplesMoney4 {

		// Convert sample with passed function before proceeding
		mySample, err := convertSmallMoney(sample)
		if err != nil {
			t.Errorf("Failed to convert sample %v: %v", sample, err)
			return
		}

		pass[i] = mySample

		// Add passed sample for the later validation (for every
		// insert)
		for j := 0; j < insert; j++ {
			mySamples[i+(len(samplesMoney4)*j)] = mySample
		}
	}

	rows, teardownFn, err := SetupTableInsert(db, tableName, "smallmoney", pass...)
	if err != nil {
		t.Errorf("Error preparing table: %v", err)
		return
	}
	defer rows.Close()
	defer teardownFn()

	i := 0
	var recv *asetypes.Decimal
	for rows.Next() {
		if err := rows.Scan(&recv); err != nil {
			t.Errorf("Scan failed on %dth scan: %v", i, err)
			continue
		}

		if compareDecimal(recv, mySamples[i]) {

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
