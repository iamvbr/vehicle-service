package services

import "errors"

var CarExists = errors.New("car exists")
var CarNotFound = errors.New("car not found")
