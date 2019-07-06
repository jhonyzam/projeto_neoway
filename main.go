package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projeto_neoway/models"
)

type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

func main() {
	db, err := models.NewDB("postgres://postgres:@postSenha123@db/neoway?sslmode=disable")

	if err != nil {
		log.Panic(err)
	}
	ctx := context.WithValue(context.Background(), "db", db)

	http.Handle("/datastore", &ContextInjector{ctx, http.HandlerFunc(dataStoresIndex)})
	http.Handle("/datastore/execute", &ContextInjector{ctx, http.HandlerFunc(dataStoresExecute)})
	http.ListenAndServe(":3000", nil)
}

//Busca os registros no BD
func dataStoresIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	db, ok := r.Context().Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "could not get database connection pool from context", 500)
		return
	}

	dataStores, err := models.AllDataStores(db)

	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	fmt.Fprintf(w, "%s", "[")

	size := (len(dataStores) - 1)
	for i, ds := range dataStores {
		js, err := json.Marshal(ds)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if size == i {
			fmt.Fprintf(w, "%s", js)
		} else {
			fmt.Fprintf(w, "%s,", js)
		}

	}
	fmt.Fprintf(w, "%s", "]")
}

//Executa processo de extracão e inserção no BD
func dataStoresExecute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	db, ok := r.Context().Value("db").(*sql.DB)
	if !ok {
		http.Error(w, "could not get database connection pool from context", 500)
		return
	}

	//Limpa a tabela no postgresql
	err1 := models.DeleteAllDataStore(db)
	if err1 != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	//Busca Arquivos do csv
	dataStores := models.GetBase()

	//Insere registros
	for _, ds := range dataStores {
		valid := (models.IsCNPJ(ds.Cpf) || models.IsCPF(ds.Cpf))
		if !valid {
			fmt.Println("Código inválido: " + ds.Cpf)
			continue
		}

		models.Clean(&ds.Cpf)
		models.Clean(&ds.StoreFrequent)
		models.Clean(&ds.StoreLast)
		err2 := models.CreateDataStore(db, ds)

		if err2 != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
	}
}
