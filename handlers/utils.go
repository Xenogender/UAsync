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
	para