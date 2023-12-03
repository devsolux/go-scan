package main

import (
	"fmt"
	"github.com/devsolux/go-scan/Plugins"
	"github.com/devsolux/go-scan/common"
	"time"
)

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	fmt.Printf("[*] Scanning completed, duration: %s\n", time.Since(start))
}
