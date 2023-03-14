package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func main() {
	// Define a URI de conexão com o ‘cluster’ MongoDB
	uri := "mongodb://mongo1:27017,mongo2:27017,mongo3:27017/?replicaSet=myReplicaSet"

	// Define as opções de conexão
	clientOptions := options.Client().ApplyURI(uri)

	// Conecta ao cluster MongoDB
	ctx := context.Background()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verifica se a conexão está ativa
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Prereq: Cria as collections.
	fooColl := client.Database("mydb1").Collection("foo")
	barColl := client.Database("mydb2").Collection("bar")
	oldColl := client.Database("mydb3").Collection("old")

	// ‘Step’ 1: Defina o retorno de chamada que especifica a sequência de operações a serem executadas na transação
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// Importantante: Deve passar sessCtx como parâmetro Context para as operações para serem executadas na
		// transação
		if _, err := fooColl.InsertOne(sessCtx, bson.D{{"abc", 1}}); err != nil {
			return nil, err
		}
		if _, err := barColl.InsertOne(sessCtx, bson.D{{"xyz", 2}}); err != nil {
			return nil, err
		}
		if _, err := oldColl.InsertOne(sessCtx, bson.D{{"kkk", 3}}); err != nil {
			return nil, err
		}
		if _, err := barColl.DeleteOne(sessCtx, bson.D{{"xyz", 2}}); err != nil {
			return nil, err
		}
		return nil, nil
	}

	// ‘Step’ 2: Inicie uma sessão
	session, err := client.StartSession()
	if err != nil {
		log.Fatal(err)
	}
	defer session.EndSession(ctx)

	// ‘Step’ 3: Execute o retorno de chamada usando WithTransaction
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transações efetuadas com sucesso!")
}
