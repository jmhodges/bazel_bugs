package main

import (
	"fmt"

	core "k8s.io/client-go/kubernetes/typed/core/v1"
)

func main() {
	var secrets core.SecretInterface
	fmt.Println(secrets)
}
