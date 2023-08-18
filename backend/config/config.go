package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	//Db settings
	DBURI     string
	DBTimeout int

	//Server settings
	ServerPort  string
	GatewayPort string
}

// Instantiate a new configuration -- reads from .env file
func New() (*Config, error) {

	readEnvfile()

	dbURI, err := stringEnvVar("DB_URI")
	if err != nil {
		return nil, fmt.Errorf("failed to get DB_URI: %v", err)
	}

	dbTimeout, err := intEnvVar("DB_TIMEOUT")
	if err != nil {
		return nil, fmt.Errorf("failed to get DB_TIMEOUT: %v", err)
	}

	serverPort, err := stringEnvVar("SERVER_PORT")
	if err != nil {
		return nil, fmt.Errorf("failed to get SERVER_PORT: %v", err)
	}

	gatewayPort, err := stringEnvVar("GATEWAY_PORT")
	if err != nil {
		return nil, fmt.Errorf("failed to get GATEWAY_PORT: %v", err)
	}

	return &Config{
		DBURI:       dbURI,
		DBTimeout:   dbTimeout,
		ServerPort:  serverPort,
		GatewayPort: gatewayPort,
	}, nil
}

func readEnvfile() {
	//Read the .env file
	cwdEnvPath, err := filepath.Abs(".env")
	if err == nil {
		err = godotenv.Load(cwdEnvPath)
		if err == nil {
			log.Println("Loaded .env file")
			return
		}
		fmt.Println(cwdEnvPath)
	}
	err = godotenv.Load("../../.env")
	if err != nil {
		log.Println("Loading .env file failed, using production environment")
	}
	if err == nil {
		log.Println("Loaded .env file")
	}
}

func stringEnvVar(envname string) (string, error) {
	val, ok := os.LookupEnv(envname)
	if !ok {
		return "", fmt.Errorf("missing env var '%s' string value", envname)
	}

	return val, nil
}

func intEnvVar(envname string) (int, error) {
	val, ok := os.LookupEnv(envname)
	if !ok {
		return 0, fmt.Errorf("missing env var '%s' int value", envname)
	}
	//Convert val to int
	intVal, err := strconv.Atoi(val)
	if err != nil {
		return 0, fmt.Errorf("failed to convert env var '%s' to int: %v", envname, err)
	}
	return intVal, nil
}
