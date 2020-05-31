package graphs

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	dimacs  = "col"
	edge    = 'e'
	comment = 'c'
	problem = 'p'
)

var instances = [5]string{"queen6.col"}

type DimacsReader struct {
}

func (r DimacsReader) Read(debug bool) *Graph {
	if debug {
		return r.read("queen5_5.col")
	}
	files, err := ioutil.ReadDir("./instances")
	if err != nil {
		log.Fatal(err)
	}
	for i, f := range files {
		fmt.Printf("%d -- %s\n", i, f.Name())
	}
	fmt.Println("Choose one of the above instances:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	fi, _ := strconv.Atoi(in.Text())
	return r.read(files[fi].Name())
}

func (r DimacsReader) read(f string) *Graph {
	file, err := os.Open("instances/" + f)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	g := NewGraph()
	n := 0
	for scanner.Scan() {
		l := scanner.Text()
		switch l[0] {
		case edge:
			if e := strings.Fields(l[1:]); len(e) == 2 {
				v1, err := strconv.ParseUint(e[0], 10, 32)
				if err != nil {
					log.Fatal("error occurred while parsing first vertex to uint16")
				}
				v2, err := strconv.ParseUint(e[1], 10, 32)
				if err != nil {
					log.Fatal("error occurred while parsing second vertex to uint16")
				}
				g.Add(v1, v2)
				continue
			}
			log.Printf("wrong line format {n=%d, line='%s', split=%v}\n", n, l, strings.Fields(l[1:]))
		case problem:
			if p := strings.Fields(l[1:]); len(p) == 3 {
				fmt.Printf("Undirected graph (V,E) : [%s | %s]\n", p[1], p[2])
			}
		case comment:
			fmt.Println(l[1:])
		}
		n++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return g
}
