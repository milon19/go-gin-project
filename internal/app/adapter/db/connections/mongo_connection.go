package db

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"healthousedemo/cmd/app/config"
	"io/ioutil"
	"log"
	"time"
)

var (
	client   *mongo.Client
	database *mongo.Database
	err      error
)

const (
	caFilePath               = "configurations/db/rds-combined-ca-bundle.pem"
	connectTimeout           = 30
	connectionStringTemplate = "%s?ssl=%s&readpreference=secondaryPreferred"
)

func Connect() {
	uri := fmt.Sprintf(connectionStringTemplate, config.MongoUrl, config.MongoSSL)

	databaseName := config.DatabaseName

	if config.MongoSSL == "true" {
		uri += "&replicaSet=rs0&retryWrites=false"
		tlsConfig, err := getCustomTLSConfig(caFilePath)
		if err != nil {
			log.Fatalf("Failed getting TLS configuration: %v", err)
		}
		client, err = mongo.NewClient(options.Client().ApplyURI(uri).SetTLSConfig(tlsConfig))
	} else {
		client, err = mongo.NewClient(options.Client().ApplyURI(uri))
	}

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to cluster: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping cluster: %v", err)
	}

	log.Printf("Connected to healthumentDB!")

	database = client.Database(databaseName)
	log.Printf("Connected to database: %v", database)
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.New("failed parsing pem file")
	}
	return tlsConfig, nil
}

func GetDB(collection string) *mongo.Collection {
	col := database.Collection(collection)
	return col
}

func GetCtx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 50000*time.Second)
	return ctx, cancel
}
