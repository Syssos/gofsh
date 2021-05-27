# Go Shell [![Build Status](https://travis-ci.com/Syssos/Go_Shell.svg?branch=main)](https://travis-ci.com/github/Syssos/Go_Shell)  [![Cody Code](https://syssos.app/static/images/index/cody_code.svg)](https://syssos.app)

![Go Shell img](https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Go_Shell.png)

This is a simple shell built in go designed to give you basic functionality such as changing directories or listing files. This is intended to be used in future projects, if need arise for a shell. Over time this shell will develop to account for tasks I use often such as running nmap scans or sorting/refining output files.


![Go Shell Example](https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Shell.PNG)

## Install

All of the packages used in this repository were designed with external usablility in mind. Each package directory will contain a readme with all of the information you'll need to get started. The directions below should get you started with using the shell as.

These instructions are for targeted towards linux user's.

To start things off this will require some knowledge of working with Go modules, and packages. As well as how to build and install programs. If you have any questions about Go or how to do these things I just so happen to have a repository [here](https://github.com/Syssos/Learning_Go) I created as I learned Go, leaving notes behind that may be useful to you.

To get the files we need clone the package to a location you can use.

``` 
git clone https://github.com/Syssos/Go_Shell.git
```
Once the repository is cloned cd into it.
```
cd Go_Shell
```
Before we continue I need to note there is no go.mod file included in this repository so one will need to be created. The command I used to create my module is below.

```
go mod init github.com/Syssos/Go_Shell
```
If this is not the module you choose you will need to either run the ``` go get ``` command for the packages included, or change the import statement in the nested packages.

Once the go.mod file is created we can install it.

```
go install .
```

This will place a bin file in "$PATH/bin" called "Go_Shell", if the module name was used above. Running this command will give you the Go Shell.

For me this looks like

```
~/go/bin/Go_Shell
```
Alternativly you can install the "cmds" package with ``` go get ``` and create a script that utilizes the Loop function much like [main.go](https://github.com/Syssos/Go_Shell/blob/main/main.go) does.

## Usage

There are currently 3 working commands and no help system set up yet. I aplogize, but becuase this project is actively being worked on there should be self help system soon. To make up for the lack of self help below I will list the commands and how to utilize them

### ls
This command should work just about the same as it does an a native linux system, minus advanced functionality. While it takes a location to list it cannot except flags at the moment.

```
$ ls
$ ls ../
$ ls /Go_Shell
```
### pwd
This command will print the working directory. Its a pretty straight forward command and doesn't take arguments.

```
$ pwd
```
### cd
This command changes the working directory much like cd in linux, like ls this will take a location to change to but will not accept flags.

```
$ cd ../
$ cd Go_Shell/
```
### exit
This exits the shell properly and allows for logging of the user quiting.

```
$ exit
```
## Main.go

This file will call the Loop function from the cmds package. This loop is what is responsible for all of the commands. The plans to re-use this code mean I need all of the code in one package, more or less, that I can grab and use in another project.

## Logging

This program will log errors to a file called "GoShellLogfile.txt". Due to the state of this shell the logging file location is undecided, for know the shell will live int the ```/tmp``` directory. Be aware that this file location may change in future version as there could be a need for it in a new location to better suit personal use cases.

![Logging Example](https://raw.githubusercontent.com/Syssos/Go_Shell/main/settings/images/Error_Log_Example.PNG)

## Travis

Travis-CI is a bonus feature of this repository. Due to the nature of this project the code can change at a fast pace. To prevent errors from occuring to someone who clones the repo, the travis build state is indicated at the top of this readme. I have multiple machines and need to share code between them. Travis allows me to do that and ensure the code is working at the same time.
