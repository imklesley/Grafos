package grafo

import (
	"github.com/golang-collections/collections/stack"
	"fmt"
	"bytes"
)

type Grafo struct {
	nome   string
	lista  []*vertice
	labels map[string]*vertice
	n      int
	m      int
}

type vertice struct {
	index int
	label string
	adj   []*vertice
}

func Novo(nome string) *Grafo {
	return &Grafo{nome, make([]*vertice, 0), make(map[string]*vertice), 0, 0}
}

func (g *Grafo) Vertice(labels ... string) {
	for _, value := range labels {
		g.criarVerticeAux(value)
	}
}
func (g *Grafo) criarVerticeAux(label string) int {
	index := g.n
	// verifica se o vertice já existe
	for _, vertice := range g.lista {
		if vertice.label == label {
			return vertice.index
		}
	}
	v := &vertice{index, label, make([]*vertice, 0)}
	g.labels[label] = v
	g.lista = append(g.lista, v)
	g.n++
	return index
}

func (g *Grafo) Aresta(label1, label2 string) *Grafo {
	v1 := g.labels[label1]
	v2 := g.labels[label2]
	v1.adj = append(v1.adj, v2)
	g.m++
	return g
}

func (g *Grafo) Label(index int) string {
	return g.lista[index].label
}

func (g *Grafo) Kosaraju() [][]string {
	gt := g.Transposto("GT")
	permutacao := gt.permutacao()

	return g.encontrarComponentes(permutacao)
}

func (g *Grafo) iniciaVisitados() []bool {
	return make([]bool, g.n)
}

func (g *Grafo) permutacao() *stack.Stack {
	pilha := stack.New()
	visitados := g.iniciaVisitados()
	for index := range g.lista {
		g.permutacaoAux(index, pilha, visitados)
	}
	return pilha
}

func (g *Grafo) permutacaoAux(index int, pilha *stack.Stack, visitados []bool) {
	if visitados[index] == true {
		return
	}
	visitados[index] = true
	for _, adj := range g.lista[index].adj {
		g.permutacaoAux(adj.index, pilha, visitados)
	}
	pilha.Push(index)
}
func (g *Grafo) encontrarComponentes(permutacao *stack.Stack) [][]string {
	var index int
	visitados := g.iniciaVisitados()
	componentes := make([][]string, 0)
	contar := 0
	for permutacao.Len() > 0 {
		index = permutacao.Pop().(int)
		if visitados[index ] == true {
			continue
		}
		if contar >= len(componentes) {
			componentes = append(componentes, make([]string, 0))
		}
		g.BuscaProfundidade(index, visitados, func(i int) {
			componentes[contar] = append(componentes[contar], g.Label(i))
		})
		contar++
	}
	componentes = g.filtrarComponentes(componentes)
	return componentes
}

func (g *Grafo) filtrarComponentes(componentes [][]string, ) [][]string {
	componentesFiltrados := make([][]string, 0)
	for key, componente := range componentes {
		if len(componente) < 2 {
			componentes[key] = nil
			continue
		}
		componentesFiltrados = append(componentesFiltrados, componente)
	}
	componentes = nil
	return componentesFiltrados
}

func (g *Grafo) BuscaProfundidade(index int, visitados []bool, faz func(i int)) {
	if visitados[index] == true {
		return
	}
	visitados[index] = true
	faz(index)
	for _, adj := range g.lista[index].adj {
		if visitados[adj.index] == true {
			continue
		}
		g.BuscaProfundidade(adj.index, visitados, faz)
	}

}

func (g *Grafo) Transposto(nome string) *Grafo {
	inverso := Novo(nome)
	g.PercorreVertices(func(vertice string, _ bool) {
		inverso.Vertice(vertice)
	})
	g.PercorreArestas(func(vertice1, vertice2 string, _ bool) {
		inverso.Aresta(vertice2, vertice1)
	})
	return inverso
}

func (g *Grafo) PercorreVertices(faz func(vertice string, ehUltimo bool)) {
	for key, v1 := range g.lista {
		faz(v1.label, key+1 == len(g.lista))
	}
}

func (g *Grafo) PercorreArestas(faz func(vertice1, vertice2 string, ehUltimo bool)) {
	for _, v1 := range g.lista {
		for key, v2 := range v1.adj {
			faz(v1.label, v2.label, key+1 == len(v1.adj))
		}
	}
}
func (g *Grafo) M() int {
	return g.m
}
func (g *Grafo) N() int {
	return g.n
}

func (g *Grafo) Imprimir() []byte {
	buff := new(bytes.Buffer)
	fmt.Fprintf(buff, "%s(v,e)\n", g.nome)
	fmt.Fprintf(buff, "e = {")
	g.PercorreVertices(func(vertice string, ehUltimo bool) {
		if ehUltimo == false {
			fmt.Fprintf(buff, "%s,", vertice)
			return
		}
		fmt.Fprintf(buff, "%s", vertice)
	})

	fmt.Fprintf(buff, "}\n")
	fmt.Fprintf(buff, "v = {")
	g.PercorreArestas(func(v1, v2 string, ehUltimo bool) {
		if ehUltimo == false {
			fmt.Fprintf(buff, "(%s->%s),", v1, v2)
			return
		}
		fmt.Fprintf(buff, "(%s->%s)", v1, v2)
	})
	fmt.Fprintf(buff, "}\n")
	fmt.Fprintf(buff, "m = %d\n", g.m)
	fmt.Fprintf(buff, "n = %d\n", g.n)

	return buff.Bytes()
}

func (g *Grafo) ImprimirComponentes(componentes [][]string) []byte {
	buff := new(bytes.Buffer)
	for _, lista := range componentes {
		for key2, componente := range lista {
			if key2+1 == len(lista) {
				fmt.Fprintf(buff, "%s", componente)
				continue
			}
			fmt.Fprintf(buff, "%s -> ", componente)
		}

		fmt.Fprintln(buff)

	}
	return buff.Bytes()
}
