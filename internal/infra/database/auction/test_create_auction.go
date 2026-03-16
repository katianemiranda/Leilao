package auction

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/katianemiranda/leilao/internal/entity/auction_entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAuctionAutomaticClose(t *testing.T) {

	// Configura intervalo pequeno para o teste
	os.Setenv("AUCTION_INTERVAL", "3s")

	ctx := context.Background()

	// Conexão Mongo
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Erro conectando no Mongo: %v", err)
	}

	db := client.Database("auction_test")

	repository := NewAuctionRepository(db)

	// Criar leilão
	auction := &auction_entity.Auction{
		Id:          "auction-test-1",
		ProductName: "Notebook",
		Category:    "Eletrônicos",
		Description: "Notebook para teste",
		Condition:   auction_entity.New,
		Status:      auction_entity.Active,
		Timestamp:   time.Now(),
	}

	errInternal := repository.CreateAuction(ctx, auction)
	if errInternal != nil {
		t.Fatalf("Erro ao criar leilão")
	}

	// Espera o tempo do fechamento automático
	time.Sleep(5 * time.Second)

	// Busca leilão novamente
	filter := bson.M{"_id": auction.Id}

	var auctionMongo AuctionEntityMongo

	err = repository.Collection.FindOne(ctx, filter).Decode(&auctionMongo)
	if err != nil {
		t.Fatalf("Erro buscando leilão: %v", err)
	}

	// Verifica se fechou automaticamente
	if auctionMongo.Status != auction_entity.Completed {
		t.Fatalf("Leilão não foi fechado automaticamente")
	}

	t.Log("Leilão fechado automaticamente com sucesso")
}
