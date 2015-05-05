package main

import (
	"fmt"
	"testing"
)

func TestOToFE(t *testing.T) {
	argList := []string{
		"-o", "test.exe", "test.c",
	}

	w := wrapper{}
	w.setArgs(argList)

	if w.argString() != "/Fe:test.exe test.c" {
		fmt.Println(w.argString())
		t.Error()
	}

}

func TestOToFEMultipleSpaces(t *testing.T) {
	argList := []string{
		"-o", "    test.exe", "test.c",
	}
	w := wrapper{}
	w.setArgs(argList)
	args := w.argString()
	if args != "/Fe:test.exe test.c" {
		fmt.Println(args)
		t.Error()
	}

}

func TestOToFENoSpaces(t *testing.T) {
	argList := []string{
		"-otest.exe", "test.c",
	}
	w := wrapper{}
	w.setArgs(argList)
	args := w.argString()
	if args != "/Fe:test.exe test.c" {
		fmt.Println(args)
		t.Error()
	}

}

func TestOToFo(t *testing.T) {
	argList := []string{
		"-o", "test.obj", "test.c",
	}
	w := wrapper{}
	w.setArgs(argList)
	args := w.argString()
	if args != "/Fo:test.obj test.c" {
		fmt.Println(args)
		t.Error()
	}

}

func TestOToFoMultipleSpaces(t *testing.T) {
	argList := []string{
		"-o", "    test.obj", "test.c",
	}
	w := wrapper{}
	w.setArgs(argList)
	args := w.argString()

	if args != "/Fo:test.obj test.c" {
		fmt.Println(args)
		t.Error()
	}

}

func TestOToFoNoSpaces(t *testing.T) {
	argList := []string{
		"-otest.obj", "test.c",
	}
	w := wrapper{}
	w.setArgs(argList)
	args := w.argString()
	if args != "/Fo:test.obj test.c" {
		fmt.Println(args)
		t.Error()
	}

}
