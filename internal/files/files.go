package files

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"unicode"
)

type File struct {
	Bytes      int64
	LineCount  int64
	WordCount  int64
	CharCount  int64
}

func FileInfo(file *os.File) (File, error) {
	if file == nil {
		return File{}, fmt.Errorf("file is nil")
	}

	var bytes int64 = 0
	var lineCount int64 = 0
	var charCount int64 = 0
	var wordCount int64 = 0
	
	reader := bufio.NewReader(file)
	var errr error
	var inWord bool = false

	for {
		char, size, err := reader.ReadRune()

		if err != nil {
			if err == io.EOF {
				break
			}
			errr = err
			break
		}
		bytes = bytes + int64(size)
		charCount = charCount + 1

		if unicode.IsSpace(char){
			if inWord {
				wordCount = wordCount + 1
				inWord = false
			}

			if char == '\n' {
				lineCount = lineCount + 1
			}
			
		} else {
			inWord = true
		}
	}

	// If the file ends with a word, increment the word count
	if inWord {
		wordCount = wordCount + 1
	}

	fileInfo := File{
		Bytes: bytes,
		LineCount: lineCount,
		WordCount: wordCount,
		CharCount: charCount,
	}

	return fileInfo, errr
}