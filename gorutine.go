package main

import (
	"fmt"
	"runtime"
	"time"
)

func concurrent () {
	go func() {fmt.Println("task1")} ()
	go func() {fmt.Println("task2")} ()
	time.Sleep(time.Second)
}

func pararel () {
	runtime.GOMAXPROCS(4)
	go func() {fmt.Println("task1")} ()
	go func() {fmt.Println("task2")} ()
	time.Sleep(time.Second)
}