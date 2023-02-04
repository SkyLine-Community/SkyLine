package SkyLine

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	SkyLine_Crypto "main/Modules/SkyLine_Builtin/Cryptography"
	SkyLine_BuiltIn_System "main/Modules/SkyLine_Builtin/SystemFunctions"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var PI ProcessInformation

var builtins = map[string]*Builtin{
	"HASHIT": {
		Fn: func(args ...Object) Object {
			if len(args) != 2 {
				return newError(ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
					"HASHIT",
					fmt.Sprint(fmt.Sprint(2)),
					fmt.Sprint(len(args)),
				))
			}
			arg2 := args[1].Inspect()
			var newhash string
			switch args[0].Inspect() {
			case "MD5":
				newhash = SkyLine_Crypto.Hasher["MD5"](arg2)
			case "SHA1":
				newhash = SkyLine_Crypto.Hasher["SHA1"](arg2)
			case "SHA224":
				newhash = SkyLine_Crypto.Hasher["SHA224"](arg2)
			case "SHA256":
				newhash = SkyLine_Crypto.Hasher["SHA256"](arg2)
			case "SHA384":
				newhash = SkyLine_Crypto.Hasher["SHA384"](arg2)
			case "SHA512":
				newhash = SkyLine_Crypto.Hasher["SHA512"](arg2)
			}
			if newhash != "" {
				return &String{Value: newhash}
			} else {
				return &String{Value: "HASH empty, might have been an empty string or unicode"}
			}
		},
	},
	"OS_": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError(
					ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
						"OS_()",
						"1",
						fmt.Sprint(len(args)),
					),
				)
			}
			switch args[0].Inspect() {
			case "os_name":
				name, _ := SkyLine_BuiltIn_System.GrabOperatingSystemDataBasedOnKey["os_name"]()
				return &String{Value: name}
			case "os_arch":
				arch, _ := SkyLine_BuiltIn_System.GrabOperatingSystemDataBasedOnKey["os_arch"]()
				return &String{Value: arch}
			default:
				return &String{Value: "unknown value | run SkyLine__('OS') for more information"}
			}
		},
	},
	"USER_": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError(
					ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
						"USER_()",
						"1",
						fmt.Sprint(len(args)),
					),
				)
			}
			switch args[0].Inspect() {
			case "name":
				name, x := SkyLine_BuiltIn_System.GrabUserInformationFromOS["name"]()
				if x != nil {
					return newError("SkyLine backend (ERR_OS_INFO) => got error when working with OS information %s", x)
				} else {
					return &String{Value: name}
				}
			case "gid":
				gid, x := SkyLine_BuiltIn_System.GrabUserInformationFromOS["gid"]()
				if x != nil {
					return newError("SkyLine backend (ERR_OS_INFO) => got error when working with OS information %s", x)
				} else {
					return &String{Value: gid}
				}
			case "uid":
				uid, x := SkyLine_BuiltIn_System.GrabUserInformationFromOS["uid"]()
				if x != nil {
					return newError("SkyLine backend (ERR_OS_INFO) => got error when working with OS information %s", x)
				} else {
					return &String{Value: uid}
				}
			case "username":
				username, x := SkyLine_BuiltIn_System.GrabUserInformationFromOS["username"]()
				if x != nil {
					return newError("SkyLine backend (ERR_OS_INFO) => got error when working with OS information %s", x)
				} else {
					return &String{Value: username}
				}
			case "hdir":
				hdir, x := SkyLine_BuiltIn_System.GrabUserInformationFromOS["hdir"]()
				if x != nil {
					return newError("SkyLine backend (ERR_OS_INFO) => got error when working with OS information %s", x)
				} else {
					return &String{Value: hdir}
				}
			default:
				return &String{Value: "unknown value | run SkyLine__('USER') for more information"}
			}
		},
	},
	"SkyLine__": {
		Fn: func(args ...Object) Object {
			if len(args) != 1 {
				return newError(
					ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
						"SkyLine_()",
						"1",
						fmt.Sprint(len(args)),
					),
				)
			}
			var msg string
			switch args[0].Inspect() {
			case "OPERATORS":
				// List of operators
			case "OS":
				msg += `
				OS or Operating System is a standard SkyLine 
				function to grab or view information about 
				the current operating system in which the 
				SkyLine interpreter is running on. This 
				function has the following values

				OS_("os_name")    | Grabs the current operating system
				OS_("os_arch")    | Grabs the current operating system architecture
				`
			case "USER":
				msg += `
                USER or Username is a standard SkyLine 
                function to grab or view information about 
                the current user in which the SkyLine 
                interpreter is running on. This function has 
                the following values
				
				USER_("username")    | Grabs the current username
				USER_("uid")         | Grabs the current uid
				USER_("gid")         | Grabs the current gid
				USER_("name")        | Grabs the name
				USER_("hdir")        | Grabs the home directory of the user
				`
			default:
				msg += `METHOD DOES NOT EXIST -> `
				msg += args[0].Inspect()
			}
			return &String{Value: msg}
		},
	},

	"length": {
		Fn: func(args ...Object) Object {
			if l := len(args); l != 1 {
				return newError(
					ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
						"length()",
						"1",
						fmt.Sprint(l),
					),
				)
			}

			switch arg := args[0].(type) {
			case *String:
				return &Integer{Value: int64(len(arg.Value))}
			case *Array:
				return &Integer{Value: int64(len(arg.Elements))}
			default:
				return newError("argument to `len` not supported, got %s", arg.Type_Object())
			}
		},
	},
	"reverse": {
		Fn: func(args ...Object) Object {
			if l := len(args); l != 1 {
				return newError(
					ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
						"reverse()",
						"1",
						fmt.Sprint(l),
					),
				)
			}

			if typ := args[0].Type_Object(); typ != ArrayType {
				return newError("argument to `last` must be Array, got %s", typ)
			}
			var a []Object
			arr := args[0].(*Array)
			l := len(arr.Elements)
			if l == 0 {
				return NilValue
			}
			for i := 0; i < len(arr.Elements); i++ {
				a = append(a, arr.Elements[i])
			}
			ReverseObjectArrayWithinCSCF(a)
			newElems := a
			return &Array{Elements: newElems}
		},
	},
	"first": {
		Fn: func(args ...Object) Object {
			if l := len(args); l != 1 {
				return newError(
					ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
						"first()",
						"1",
						fmt.Sprint(l),
					),
				)
			}

			if typ := args[0].Type_Object(); typ != ArrayType {
				return newError("argument to `first` must be Array, got %s", typ)
			}

			arr := args[0].(*Array)
			if len(arr.Elements) == 0 {
				return NilValue
			}
			return arr.Elements[0]
		},
	},

	"last": {
		Fn: func(args ...Object) Object {
			if l := len(args); l != 1 {
				return newError(
					ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
						"last()",
						"1",
						fmt.Sprint(l),
					),
				)
			}

			if typ := args[0].Type_Object(); typ != ArrayType {
				return newError("argument to `last` must be Array, got %s", typ)
			}

			arr := args[0].(*Array)
			l := len(arr.Elements)
			if l == 0 {
				return NilValue
			}
			return arr.Elements[l-1]
		},
	},

	"rest": {
		Fn: func(args ...Object) Object {
			if l := len(args); l != 1 {
				return newError(
					ErrorSymBolMap[CODE_WRONG_NUMBER_OF_ARGUMENTS](
						"rest()",
						"1",
						fmt.Sprint(l),
					),
				)
			}

			if typ := args[0].Type_Object(); typ != ArrayType {
				return newError("argument to `last` must be Array, got %s", typ)
			}

			arr := args[0].(*Array)
			l := len(arr.Elements)
			if l == 0 {
				return NilValue
			}

			newElems := make([]Object, l-1)
			copy(newElems, arr.Elements[1:l])
			return &Array{Elements: newElems}
		},
	},

	"push": {
		Fn: func(args ...Object) Object {
			if l := len(args); l != 2 {
				return newError("wrong number of arguments. want=%d, got=%d", 2, l)
			}

			if typ := args[0].Type_Object(); typ != ArrayType {
				return newError("first argument to `push` must be Array, got %s", typ)
			}

			arr := args[0].(*Array)
			l := len(arr.Elements)

			newElems := make([]Object, l+1)
			copy(newElems, arr.Elements)
			newElems[l] = args[1]
			return &Array{Elements: newElems}
		},
	},

	"print": {
		Fn: func(args ...Object) Object {
			for _, arg := range args {
				fmt.Print(arg.Inspect())
			}
			return &String{Value: ""}
		},
	},
	"println": {
		Fn: func(args ...Object) Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			return &String{Value: ""}
		},
	},
	"sprint": {
		Fn: func(args ...Object) Object {
			if l := len(args); l == 0 {
				return newError("SkyLine Builtin (sprint): SPRINT function requires 1 (argument), the argument should be a variable you want to convert to a string")
			}
			return &String{Value: fmt.Sprint(args[0].Inspect())}
		},
	},
	"input": {
		Fn: func(args ...Object) Object {
			if l := len(args); l != 2 {
				return newError("wrong number of arguments. want=1, got=%d | SkyLine's builtin functions such as INPUT require you to enter a character and the name of the input such as 'input' and 'n' where n is the second argument to tell the parser when to use that input. Current supported characters are (n) -> newline ", l)
			}
			input := bufio.NewReader(os.Stdin)
			var Payload string
			fmt.Print(args[0].Inspect())
			for {
				switch args[1].Inspect() {
				case "newline":
					Payload, _ = input.ReadString('\n')              // read input until new line
					Payload = strings.Replace(Payload, "\n", "", -1) // read and replace input state
				case "n":
					Payload, _ = input.ReadString('\n')              // read input until new line
					Payload = strings.Replace(Payload, "\n", "", -1) // read and replace input state
				}
				if Payload != "" {
					break
				}
			}
			return &String{Value: Payload}
		},
	},
	"uppercase": {Fn: func(args ...Object) Object {
		if l := len(args); l != 1 {
			return newError("wrong number of arguments. want=1, got=%d", l)
		}
		switch args[0].(type) {
		case *String:
			return &String{Value: strings.ToUpper(args[0].(*String).Value)}
		default:
			return newError("argument to `uppercase` must be String, got %s", args[0].Type_Object())
		}
	}},
	"lowercase": {Fn: func(args ...Object) Object {
		if l := len(args); l != 1 {
			return newError("wrong number of arguments. want=1, got=%d", l)
		}
		switch args[0].(type) {
		case *String:
			return &String{Value: strings.ToLower(args[0].(*String).Value)}
		default:
			return newError("argument to `lowercase` must be String, got %s", args[0].Type_Object())
		}
	}},
	"typeof": {Fn: func(args ...Object) Object {
		if l := len(args); l != 1 {
			return newError("wrong number of arguments. want=1, got=%d", l)
		}
		return &String{Value: fmt.Sprint(args[0].Type_Object())}
	}},
	"index": {Fn: func(args ...Object) Object {
		if l := len(args); l != 2 {
			return newError("wrong number of arguments. want=2, got=%d", l)
		}
		var StrToIndex string
		var LookingFor string
		switch args[0].(type) {
		case *String:
			StrToIndex = args[0].(*String).Value
		default:
			return newError("argument to `index` at (1) must be String, got %s", args[0].Type_Object())
		}
		switch args[1].(type) {
		case *String:
			LookingFor = args[1].(*String).Value
		default:
			return newError("argument to `index` at (2) must be String, got %s", args[1].Type_Object())
		}
		IDX, x := strconv.ParseInt(fmt.Sprint(strings.Index(StrToIndex, LookingFor)), 10, 64)
		if x != nil {
			log.Fatal(x)
		}
		return &Integer{Value: IDX}
	}},
	"exit": {Fn: func(args ...Object) Object {
		if l := len(args); l != 1 {
			return newError("wrong number of arguments. want=1, got=%d", l)
		}
		var code int64
		switch args[0].(type) {
		case *Integer:
			code = int64(args[0].(*Integer).Value)
		default:
			return newError("argument to `exit` at (1) must be Integer, got %s", args[0].Type_Object())
		}
		if code == 0 {
			os.Exit(0)
		} else {
			IDX, x := strconv.ParseInt(fmt.Sprint(code), 10, 64)
			if x != nil {
				log.Fatal(x)
			}
			os.Exit(int(IDX))
		}
		return &Nil{}
	}},
	"sleep": {Fn: func(args ...Object) Object {
		if l := len(args); l != 2 {
			return newError("wrong number of arguments. want=2, got=%d", l)
		}
		var t time.Duration
		var timer int
		switch l := args[0].(type) {
		case *String:
			t = SkyLine_BuiltIn_System.SleepCodes[l.Inspect()]
		default:
			return newError("argument to `sleep` at (1) must be String, got %s", args[0].Type_Object())
		}
		switch args[1].(type) {
		case *Integer:
			timer = int(args[1].(*Integer).Value)
		default:
			return newError("argument to `sleep` at (2) must be Integer, got %s", args[1].Type_Object())
		}
		time.Sleep(time.Duration(timer) * t)
		return &Nil{}
	}},
	"repeat": {Fn: func(args ...Object) Object {
		if l := len(args); l != 2 {
			return newError("wrong number of arguments. want=2, got=%d", l)
		}
		var count int
		var str string
		switch args[0].(type) {
		case *Integer:
			count = int(args[0].(*Integer).Value)
		default:
			return newError("argument to `repeat` at (1) must be Integer, got %s", args[0].Type_Object())
		}
		switch args[1].(type) {
		case *String:
			str = args[1].(*String).Value
		default:
			return newError("argument to `repeat` at (2) must be String, got %s", args[1].Type_Object())
		}
		return &String{Value: strings.Repeat(str, count)}
	}},
	"Isnil": {Fn: func(args ...Object) Object {
		if l := len(args); l != 1 {
			return newError("wrong number of arguments. want=1, got=%d", l)
		}
		if args[0].Inspect() == "" {
			return &Boolean_Object{Value: true}
		} else {
			return &Boolean_Object{Value: false}
		}
	}},
	// Processing
	"LoadProcess": {Fn: func(args ...Object) Object {
		if l := len(args); l != 1 {
			return newError("wrong number of arguments. want=1, got=%d", l)
		}
		switch args[0].(type) {
		case *String:
			PI.PIDbyProgramName(args[0].Inspect())
			return &String{Value: fmt.Sprint(PI.ProcessID)}
		default:
			return &Error{
				Message: "SkyLine erorr: Mismatched data types,function GetProcess requires a value of type string",
			}
		}
	}},
	"FIRE_PROC": {Fn: func(args ...Object) Object {
		if l := len(args); l != 1 {
			return newError("wrong number of arguments. want=2, got=%d", l)
		}
		if PI.ProcessID == "" {
			return newError("False ID : Process ID must be specified by calling GetProcess...")
		}
		file := fmt.Sprintf("/proc/%s/fd/0", PI.ProcessID)
		x := ioutil.WriteFile(file, []byte(args[0].Inspect()), 0600)
		VerErr(x)
		return &Boolean_Object{Value: true}
	}},
	// Script tags...
	"__ID__": {Fn: func(args ...Object) Object {
		result := os.Getpid()
		return &String{
			Value: fmt.Sprint(result),
		}
	}},
	"__NAME__": {Fn: func(args ...Object) Object {
		return &String{
			Value: os.Args[0],
		}
	}},
	// Mathematical operations standard | These operations will be base such as ATAN, ATAN2, SQRT, ABS, TAN, SQRT and including other small operations
	// tale : make mathematical module
	"sqrt": {Fn: func(args ...Object) Object {
		if l := len(args); l != 1 {
			return newError("wrong number of arguments. want=1, got=%d", l)
		}
		var num float64
		var x error
		switch args[0].(type) {
		case *Float:
			num, x = strconv.ParseFloat(fmt.Sprint(args[0].Inspect()), 64)
		case *Integer:
			num, x = strconv.ParseFloat(fmt.Sprint(args[0].Inspect()), 64)
		default:
			return newError("Wrong data type, mathematical function SQRT (Square Root) requires either integer or float argument")
		}
		if x != nil {
			return &Error{Message: fmt.Sprintf("Format error when trying to convert data type to float -> %s ", x)}
		} else {
			return &Float{Value: math.Sqrt(num)}
		}
	}},
}

type FileExec struct {
	FileToLoad string
}

type FileInformation struct {
	IsDir     bool
	IsRegular bool
	Size      int64
	Name      string
	Base      string
	Lines     []string
	LineC     int
	Fdata     map[string]string
	Mode      string
	ModTime   string
}

var fexec FileExec
var finfo FileInformation

// File information map
var FInfo = map[string]*Builtin{
	"new": {Fn: func(args ...Object) Object {
		var x error
		var f fs.FileInfo
		if len(args) != 1 {
			return newError("wrong number of arguments. want=1, got=%d", len(args))
		}
		var a string
		switch args[0].(type) {
		case *String:
			a = args[0].(*String).Value
		default:
			return newError("argument to `new` at (1) must be String, got %s", args[0].Type_Object())
		}
		if f, x = os.Stat(a); x == nil {
			fexec.FileToLoad = a
			finfo.IsDir = f.IsDir()
			finfo.IsRegular = f.Mode().IsRegular()
			finfo.Name = f.Name()
			finfo.Size = f.Size()
			finfo.ModTime = f.ModTime().String()
			finfo.Mode = f.Mode().String()
			return &Nil{}
		}
		return &Error{Message: "File does not exist"}
	}},
	"exists?": {Fn: func(args ...Object) Object {
		if _, x := os.Stat(fexec.FileToLoad); os.IsNotExist(x) {
			return &Boolean_Object{Value: false}
		} else {
			return &Boolean_Object{Value: true}
		}
	}}, // File exists
	"fsize": {Fn: func(args ...Object) Object {
		return &Integer{Value: finfo.Size}
	}}, // File size
	"isdir?": {Fn: func(args ...Object) Object {
		return &Boolean_Object{Value: finfo.IsDir}
	}}, // Is a directory
	"isfile?": {Fn: func(args ...Object) Object {
		return &Boolean_Object{Value: finfo.IsRegular}
	}}, // Is a file
	"flinec": {Fn: func(args ...Object) Object {
		f, x := os.Open(finfo.Name)
		if x != nil {
			log.Fatal(x) // change per error
		}
		defer f.Close()
		var l int
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			l++
		}
		return &Integer{Value: int64(l)}
	}}, // File data or lines
	"fmodtime": {Fn: func(args ...Object) Object {
		return &String{Value: fmt.Sprint(finfo.ModTime)}
	}}, // File modification time
	"fmode": {Fn: func(args ...Object) Object {
		return &String{Value: finfo.Mode}
	}}, // File mode
	"fext": {Fn: func(args ...Object) Object {
		x := filepath.Ext(finfo.Name)
		return &String{Value: x}
	}}, // Extension
}
