// Code generated by "stringer -type=PathVar"; DO NOT EDIT.

package tcp_estats

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PATH_TABLE_NONRECOVDAEPISODES-0]
	_ = x[PATH_TABLE_SUMOCTETSREORDERED-1]
	_ = x[PATH_TABLE_NONRECOVDA-2]
	_ = x[PATH_TABLE_SAMPLERTT-3]
	_ = x[PATH_TABLE_MAXRTT-4]
	_ = x[PATH_TABLE_MINRTT-5]
	_ = x[PATH_TABLE_SUMRTT-6]
	_ = x[PATH_TABLE_COUNTRTT-7]
	_ = x[PATH_TABLE_MAXRTO-8]
	_ = x[PATH_TABLE_MINRTO-9]
	_ = x[PATH_TABLE_PTTL-10]
	_ = x[PATH_TABLE_PTOSIN-11]
	_ = x[PATH_TABLE_PRECONGSUMCWND-12]
	_ = x[PATH_TABLE_PRECONGSUMRTT-13]
	_ = x[PATH_TABLE_POSTCONGSUMRTT-14]
	_ = x[PATH_TABLE_POSTCONGCOUNTRTT-15]
	_ = x[PATH_TABLE_ECNSIGNALS-16]
	_ = x[PATH_TABLE_DUPACKEPISODES-17]
	_ = x[PATH_TABLE_DUPACKSOUT-18]
	_ = x[PATH_TABLE_CERCVD-19]
	_ = x[PATH_TABLE_ECESENT-20]
}

const _PathVar_name = "PATH_TABLE_NONRECOVDAEPISODESPATH_TABLE_SUMOCTETSREORDEREDPATH_TABLE_NONRECOVDAPATH_TABLE_SAMPLERTTPATH_TABLE_MAXRTTPATH_TABLE_MINRTTPATH_TABLE_SUMRTTPATH_TABLE_COUNTRTTPATH_TABLE_MAXRTOPATH_TABLE_MINRTOPATH_TABLE_PTTLPATH_TABLE_PTOSINPATH_TABLE_PRECONGSUMCWNDPATH_TABLE_PRECONGSUMRTTPATH_TABLE_POSTCONGSUMRTTPATH_TABLE_POSTCONGCOUNTRTTPATH_TABLE_ECNSIGNALSPATH_TABLE_DUPACKEPISODESPATH_TABLE_DUPACKSOUTPATH_TABLE_CERCVDPATH_TABLE_ECESENT"

var _PathVar_index = [...]uint16{0, 29, 58, 79, 99, 116, 133, 150, 169, 186, 203, 218, 235, 260, 284, 309, 336, 357, 382, 403, 420, 438}

func (i PathVar) String() string {
	if i >= PathVar(len(_PathVar_index)-1) {
		return "PathVar(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _PathVar_name[_PathVar_index[i]:_PathVar_index[i+1]]
}