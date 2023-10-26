package main

import (
	"github.com/sirupsen/logrus"
	"github.com/yzx9/motion/interface/web"
)

func main() {
	log := logrus.WithField("topic", "main")
	log.Fatal(web.New())
}
