# Filelog

<p align="center">
  <img src="https://github.com/Syssos/Go_Shell/blob/main/etc/images/Error_Log_Example.PNG" alt="Go Shell img"/>
  <p align="center">Example held in <a href="https://github.com/Syssos/Go_Shell/blob/main/settings/example_log.txt">settings/example_log.txt</a></p>
</p>

Filelog is a tool designed to quickly and easily log an incident to a file. This was originally created to log any errors as well as a start/end timestamp.

## Install
To start off the package needs to be collected from github. You can do this via the go get command
```
go get github.com/Syssos/Go_Shell/filelog
```

After the package is downloaded it should be able to be imported and used
```go
import "github.com/Syssos/Go_Shell/filelog"
```
## Usage

This package was designed to keep logging lightweight and allow for an easy way to implement multiple log files.

To use this package, create an instance of the ``` Flog ``` struct.

```go
flog := filelog.Flog{ Greeting, Salute, LogFile, DtFormat, location, nil}
```
**Note the last value when creating a Flog struct is nil because it will hold an err if one is present**

Once the instance is created, to log a file, the Log method can be utilized.

### Logging strings (no timestamp)

```go
LogOne.Log("Message to log")
```
### Logging errors (logs timestamp)

```go
LogOne.Errormsg = errors.New("Error to log")
LogOne.Err()
```
### Logging errors from function return

```go
res, err := somefun(a, b)
if err != nil {
	LogOne.Errormsg = err
	LogOne.Err()
}
```

The code above will log the error message to the log file as well as print the error to the screen using a Println function.

If you are interested in how toml was utilized, to store preset values for the logger in the Go Shell project, look inside of the [main.go](https://github.com/Syssos/Go_Shell/blob/main/main.go) file is an example of the Flog instance being created from a TOML file.
