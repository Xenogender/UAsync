
<h1 align="center">
  <br>
  <img src="./.github/img/golang_logo.png" alt="Golang logo" width="200">
  <br>
  GO DIEGO - BOT
  <br>
</h1>

<h4 align="center">Bot features made with golang for <a href="https://www.discord.com" target="_blank">Discord</a> (is unfinished and close project yet).</h4>

<p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#download">Download</a> •
  <a href="#credits">Credits</a> •
  <a href="#related">Related</a> •
  <a href="#license">License</a> •
  <a href="#others">Others</a>
</p>

## Key Features

* Command prefix
  - The bot responses with a command prefix, that's a "!".
* Ping-pong
  - If you send !ping, the bot responses pong, else if !pong, the bot responses ping.
* Help java!
  - if you type a phrase that have java in the string, the bot will tag the some members (i add example member ID) in the server to help you. (CUSTOMIZABLE)
* Greeting
  - The bot will greeting you if you send "!Oi diego" to him.
* Image generator
  - Diego bot responses with an image if you send "!picture (your substantive).

## How To Use

To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://go.dev/learn/) and the [Docker](https://www.docker.com/) to run the app in container 
```bash
# Clone this repository
$ git clone https://github.com/Leoff00/go-diego-bot.git

# Go into the repository
$ cd go-diego-bot

# Run the project locally
$ go mod download && go mod tidy && go run main.go
