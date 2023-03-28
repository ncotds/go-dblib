// SPDX-FileCopyrightText: 2023 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "stringer -type=CursorOStatus"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_CUR_OSTAT_UNUSED-0]
	_ = x[TDS_CUR_OSTAT_HASARGS-1]
	_ = x[TDS_CUR_CONSEC_UPDS-2]
}

const _CursorOStatus_name = "TDS_CUR_OSTAT_UNUSEDTDS_CUR_OSTAT_HASARGSTDS_CUR_CONSEC_UPDS"

var _CursorOStatus_index = [...]uint8{0, 20, 41, 60}

func (i CursorOStatus) String() string {
	if i >= CursorOStatus(len(_CursorOStatus_index)-1) {
		return "CursorOStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CursorOStatus_name[_CursorOStatus_index[i]:_CursorOStatus_index[i+1]]
}
