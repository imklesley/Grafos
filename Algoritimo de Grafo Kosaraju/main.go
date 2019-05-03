package main

import (
	"Grafos-em-Golang/grafo"
	"bytes"
	"fmt"
	"os"
)

func main() {
	g := grafo.Novo("G")
	buff := new(bytes.Buffer)
	g.Vertice("a", "b", "c", "d", "e", "f", "h", "g")
	g.Aresta("a", "c")
	g.Aresta("c", "b")
	g.Aresta("b", "a")
	g.Aresta("c", "d")
	g.Aresta("d", "e")
	g.Aresta("e", "f")
	g.Aresta("f", "d")
	g.Aresta("h", "g")
	buff.Write(g.Imprimir())
	buff.Write(g.Transposto("GT").Imprimir())

	componentes := g.Kosaraju()
	fmt.Fprintln(buff, "Componentes fortemente conexos")
	buff.Write(g.ImprimirComponentes(componentes))

	buff.WriteTo(os.Stdout)
}
