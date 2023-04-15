package readfile

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

/* Write a function that reads from a file and captures its output and prints it to the screen */
func ReadFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		return errors.New("Could not open file: " + err.Error())
	}

	defer func() {
		if cerr := file.Close(); err != nil {
			fmt.Println("Failed to close file: " + cerr.Error())
		}
	}()

	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return errors.New("Could not read file: " + err.Error())
	}

	log.Printf("%s", bytes)
	return nil
}

/* Write a function that reads a file using the bufio package to handle large files or streaming from an input */
func ReadLarge(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {

		/* We can provide more context to the error by wrapping it with fmt.Errorf and the %w format verb */
		if os.IsNotExist(err) {
			return fmt.Errorf("File does not exist: %w", err)
		} else if os.IsPermission(err) {
			return fmt.Errorf("Permissions error: %w", err)
		} else {
			return err
		}
	}
	if err != nil {
		return errors.New("Could not open file: " + err.Error())
	}

	defer file.Close()

	newReader := bufio.NewReader(file)

	for {
		delim := byte('\n')
		s, err := newReader.ReadSlice(delim)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return errors.New("Could not read file: " + err.Error())
			}
		} else {
			fmt.Printf("%s", s)
		}
	}

	return nil
}
