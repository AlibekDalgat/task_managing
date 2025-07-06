package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
}
