package filelog

import (
	"os"
	"os/user"
	"fmt"
	"time"
	"github.com/Syssos/Go_Shell/color"
)

// Flog structure to keep track of needed material
type Flog struct {

	Greeting, Salutation, LogFile string
	Errormsg error
}

// Prints greeting message and logs when the shell instance was started
func (m Flog) Greet() {

	dt := time.Now()
	m.Log(fmt.Sprintf("%v - User %v started instance",  dt.Format("01-02-2006 15:04:05"), GetUser()))
	fmt.Println(m.Greeting, "\n")
}

// Prints salute message and logs when the shell instance is closed
func (m Flog) Salute() {

	dt := time.Now()
	m.Log(fmt.Sprintf("%v - User %v exited instance",  dt.Format("01-02-2006 15:04:05"), GetUser()))
	fmt.Println(m.Salutation)
}

// Prints error message and logs it
func (m Flog) Err() {

	dt := time.Now()
	m.Log(fmt.Sprintf("%v - %v",  dt.Format("01-02-2006 15:04:05"), m.Errormsg))
	fmt.Println(color.Red + fmt.Sprintf("%v", m.Errormsg) + color.Reset)
}

// Command that opens log file and writes content to it.
func (m Flog) Log(msg string) {

	f, err := os.OpenFile(m.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	    fmt.Println(err)
	}
	
	f.Write([]byte(msg+"\n"))
	f.Close()
}

// Command to get the current users name
func GetUser() string{

	use, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	return use.Name
}

// Gets the current user direectory 
func GetHomeDir() string {

	dirname, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	return dirname
}

// Gets the current working directory for the current directory the user is in when the program is ran.
func GetCurrentDir() string{

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return cwd
}