package transaction

import (
	"testing"

	"github.com/elastic/apm-server/tests/loader"
)

func BenchmarkWithFileLoading(b *testing.B) {
	processor := NewProcessor(nil)
	for i := 0; i < b.N; i++ {
		data, _ := loader.LoadValidData("transaction")
		err := processor.Validate(data)
		if err != nil {
			b.Fatalf("Error: %v", err)
		}
		processor.Transform(data)
	}
}

func BenchmarkTransactionFileLoadingOnce(b *testing.B) {
	processor := NewProcessor(nil)
	data, _ := loader.LoadValidData("transaction")
	for i := 0; i < b.N; i++ {
		err := processor.Validate(data)
		if err != nil {
			b.Fatalf("Error: %v", err)
		}
		processor.Transform(data)
	}
}
