# GoPkgr &nbsp; ![GitHub release (latest by date)](https://img.shields.io/github/v/release/cipheras/gopkgr?style=flat-square&logo=superuser)
#### A packager module/tool to pack templates/static files with the binary and extract it at the runtime.  

![Lines of code ](https://img.shields.io/tokei/lines/github/cipheras/gopkgr?style=flat-square)
&nbsp;&nbsp;&nbsp;&nbsp;![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/cipheras/gopkgr?style=flat-square)
&nbsp;&nbsp;&nbsp;&nbsp;![GitHub All Releases](https://img.shields.io/github/downloads/cipheras/gopkgr/total?style=flat-square)
&nbsp;&nbsp;&nbsp;&nbsp;![Platform](https://img.shields.io/badge/dynamic/json?url=https://jsonkeeper.com/b/L6FD&label=platform&query=platform&style=flat-square&labelColor=grey&color=darkgreen&cacheSeconds=3600)

## Installation
You can import this as a module and start using it.
```
go get github.com/cipheras/gopkgr
```
Or you can either use a **precompiled binary** package for your architecture or you can compile **gopkgr** from source.

### Download precompiled binary
Windows | Linux
--------|-------
[win-x64](https://github.com/cipheras/gopkgr/releases/download/v1.4.0/gopkgr-win-1.4.exe) | [linux-x64](https://github.com/cipheras/gopkgr/releases/download/v1.4.0/gopkgr-linux-1.4)

For other versions or releases go to [release page](https://github.com/cipheras/gopkgr/releases).

### Compiling from source
In order to compile from source, make sure you have installed **GO** of version at least **1.15.0** (get it from [here](https://golang.org/doc/install)).
When you have GO installed, type in the following:
```
go build 
```
## Usage
Running **gopkgr** will generate a **GO** file which will contain data from all packed files.
<br>***Note**: This file can become big if there are too many static files. Don't include images.

### Using precompiled binary
Put the *binary in the same dir* in which the files which are to be packed are located and execute it.
<br>It will automatically pack all files and files that are in folders too.

Now you can use this file with your **GO** project and use **unpacker** function to unpack it at runtime.
<br>You can give path to where it should be extracted at the runtime. 
<br>Finally build your project.

Packing Example:
![example](../assets/example.gif?raw=true)

## License
**gopkgr** is made by **@cipheras** and is released under the terms of the &nbsp;![GitHub License](https://img.shields.io/github/license/cipheras/gopkgr?color=darkgreen)

## Contact &nbsp; [![Twitter](https://img.shields.io/twitter/url?style=social&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fgopkgr&label=Tweet)](https://twitter.com/intent/tweet?text=Hi:&url=https%3A%2F%2Fgithub.com%2Fcipheras%2Fgopkgr)
> Feel free to submit a bug, add features or issue a pull request.
