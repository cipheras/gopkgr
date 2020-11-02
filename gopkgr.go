package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\n[+] Exiting...")
		os.Exit(0)
	}()
	pkr()
}

func pkr() {
	f, err := os.OpenFile("pkg.go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintln(f, p)
	var bytstr []string
	// fmt.Fprintln(f, "func pkg() ([]string, [][]byte) {") //main function start
	fmt.Fprintf(f, "pth := []string{") //open pth
	count := 0
	ignore := map[string]bool{
		"pkg.go":     true,
		"go.mod":     true,
		"LICENSE":    true,
		"README.md":  true,
		"gopkgr.go":  true,
		"gopkgr":     true,
		"gopkgr.exe": true,
	}
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && !ignore[info.Name()] {
			filebyt, err := ioutil.ReadFile(path)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(path)
			fmt.Fprintf(f, "\"%v\",", path)
			replacer := strings.NewReplacer("[", "{", "]", "}", " ", ",")
			bytstr = append(bytstr, replacer.Replace(fmt.Sprint(filebyt)))
			count++
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(f, "}")                //close pth
	fmt.Fprintf(f, "file := [][]byte{") //open byt
	replacer := strings.NewReplacer("[", "", "]", "", " ", "")
	for i, v := range bytstr {
		var ss []string
		for i, c := range v {
			if i%500 == 0 && i > 0 && string(c) != "," && string(c) != "0" && string(c) != "{" && string(c) != "}" {
				ss = append(ss, "+ \n")
			}
			ss = append(ss, string(c))
		}
		fmt.Fprintf(f, "%v,", replacer.Replace(fmt.Sprint(ss)))
		fmt.Printf("\033[100D \033[48;5;22m:Packing [%v/%v]\033[0m [%v%v]", i+1, len(bytstr), strings.Repeat("#", i+1), strings.Repeat(".", len(bytstr)-i-1))
	}
	fmt.Fprintln(f, "}")                //close byt
	fmt.Fprintln(f, "return pth, file") //return all paths
	fmt.Fprintln(f, "}")                //end main
	fmt.Println("\nFiles Packed:", count)
	fmt.Fprintln(f, u)
	f.Close()
}

var p string = `
package main
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)
func pkg() ([]string, [][]byte) {
`
var u string = `
func unpkr() {
	pth, file := pkg()
	for i, p := range pth {
		fd := strings.SplitAfter(p, "/")
		var fp string
		for _, v := range fd[:len(fd)-1] {
			fp = fp + v
		}
		err := os.MkdirAll(filepath.Join("tmp", fp), os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		err = ioutil.WriteFile(filepath.Join("tmp", p), file[i], os.ModePerm)
		if err != nil {
			log.Fatalln(err)
		}
		// defer os.RemoveAll("tmp")
		fmt.Printf("\033[1000D \033[48;5;22m:Unpacking [%v/%v]\033[0m [%v%v]", i+1, len(pth), strings.Repeat("#", i+1), strings.Repeat(".", len(pth)-i-1))
	}
}
	`
