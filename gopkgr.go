package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"time"

	. "github.com/cipheras/gohelper"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()
	Flog()
	Cwindows()
	pkr()
	fmt.Printf("\nPress any key to exit...")
	fmt.Scanf("enter")
}

func pkr() {
	f, err := os.OpenFile("pkg.go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	Try(err, true, "Creating \"pkg.go\" file")
	fmt.Fprintln(f, p)
	var bytstr []string
	fmt.Fprintf(f, "pth := []string{") //open pth
	ignore := map[string]bool{
		"log.txt":                                true,
		"pkg.go":                                 true,
		"go.mod":                                 true,
		"go.sum":                                 true,
		"LICENSE":                                true,
		"README.md":                              true,
		"CHANGELOG":                              true,
		strings.ReplaceAll(os.Args[0], "./", ""): true,
	}
	Cprint(N, "Reading dir structure")
	time.Sleep(1 * time.Second)
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		exe, _ := regexp.MatchString("^(.*\\.exe)$", info.Name())
		if info.Mode().IsRegular() && !ignore[info.Name()] && !exe && !strings.Contains(path, ".git") {
			filebyt, err := ioutil.ReadFile(path)
			Try(err, false, "Reading file \""+path+"\"")
			fmt.Println(path)
			fmt.Fprintf(f, "\"%v\",", path)
			replacer := strings.NewReplacer("[", "{", "]", "}", " ", ",")
			bytstr = append(bytstr, replacer.Replace(fmt.Sprint(filebyt)))
		}
		return nil
	})
	Try(err, false)
	fmt.Fprintln(f, "}")                //close pth
	fmt.Fprintf(f, "file := [][]byte{") //open byt
	Cprint(N, "Starting packing")
	time.Sleep(1 * time.Second)
	// replacer := strings.NewReplacer("[", "", "]", "", " ", "")
	for i, v := range bytstr {
		// var ss []string
		// for i, c := range v {
		// 	if i%1000 == 0 && i > 0 && string(c) != "," && string(c) != "0" && string(c) != "{" && string(c) != "}" {
		// 		ss = append(ss, "+ \n")
		// 	}
		// 	ss = append(ss, string(c))
		// }
		// fmt.Fprintf(f, "%v,", replacer.Replace(fmt.Sprint(ss)))
		fmt.Fprintf(f, "%v,", v)
		fmt.Printf("\033[100D \033[48;5;22m:Packing [%v/%v]\033[0m [%v%v]", i+1, len(bytstr), strings.Repeat("#", i+1), strings.Repeat(".", len(bytstr)-i-1))
	}
	fmt.Fprintln(f, "}")                //close byt
	fmt.Fprintln(f, "return pth, file") //return all paths
	fmt.Fprintln(f, "}")                //end main
	fmt.Fprintln(f, u)
	f.Close()
	fmt.Printf("\n")
	Cprint(N, "Packing complete")
}

var p string = `
package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)
func pkg() ([]string, [][]byte) {
`
var u string = `
func unpkr(unpdir string) error{
	pth, file := pkg()
	for i, p := range pth {
		fd := strings.SplitAfter(p, "/")
		var fp string
		for _, v := range fd[:len(fd)-1] {
			fp = fp + v
		}
		err := os.MkdirAll(filepath.Join(unpdir, fp), os.ModePerm)
		if err != nil { 
			return err
		}
		err = ioutil.WriteFile(filepath.Join(unpdir, p), file[i], os.ModePerm)
		if err != nil {
			return err
		}
		// defer os.RemoveAll(unpdir)
		fmt.Printf("\033[1000D \033[48;5;22m:Unpacking [%v/%v]\033[0m [%v%v]", i+1, len(pth), strings.Repeat("#", i+1), strings.Repeat(".", len(pth)-i-1))
	}
	fmt.Printf("\n")
	return nil
}
`
