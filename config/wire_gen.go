// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

// Injectors from injector.go:

func Build() Config {
	config := Provider()
	return config
}
