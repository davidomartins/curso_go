package main

import (
	"curso_go/loja/produto"
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pg"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := ConectaBancodeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := ConectaBancodeDados()

	selectDeTodososProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := produto.Produto{}

	produtos := []produto.Produto{}

	for selectDeTodososProdutos.Next() {
		err := selectDeTodososProdutos.Scan(p.Id, p.Nome, p.Descricao, p.Preco, p.Quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtos = append(produtos, p)
	}

	// produtos := []produto.Produto{{"Camiseta", "Bem bonita", 29, 10}, {"Notebook", "Muito rápido", 1999, 2}, {"Fone", "Muito legal", 89, 2}}
	temp.ExecuteTemplate(w, "Index", produtos)
}

func ConectaBancodeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=@dmartins01 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	} else {
		return db
	}

}
