package auction

import (
	"context"
	"fmt"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/katianemiranda/leilao/internal/entity/auction_entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestAuctionAutomaticCloseLoad(t *testing.T) {

	os.Setenv("AUCTION_INTERVAL", "3s")

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("Erro conectando no Mongo: %v", err)
	}

	db := client.Database("auction_test")

	repository := NewAuctionRepository(db)

	totalAuctions := 100
	var wg sync.WaitGroup

	// Criar vários leilões simultaneamente
	for i := 0; i < totalAuctions; i++ {

		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			auction := &auction_entity.Auction{
				Id:          fmt.Sprintf("auction-test-%d", i),
				ProductName: "Produto Teste",
				Category:    "Teste",
				Description: "Teste de carga",
				Condition:   auction_entity.New,
				Status:      auction_entity.Active,
				Timestamp:   time.Now(),
			}

			repository.CreateAuction(ctx, auction)

		}(i)
	}

	wg.Wait()

	// Esperar fechamento automático
	time.Sleep(5 * time.Second)

	// Verificar quantos fecharam
	filter := bson.M{
		"status": auction_entity.Completed,
	}

	count, err := repository.Collection.CountDocuments(ctx, filter)
	if err != nil {
		t.Fatalf("Erro contando leilões: %v", err)
	}

	if count != int64(totalAuctions) {
		t.Fatalf("Esperado %d leilões fechados, mas encontrou %d", totalAuctions, count)
	}

	t.Logf("Todos os %d leilões fecharam automaticamente com sucesso", totalAuctions)
}
