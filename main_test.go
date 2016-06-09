package main_test

import (
	"flag"
	"log"
	"os"
	"testing"
)

func Test(t *testing.T) {
	expected := 0
	actual := 0
	if expected != actual {
		t.Errorf("%v is not %v", actual, expected)
	}
}

func SetupTest() (err error) {
	return
}

func TearDownTest() (err error) {
	return
}

func TestMain(m *testing.M) {
	var (
		err error
		res int
	)
	defer func() {
		if err != nil {
			log.Fatal(err)
		}
	}()

	flag.Parse()

	err = SetupTest()
	if err != nil {
		return
	}

	res = m.Run()
	err = TearDownTest()
	if err != nil {
		return
	}

	os.Exit(res)
}

func BenchmarkMain(b *testing.B) {
}
