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
	pictureGenerator func(param string, resC chan *AiResponse,