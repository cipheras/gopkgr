package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	f, err := os.OpenFile("pkg.go", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	fmt.Fprintln(f, "package main")
	var bytstr []string
	fmt.Fprintln(f, "func pkg() ([]string, [][]byte) {") //main function start
	fmt.Fprintf(f, "pth := []string{")                   //open pth
	count := 0
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.Mode().IsRegular() && info.Name() != "pkgr.go" && info.Name() != "pkg.go" && info.Name() != "templates.json" {
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
			if i%200 == 0 && i > 0 && string(c) != "," && string(c) != "0" && string(c) != "{" && string(c) != "}" {
				ss = append(ss, "+ \n")
			}
			ss = append(ss, string(c))
		}
		fmt.Fprintf(f, "%v,", replacer.Replace(fmt.Sprint(ss)))
	}
	fmt.Fprintln(f, "}")                //close byt
	fmt.Fprintln(f, "return pth, file") //return all paths
	fmt.Fprintln(f, "}")                //end main
	fmt.Println("\nFiles Packed:", count)
}

/*
func unpkr() {
	pth, file := pkg()
	fmt.Println("kakak")
	for i, p := range pth {
		fd := strings.SplitAfter(p, "/")
		var fp string
		for _, v := range fd[:len(fd)-1] {
			fp = fp + v
		}
		err := os.MkdirAll(filepath.Join("tmp", fp), os.ModePerm)
		Try("", err, true)
		err = ioutil.WriteFile(filepath.Join("tmp", p), file[i], os.ModePerm)
		Try("", err, true)
		// defer os.RemoveAll(filepath.Join("tmp", p))
		defer os.RemoveAll("tmp")
		fmt.Printf("\u001b[1000D :Unpacking[%v/%v] [%v%v]", i+1, len(pth), strings.Repeat("#", i+1), strings.Repeat(".", len(pth)-i-1))
	}
}
*/
