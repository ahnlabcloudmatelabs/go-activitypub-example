package constants

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	ENV  = "local"
	PORT = ""

	DB_TYPE     = ""
	DB_HOST     = ""
	DB_PORT     = ""
	DB_NAME     = ""
	DB_USER     = ""
	DB_PASSWORD = ""

	APP_NAME             = ""
	APP_VERSION          = ""
	APP_ADDRESS          = ""
	APP_DESCRIPTION      = ""
	APP_MAINTAINER_NAME  = ""
	APP_MAINTAINER_EMAIL = ""

	USER_JSON_URL_FORMAT = ""
	USER_HTML_URL_FORMAT = ""
	USER_JSON_ENDPOINT   = ""
	USER_HTML_ENDPOINT   = ""

	ACTIVITY_JSON_CONTENT_TYPE = `application/ld+json; profile="https://www.w3.org/ns/activitystreams"`
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	setEnv()
}

func setEnv() {
	ENV = os.Getenv("ENV")
	PORT = os.Getenv("PORT")

	DB_TYPE = os.Getenv("DB_TYPE")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")

	APP_NAME = os.Getenv("APP_NAME")
	APP_VERSION = os.Getenv("APP_VERSION")
	APP_ADDRESS = os.Getenv("APP_ADDRESS")
	APP_DESCRIPTION = os.Getenv("APP_DESCRIPTION")
	APP_MAINTAINER_NAME = os.Getenv("APP_MAINTAINER_NAME")
	APP_MAINTAINER_EMAIL = os.Getenv("APP_MAINTAINER_EMAIL")

	USER_JSON_URL_FORMAT = os.Getenv("USER_JSON_URL_FORMAT")
	USER_HTML_URL_FORMAT = os.Getenv("USER_HTML_URL_FORMAT")
	USER_JSON_ENDPOINT = os.Getenv("USER_JSON_ENDPOINT")
	USER_HTML_ENDPOINT = os.Getenv("USER_HTML_ENDPOINT")
}
