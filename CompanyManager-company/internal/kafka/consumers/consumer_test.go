package consumers

//import (
//	_ "github.com/lib/pq"
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//	"testing"
//)

//func TestCart(t *testing.T) {
//	RegisterFailHandler(Fail)
//	RunSpecs(t, "Consumer Suite")
//}
//
//
//var _ = Describe("Consumer", func() {
//	Describe("StringToInt64", func() {
//		It("Should return an error and zero if can't do Atoi conversion", func() {
//			msg1 := make(chan []byte)
//			KafkaConsumer("WrongTopic", "wrong broker", msg1)
//			Ω().Should(Panic())
//		})
//
//		//It("Should return an int64 value and nil error if Atoi conversion works", func() {
//		//	got, err := StringToInt64("2")
//		//	Expect(err).NotTo(HaveOccurred())
//		//	Expect(got).To(Equal(int64(2)))
//		//})
//	})
//
//
//})