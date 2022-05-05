package main

import (
  "os
  "..."
)

func main() {
  bot := twillight.Beat()
  bot.Config("token", os.Getenv("TOKEN"))

  bot.Beat("webhook", twillight.WebHook{URL: "https://site.com",})
}
