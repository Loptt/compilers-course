package duck

import (
	"github.com/Loptt/compilers-course/little_duck/gocc/lexer"
	"github.com/Loptt/compilers-course/little_duck/gocc/parser"
	"testing"
	"os"
)


func readFile(file string) []byte, error {
	file, err := os.Open("filetoread.txt")
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

	bytesread, err := file.Read(buffer)

	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func TestDuck(t *testing.T) {
	p := parser.NewParser()
	tests := []string {
		"test1.duck",
		"test2.duck",
	}

	for _, test := range tests {
		input, err := readFile(test)

		if err != nil {
			t.Fatalf("Error reading file %s", test);
		}

		s := lexer.NewLexer(input);
		_, errtest := p.Parse(s);

		if errtest != nil {
			t.Errorf("Error: %v", errtest);
		}
	}
}