# The Language Basics

[Home](README.md)

This is going to be a very brief tour of the language but if you want a bit more you can go to [A Tour of Go](https://tour.golang.org/)
where you will find a great intor to the language that can be done betwen 2 and 4 hours).  If you are still interested there are now
several books written on go and also many more tutorials online about a very comprehensive list of subjects that you might want to 
play with.

1. Installation
2. Path
3. IDEs
4. Variables
5. Functions
 * Multi returns
 * Named returns
6. Arrays, Slices, Structs and Maps
7. Concurrency
8. But what about OO?

### Installation

There are 2 common ways to install Go on your computer:
* An installer for your system (Windows, Linux, Mac (Homebrew)), or;
* GVM (Linux / Mac only)

I recommend GVM as it gives you the most flexibility to switch from one version to another easily.  The instructions 
are simple and you can go thru it on your own.  

* [Windows Installer](https://golang.org/doc/install)
* [Linux/Mac GVM](https://github.com/moovweb/gvm)



### Path

There are 2 very important environment variables that GO uses:

* GOROOT - Where the GO tools are installed on your system (in the case of GVM this is set for you)
* GOPATH - Where GO will download and install libraries that you want to use in your applications
   * some packages that are available for you to use in your apps [https://github.com/avelino/awesome-go](https://github.com/avelino/awesome-go)

Additionally there are 2 more that are important for cross compiling

* GOOS   - compiler/linker target OS
* GOARCH - compiler/linker target architecture

Example usage: GOOS=linux GOARCH=386 go build -o hello hello.go  (compiles for the linux/386 architecture)


### IDEs

There is a vsast array of choices from Text Editors to an Eclipse plugin for GO.  I will be using Webstorm from JetBrains
as I have found that it has a lot of nice built-in functions plus I can use all the frontend tools I want with it.

### Variables

Simple variables are declared in 2 different ways - checkout the [example here](https://play.golang.org/p/j8bXIVwWNx)

### Functions (*Multi returns* / *Named returns*)

Functions are the equivalent of methods in Java or Python except that they have some special characteristics that we will see soon.
Here is an example of a function where we extend the code we saw before [First pass](https://play.golang.org/p/u3m-cbYasN) 
and what if we wanted to use named return variables? [Second pass](https://play.golang.org/p/ZR8mZHVeum)

### Arrays, Slices, Structs and Maps

Examples:
* [Arrays](https://play.golang.org/p/l_KQY4_CzF)
* [Slices]()
* [Structs]()
* [Maps]()

### Concurrency

### But What about OO?

Simple Go is *not* and object oriented language,...  BUT...  You can assign methods on structs which can give it an 
OO flavor.  [Check it out]()


[Home](README.md)