//+build mage

package main

import (
	"go/build"

	"github.com/magefile/mage/sh"
)

func Install() error {

	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	return sh.Run("go", "build", "-o", build.Default.GOPATH+"/bin/hoodieaf")
}

func Remove() error {

	return sh.Run("rm", build.Default.GOPATH+"/bin/hoodieaf")
}
