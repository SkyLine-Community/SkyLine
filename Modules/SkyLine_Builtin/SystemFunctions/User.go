package SkyLine_SYSTEM

import "os/user"

var GrabUserInformationFromOS = map[string]func() (string, error){
	"username": func() (string, error) {
		user, x := user.Current()
		if x != nil {
			return "FAIL", x
		} else {
			return user.Username, nil
		}
	},
	"gid": func() (string, error) {
		user, x := user.Current()
		if x != nil {
			return "FAIL", x
		} else {
			return user.Gid, nil
		}
	},
	"uid": func() (string, error) {
		user, x := user.Current()
		if x != nil {
			return "FAIL", x
		} else {
			return user.Uid, nil
		}
	},
	"hdir": func() (string, error) {
		user, x := user.Current()
		if x != nil {
			return "FAIL", x
		} else {
			return user.HomeDir, nil
		}
	},
	"name": func() (string, error) {
		user, x := user.Current()
		if x != nil {
			return "FAIL", x
		} else {
			return user.Name, nil
		}
	},
}
