package scan

import "fmt"

func Run() {
	var a string
	_, err := fmt.Sscan("job-01.aJ7", &a)

	fmt.Println(a, err)

	var instance string
	_, err = fmt.Sscanf("job-01.aJ7", "%s.%s", &instance)

	fmt.Println(instance, err)
}
