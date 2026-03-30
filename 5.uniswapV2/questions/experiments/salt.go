package experiments

import (
	"golang.org/x/crypto/sha3"
)

// SaltEncodePacked returns keccak256(abi.encodePacked(token0, token1)) as in
// UniswapV2Factory: bytes32 salt = keccak256(abi.encodePacked(token0, token1));
// token0 and token1 must be 20-byte Ethereum addresses (big-endian, left-padded in each slot).
func SaltEncodePacked(token0, token1 [20]byte) [32]byte {
	packed := make([]byte, 0, 40)
	packed = append(packed, token0[:]...)
	packed = append(packed, token1[:]...)
	h := sha3.NewLegacyKeccak256()
	_, _ = h.Write(packed)
	var out [32]byte
	h.Sum(out[:0])
	return out
}
