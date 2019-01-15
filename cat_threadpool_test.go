package elastic

import (
	"context"
	"testing"
)

func TestCatThreadPoolService_Do(t *testing.T) {
	client, err := NewClient(SetURL("http://10.8.119.68:9200"), SetSniff(false))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.CatThreadPool().Do(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", response)
}

func TestCatThreadPoolService_Fields(t *testing.T) {
	client, err := NewClient(SetURL("http://127.0.0.1:9200"), SetSniff(false))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.CatThreadPool().Fields("completed", "name").Do(context.Background())
	t.Logf("%v", response)

	if err != nil {
		t.Fatal(err)
	}
	for _, value := range response {
		_, ok := value["completed"]
		if !ok {
			t.Fatal("no completed")
		}
		_, ok = value["name"]
		if !ok {
			t.Fatal("no completed")
		}

	}
}

func TestCatThreadPoolService_ThreadPools(t *testing.T) {
	client, err := NewClient(SetURL("http://127.0.0.1:9200"), SetSniff(false))
	if err != nil {
		t.Fatal(err)
	}
	response, err := client.CatThreadPool().ThreadPools("index", "analyze").Do(context.Background())
	t.Logf("%v", response)

	if err != nil {
		t.Fatal(err)
	}
	for _, value := range response {
		_, ok := value["name"]
		if !ok {
			t.Fatal("no completed")
		}
		_, ok = value["name"]
		if !ok {
			t.Fatal("no completed")
		}

	}
}
