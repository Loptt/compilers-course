package duck

import (
	"github.com/Loptt/compilers-course/little_duck/gocc/lexer"
	"github.com/Loptt/compilers-course/little_duck/gocc/parser"
	"testing"
	"os"
	"fmt"
)


func readFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	fileinfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	bytesr, err := file.Read(buffer)
	fmt.Println(bytesr)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func TestDuck(t *testing.T) {
	//p := parser.NewParser()
	tests := []string {
		"test1.duck",
		"test2.duck",
	}

	for _, test := range tests {
		_, err := readFile(test)

		if err != nil {
			t.Fatalf("Error reading file %s", test);
		}

		//s := lexer.NewLexer(input);
		//_, _ := 1,2 // p.Parse(s);
/*
		if errtest != nil {
			t.Errorf("Error: %v", errtest);
		}*/
	}
}