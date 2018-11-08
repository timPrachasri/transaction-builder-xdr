package transactionbuilder_test

import (
	"encoding/base64"
	"strings"
	builder "transaction-builder-xdr/transaction/builder"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stellar/go/xdr"
	xdrBuilder "gitlab.com/lightnet-thailand/xdr-builder"
)

var _ = Describe("Creating transaction XDR with payment operation", func() {
	var (
		opB64 string
	)

	BeforeEach(func() {
		asset, err := xdrBuilder.SetAsset("ABC", "GAEBJVQJJO5ZPRJ2ZPNSDJLMNN64REZO7S5VUZAMNLI34B5XUQVD3URR")
		if err != nil {
			panic(err)
		}
		limit := "1000"
		opB64, err = xdrBuilder.ChangeTrust(SourceAddr, asset, limit)
		Expect(err).NotTo(HaveOccurred())
	})

	It("should return a correct xdr transaction string", func() {
		By("adding one change trust operation")
		var (
			tB64 string
		)
		tx := xdr.Transaction{
			SourceAccount: Source,
			Fee:           10,
			SeqNum:        xdr.SequenceNumber(1),
			Memo:          Memo,
		}
		transactionBuilder := builder.GetInstance(&tx)
		transactionBuilder.MakeOperation(opB64)
		tB64, err = transactionBuilder.ToBase64()
		expected := "AAAAAFsAPNHwcy2ZPYftEEoI+dAPr0ZBN+vuXUKPEKcq2mmtAAAACgAAAAAAAAABAAAAAAAAAAAAAAABAAAAAQAAAABbADzR8HMtmT2H7RBKCPnQD69GQTfr7l1CjxCnKtpprQAAAAYAAAABQUJDAAAAAAAIFNYJS7uXxTrL2yGlbGt9yJMu/LtaZAxq0b4Ht6QqPQAAAAJUC+QAAAAAAA=="
		Expect(err).NotTo(HaveOccurred())
		Expect(tB64).Should(Equal(expected))
	})

	It("should return a correct unmarshalled bytes and operation", func() {
		By("adding one change trust operation")
		var (
			tB64           string
			unmarshalledTx xdr.Transaction
			bytesRead      int
		)
		tx := xdr.Transaction{
			SourceAccount: Source,
			Fee:           10,
			SeqNum:        xdr.SequenceNumber(1),
			Memo:          Memo,
		}
		transactionBuilder := builder.GetInstance(&tx)
		transactionBuilder.MakeOperation(opB64)
		tB64, err = transactionBuilder.ToBase64()
		rawr := strings.NewReader(tB64)
		b64r := base64.NewDecoder(base64.StdEncoding, rawr)
		bytesRead, err = xdr.Unmarshal(b64r, &unmarshalledTx)
		Expect(bytesRead).Should(Equal(160))
		Expect(len(unmarshalledTx.Operations)).Should(Equal(1))
	})
})
