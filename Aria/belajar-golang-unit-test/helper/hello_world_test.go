package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct{
		name string
		request string
		expected string
	}{
		{
			name: "Adi",
			request: "Adi",
			expected: "Hello Adi",
		},
		{
			name: "Pramesta",
			request: "Pramesta",
			expected: "Hello Pramesta",
		},
		{
			name: "Khannedy",
			request: "Khannedy",
			expected: "Hello Khannedy",
		},
		{
			name: "Budi",
			request: "Budi",
			expected: "Hello Budi",
		},
		{
			name: "Ali",
			request: "Ali",
			expected: "Hello Ali",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Adi", func(t *testing.T) {
		result := HelloWorld("Adi")
		require.Equal(t, "Hello Adi", result, "Result must be 'Hello Adi'")
	})
	t.Run("Pramesta", func(t *testing.T) {
		result := HelloWorld("Pramesta")
		require.Equal(t, "Hello Pramesta", result, "Result must be 'Hello Pramesta'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	// after
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Mac OS")
	}

	result := HelloWorld("Adi")
	require.Equal(t, "Hello Adi", result, "Reslut must be 'Hello Adi'")
}

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