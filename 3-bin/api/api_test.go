package api

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var testData = []byte(`{"test": "data", "value": 42}`)

func newTestClient(t *testing.T) *Client {
	t.Helper()
	godotenv.Load("../.env")
	key := os.Getenv("KEY")
	if key == "" {
		t.Fatal("KEY не найден - создай .env файл с ключом JSONBin")
	}
	return NewClient(key)
}

func TestCreate(t *testing.T) {
	client := newTestClient(t)

	id, err := client.Create(testData, "test-bin")
	if err != nil {
		t.Fatalf("Create вернул ошибку: %v", err)
	}
	if id == "" {
		t.Fatal("Create вернул пустой ID")
	}

	t.Cleanup(func() {
		client.Delete(id)
	})
}

func TestGet(t *testing.T) {
	client := newTestClient(t)

	id, err := client.Create(testData, "test-get-bin")
	if err != nil {
		t.Fatalf("Не удалось создать bin для теста: %v", err)
	}
	t.Cleanup(func() {
		client.Delete(id)
	})

	result, err := client.Get(id)
	if err != nil {
		t.Fatalf("Get вернул ошибку: %v", err)
	}
	if result == "" {
		t.Fatal("Get вернул пустой ответ")
	}
}

func TestUpdate(t *testing.T) {
	client := newTestClient(t)

	id, err := client.Create(testData, "test-update-bin")
	if err != nil {
		t.Fatalf("Не удалось создать bin для теста: %v", err)
	}
	t.Cleanup(func() {
		client.Delete(id)
	})

	updatedData := []byte(`{"test": "updated", "value": 99}`)
	err = client.Update(id, updatedData)
	if err != nil {
		t.Fatalf("Update вернул ошибку: %v", err)
	}
}

func TestDelete(t *testing.T) {
	client := newTestClient(t)

	id, err := client.Create(testData, "test-delete-bin")
	if err != nil {
		t.Fatalf("Не удалось создать bin для теста: %v", err)
	}

	err = client.Delete(id)
	if err != nil {
		t.Fatalf("Delete вернул ошибку: %v", err)
	}
}
