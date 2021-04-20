// Code generated by "stringer -type=SignatureAlgorithm"; DO NOT EDIT.

package sema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SignatureAlgorithmUnknown-0]
	_ = x[SignatureAlgorithmECDSA_P256-1]
	_ = x[SignatureAlgorithmECDSA_secp256k1-2]
	_ = x[SignatureAlgorithmBLS_BLS12381-3]
}

const _SignatureAlgorithm_name = "SignatureAlgorithmUnknownSignatureAlgorithmECDSA_P256SignatureAlgorithmECDSA_secp256k1SignatureAlgorithmBLS_BLS12381"

var _SignatureAlgorithm_index = [...]uint8{0, 25, 53, 86, 116}

func (i SignatureAlgorithm) String() string {
	if i >= SignatureAlgorithm(len(_SignatureAlgorithm_index)-1) {
		return "SignatureAlgorithm(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SignatureAlgorithm_name[_SignatureAlgorithm_index[i]:_SignatureAlgorithm_index[i+1]]
}
