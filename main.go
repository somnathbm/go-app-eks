package main

import (
	"fmt"
	"reflect"
	"runtime"
)

// required environmemts to create
var environments = []string{"dev", "staging", "prod"}

// runtime info from the runtime package
type RuntimeInfo struct {
	os, gopath string
}

func GetCompilerInfo() {
	sys_info := RuntimeInfo{runtime.GOOS, runtime.GOROOT()}
	sys_info_fields := reflect.ValueOf(sys_info)
	for i := 0; i < sys_info_fields.NumField(); i++ {
		fmt.Printf("[INFO] %s: %s\n", sys_info_fields.Type().Field(i).Name, sys_info_fields.Field(i))
	}
}

func GetEnvs() {
	fmt.Println("\n===========================================")
	fmt.Println("[INFO] Getting requested environments to provision..")
	fmt.Println("\n===========================================")
	for i := 0; i < len(environments); i++ {
		fmt.Printf("[INFO] environment \"%s\" will be provisioned shortly\n", environments[i])
	}
}

func GetProvisioningInfo() {
	fmt.Println("[INFO] Gathering information. Please wait..")
	fmt.Println("===========================================")
	fmt.Println("[INFO] Getting GOLANG compiler info")
	fmt.Println("===========================================")
	GetCompilerInfo()
	GetEnvs()
}

func main() {
	GetProvisioningInfo()
}
