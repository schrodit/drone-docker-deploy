package main

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

const netrcFile = `
machine %s
login %s
password %s
`

// helper function to write a netrc file.
func writeNetrc(machine, login, password string) error {
	if machine == "" {
		return nil
	}
	out := fmt.Sprintf(
		netrcFile,
		machine,
		login,
		password,
	)

	home := "/root"
	u, err := user.Current()
	if err == nil {
		home = u.HomeDir
	}
	path := filepath.Join(home, ".netrc")
	return ioutil.WriteFile(path, []byte(out), 0600)
}

// func debugGit() {
// 	cmd := exec.Command("env")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	if err := cmd.Run(); err != nil {
// 		log.Printf("error in debug\n%v", err)
// 	}

// 	cmd = exec.Command("cat", "/root/.netrc")
// 	cmd.Stdout = os.Stdout
// 	cmd.Stderr = os.Stderr
// 	if err := cmd.Run(); err != nil {
// 		log.Printf("error in debug\n%v", err)
// 	}
// }
