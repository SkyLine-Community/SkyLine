package SkyLine

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type FileCurrentWithinParserEnvironment struct {
	Filename      string
	FileLocation  string
	FileExtension string
	FileBasename  string
	IsDir         bool
}

func (File *FileCurrentWithinParserEnvironment) New(filename string)   { File.Filename = filename }  // Method | Assigns new file
func (File *FileCurrentWithinParserEnvironment) Get_Name() string      { return File.Filename }      // Method | Returns file name
func (File *FileCurrentWithinParserEnvironment) Get_Extension() string { return File.FileExtension } // Method | Returns file extension
func (File *FileCurrentWithinParserEnvironment) Get_Basename() string  { return File.FileBasename }  // Method | Returns file basename

// Verifies that the file does exist
func (File *FileCurrentWithinParserEnvironment) VerifyFileExists(logging bool) bool {
	if File.Filename == "" {
		fmt.Println(ErrorSymBolMap[CODE_FILE_INTEGRITY_FILE_INVALID_FILE_NAME_WAS_EMPTY_OR_NULL_CHEC_INPUT]())
		os.Exit(0)
	}
	f, x := os.Stat(File.Filename)
	defer func() {
		if fe := recover(); fe != nil {
			if fe == "invalid memory address or nil pointer dereference	" {
				fmt.Println("File error: File was either \n  ---> 1): Empty \n ---> 2): A filepath which is not supported \n ---> 3): An invalid file format \n MAKE SURE YOU CHECK THE FILE YOU ARE TRYING TO IMPORT OR CARRY")
			}
		}
	}()
	// OS ERR non existant
	if x == os.ErrNotExist {
		if logging {
			fmt.Println(ErrorSymBolMap[CODE_FILE_INTEGRITY_FILE_INVALID_FILE_DOES_NOT_EXIST](FileCurrent.Get_Name()))
		}
		return false
	}
	// We do not want the file to be a directory but rather a full directory
	if f.IsDir() {
		File.IsDir = true
		if logging {
			fmt.Println(ErrorSymBolMap[CODE_FILE_INTEGRITY_FILE_INVALID_MUST_NOT_BE_DIRECTORY_DIR_UNSUPPORTED](FileCurrent.Get_Name()))
		}
		return false
	}
	return true
}

// Verifies the file is a CSC file
func (File *FileCurrentWithinParserEnvironment) Verify_CSC(logging bool) bool {
	if ext := filepath.Ext(File.Get_Name()); ext == ".csc" {
		return true
	} else {
		if logging {
			fmt.Println(ErrorSymBolMap[CODE_FILE_INTEGRITY_FILE_INVALID_MUST_BE_CSC_FILE](FileCurrent.Get_Name()))
		}
		return false
	}
}

// Open and check if the file is empty
func (File *FileCurrentWithinParserEnvironment) Verify_GoodToparse(logging bool) bool {
	f, x := os.Open(File.Get_Name())
	if x != nil {
		log.Fatal(x)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var t []string
	for scanner.Scan() {
		t = append(t, scanner.Text())
	}
	if t == nil {
		return false // bad for parser
	} else {
		return true
	}
}

// Return the body of the file within an array
func (File *FileCurrentWithinParserEnvironment) Get_Body(logging bool) []string {
	f, x := os.Open(File.Get_Name())
	if x != nil {
		log.Fatal(x)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var t []string
	for scanner.Scan() {
		t = append(t, scanner.Text())
	}
	if t != nil {
		return t
	} else {
		return nil
	}
}

// Inject new code into input file for execution of linker
func (File *FileCurrentWithinParserEnvironment) Inject_Body(infile string, input []string) {
	f, x := os.Create(infile)
	if x != nil {
		log.Fatal(x)
	}
	defer f.Close()
	for _, char := range input {
		if _, x = f.WriteString(char + "\n"); x != nil {
			log.Fatal(x)
		}
	}
}

// Delete in the case of injection
func (File *FileCurrentWithinParserEnvironment) Delete() {
	x := os.Remove(File.Get_Name())
	if x != nil {
		log.Fatal(x)
	}
}
