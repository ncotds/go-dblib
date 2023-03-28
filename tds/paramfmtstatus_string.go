// SPDX-FileCopyrightText: 2023 SAP SE
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by "stringer -type=ParamFmtStatus"; DO NOT EDIT.

package tds

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[TDS_PARAM_NOSTATUS-0]
	_ = x[TDS_PARAM_RETURN-1]
	_ = x[TDS_PARAM_COLUMNSTATUS-8]
	_ = x[TDS_PARAM_NULLALLOWED-32]
}

const (
	_ParamFmtStatus_name_0 = "TDS_PARAM_NOSTATUSTDS_PARAM_RETURN"
	_ParamFmtStatus_name_1 = "TDS_PARAM_COLUMNSTATUS"
	_ParamFmtStatus_name_2 = "TDS_PARAM_NULLALLOWED"
)

var (
	_ParamFmtStatus_index_0 = [...]uint8{0, 18, 34}
)

func (i ParamFmtStatus) String() string {
	switch {
	case i <= 1:
		return _ParamFmtStatus_name_0[_ParamFmtStatus_index_0[i]:_ParamFmtStatus_index_0[i+1]]
	case i == 8:
		return _ParamFmtStatus_name_1
	case i == 32:
		return _ParamFmtStatus_name_2
	default:
		return "ParamFmtStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
