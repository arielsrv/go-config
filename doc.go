// Package config
// go-config is a configuration library
// for accessing a mixture of static as well as dynamic configurations as a single configuration unit.
// There are two key concepts to note:
//
// - Properties that can be read by your code.
// - Configurations that organize properties into objects you can bootstrap your application with.

package config

import (
	_ "github.com/arielsrv/go-config/autoload"
	_ "github.com/arielsrv/go-config/env"
)
