package initializers

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"log"
)

var Enforcer *casbin.Enforcer

func InitCasbin(config *Config) {
	// Load the model from a file
	m, err := model.NewModelFromFile(config.CasbinModelPath)
	if err != nil {
		message, _ := fmt.Printf("failed to load model: %s", err)
		panic(message)
	}

	a := fileadapter.NewAdapter(config.CasbinPolicyPath)

	// Create the enforcer with the model and the adapter
	Enforcer, err = casbin.NewEnforcer(m, a)
	if err != nil {
		message, _ := fmt.Printf("failed to create enforcer: %s", err)
		panic(message)
	}

	// Load the policies from the adapter (file)
	err = Enforcer.LoadPolicy()
	if err != nil {
		log.Fatalf("failed to load policy: %v", err)
	}
}
