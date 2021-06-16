package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Adi")
	require.Equal(t, "Hello Adi", result, "Reslut must be 'Hello Adi'")
	fmt.Println("TestHelloWorld with Require Done")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Adi")
	assert.Equal(t, "Hello Adi", result, "Reslut must be 'Hello Adi'")
	fmt.Println("TestHelloWorld with Assert Done")
}

func TestHelloWorldAdi(t *testing.T) {
	result := HelloWorld("Adi")

	if result != "Hello Adi" {
		// error
		t.Error("Result must be 'Hello Adi'")
	//	t.Fail()
	}

	fmt.Println("TestHelloWorldAdi Done")	
}

func TestHelloWorldPramesta(t *testing.T) {
	result := HelloWorld("Pramesta")

	if result != "Hello Pramesta" {
		// error
		t.Fatal("Result must be 'Hello Pramesta")
	//	t.FailNow()
	}

	fmt.Println("TestHelloWorldPramesta Done")
}	