package libcmdconfig

import (
	"fmt"
)

func SetRun(name string, host string, account string, project string) (err error) {

	// Default body

	fmt.Println("dma config set", name, host, account, project)

	return err
}
