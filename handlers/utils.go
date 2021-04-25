package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/Leoff00/go-diego-bot/envs"
)

type HandlerUtilFunctions struct {
	paramSeparator   func(content string) string
	randPhrase       func(user string) string
	pictureGenerator func(param string, resC chan *AiResponse, errC chan error) (chan *AiResponse, chan error)
}

func (hu *HandlerUtilFunctions) ParamSeparator(message string) string {
	splitted := strings.Split(message, " ")
	if len(splitted) < 2 {
		log.Default().Fatalln("Param is required.")
	}

	return splitted[1]

}

func (hu *HandlerUtilFunctions) RandPh(user string) string {
	rand.Seed(time.Now().Unix())
	g1 := fmt.Sprintf("Ola %s!", user)
	g2 := fmt.Sprintf("Iaee %s!", user)
	g3 := fmt.Sprintf("Oiee %s!", user)

	arr = append(arr, g1, g2, g3)
	return arr[rand.Intn(len(arr))]
}

func (hu *HandlerUtilFunctions) PicGenerator(param string, resC chan *AiResponse, errC chan error) (chan *AiResponse, chan error) {

	rand.Seed(time.Now().Unix())

	c := &http.Client{}

	if regexp.MustCompile(param).MatchString(`$\s\D`) {
		errC <- errors.New("Cannot send more 