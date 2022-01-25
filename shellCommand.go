package main

import (
"fmt"
	"net/url"
	"os/exec"
"runtime"
)

func execute() {

	// here we perform the pwd command.
	// we can store the output of this in our out variable
	// and catch any errors in err
	podname:="ibm-storage-metrics-agent-7987948688-2j9kz"
	route:="https://private.containers.test.cloud.ibm.com"
	u, _ := url.Parse(route)
	cmd := fmt.Sprintf("kubectl exec %s -n kube-system -- bash -c 'ping -c 2 %s'", podname,u.Host)
	out, err := exec.Command("bash", "-c",cmd).Output()

	// if there is an error with our execution
	// handle it here
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(out)
	// as the out variable defined above is of type []byte we need to convert
	// this to a string or else we will see garbage printed out in our console
	// this is how we convert it to a string
	fmt.Println("Command Successfully Executed")
}

 func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Can't Execute this on a windows machine")
	} else {
		execute()
	}
}
