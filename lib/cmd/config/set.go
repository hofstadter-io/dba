package libcmdconfig

import (
	"fmt"
)

func SetRun(name string, host string, account string, project string) (err error) {

	// Default body

	fmt.Println("dba config set", name, host, account, project)

	return err
}
