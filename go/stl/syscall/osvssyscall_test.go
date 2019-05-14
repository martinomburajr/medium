package medium_syscall

import (
	"io/ioutil"
	"os"
	"testing"
)

var smallLength = getBytesInFile("testdata/small.txt")
var largeLength = getBytesInFile("testdata/large.txt")

func getBytesInFile(filepath string) int {
	//t.Helper()

	file, err := os.Open(filepath)
	if err != nil {
		//t.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		//t.Fatal(err)
	}
	return len(data)
}

func BenchmarkOSRead(b *testing.B) {
	for i := 0;  i < b.N; i++  {
		OSOpenRead("testdata/large.txt", largeLength)
	}
}

func BenchmarkSyscallOpenRead(b *testing.B) {
	for i := 0;  i < b.N; i++  {
		SyscallOpenRead("testdata/large.txt", largeLength)
	}
}

//func TestOSRead(t *testing.T) {
//	l := getBytesInFile(t, "testdata/small.txt")
//	log.Print(l)
//}