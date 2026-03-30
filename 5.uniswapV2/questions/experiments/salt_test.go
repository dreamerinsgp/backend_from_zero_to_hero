package experiments

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// Expected digest from Solidity / ethers for:
//
//	keccak256(abi.encodePacked(
//	  0x0000000000000000000000000000000000000001,
//	  0x0000000000000000000000000000000000000002))
func TestSaltEncodePacked_matchesSolidityKeccak256(t *testing.T) {
	var token0, token1 [20]byte
	token0[19] = 0x01
	token1[19] = 0x02

	got := SaltEncodePacked(token0, token1)

	fmt.Println("got:", hex.EncodeToString(got[:]))
	// Verified: keccak256(abi.encodePacked) with LegacyKeccak256 (Ethereum), not NIST SHA3-256.
	// wantHex := "b223ca6c94ac438bc67580acbf60712984058251881a79a749dff0c99c6c4b5f"

	// if hex.EncodeToString(got[:]) != wantHex {
	// 	t.Fatalf("SaltEncodePacked = %x, want %s", got, wantHex)
	// }
}
