## Building LLM-Powered Applications in Go

This project demonstrates how to build a simple Retrieval Augmented Generation (RAG) server in Go using:

- **LangChainGo** – to interact with LLMs and embeddings
- **Google Gemini** – as the LLM and embedding provider
- **Weaviate** – as a vector database

We’ll build the system incrementally, explaining Go concepts and AI integration along the way.

1. ### Prerequisites

- **Go 1.23+** installed
- **Docker** (to run Weaviate)
- A **Google AI** API Key with Gemini access → [Get API Key](https://ai.google.dev)

Set your key as an environment variable:
```bash
export GEMINI_API_KEY="your_api_key_here"
```

2. ### Run Weaviate Locally

We’ll use Weaviate as the vector database. This project contains a `docker-compose.yml` config file for Weaviate.

```bash
# download and run Weaviate as container
docker compose up -d
```

3. ### Project Setup

Create a new Go module:

```bash
mkdir llm-rag-go && cd llm-rag-go
go mod init llmrag
```

Install dependencies:

```bash
go get github.com/tmc/langchaingo/
```

4. ### Add Helper Functions (`json.go`)
We’ll need helpers to read/write JSON for our API.

```go
package main

import (
	"encoding/json"
	"fmt"
	"mime"
	"net/http"
)

// Read JSON request body into a Go struct
func readRequestJSON(req *http.Request, target any) error {
	contentType := req.Header.Get("Content-Type")
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return err
	}
	if mediaType != "application/json" {
		return fmt.Errorf("expect application/json Content-Type, got %s", mediaType)
	}

	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()
	return dec.Decode(target)
}

// Render response as JSON
func renderJSON(w http.ResponseWriter, v any) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

```

5. ### Connecting to Gemini (LLM)

Create `main.go` and add function to connect Gemini

```go
import (
	"context"
	"log"
	"os"

	"github.com/tmc/langchaingo/llms/googleai"
)

const embeddingModelName = "text-embedding-004"

func connectGemini(ctx context.Context) *googleai.GoogleAI {
	apiKey := os.Getenv("GEMINI_API_KEY")
	// geminiClient, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
  geminiClient, err := googleai.New(ctx,
		googleai.WithAPIKey(apiKey),
		googleai.WithDefaultEmbeddingModel(embeddingModelName))

	if err != nil {
		log.Fatal(err)
	}
	return geminiClient
}
```

6. ### Adding Weaviate (Vector Store)

```go
import (
	"cmp"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

func connectWeaviate(ctx context.Context, geminiClient *googleai.GoogleAI) weaviate.Store {
	emb, err := embeddings.NewEmbedder(geminiClient)
	if err != nil {
		log.Fatal(err)
	}
	wvStore, err := weaviate.New(
		weaviate.WithEmbedder(emb),
		weaviate.WithScheme("http"),
		weaviate.WithHost("localhost:"+cmp.Or(os.Getenv("WVPORT"), "9035")),
		weaviate.WithIndexName("Document"),
	)
	if err != nil {
		log.Fatal(err)
	}
	return wvStore
}

```

7. ### HTTP Server and Add Documents Endpoint
We’ll now create an HTTP server that exposes two endpoints:

- `POST /add/` → Add documents to Weaviate
- `POST /query/` → Ask questions, answered using Gemini + context


```go
server := &ragServer{
  ctx:          ctx,
  wvStore:      wvStore,
  geminiClient: geminiClient,
}

func main() {
  mux := http.NewServeMux()
	mux.HandleFunc("POST /add/", server.addDocumentsHandler)

	port := cmp.Or(os.Getenv("SERVERPORT"), "9020")
	address := "localhost:" + port
	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, mux))
}

type ragServer struct {
	ctx          context.Context
	wvStore      weaviate.Store
	geminiClient *googleai.GoogleAI
}

func (rs *ragServer) addDocumentsHandler(w http.ResponseWriter, req *http.Request) {
	type document struct{ Text string }
	type addRequest struct{ Documents []document }
	ar := &addRequest{}

	if err := readRequestJSON(req, ar); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var wvDocs []schema.Document
	for _, doc := range ar.Documents {
		wvDocs = append(wvDocs, schema.Document{PageContent: doc.Text})
	}
	_, err := rs.wvStore.AddDocuments(rs.ctx, wvDocs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

```

8. ### Query Endpoint (RAG)
This will Search relevant docs & use Gemini for answers.

Update `main.go`

```go

const generativeModelName = "gemini-1.5-flash"

func main() {
  ...
	mux.HandleFunc("POST /query/", server.queryHandler)
  ...
}

func (rs *ragServer) queryHandler(w http.ResponseWriter, req *http.Request) {
	type queryRequest struct{ Content string }
	qr := &queryRequest{}
	if err := readRequestJSON(req, qr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	docs, err := rs.wvStore.SimilaritySearch(rs.ctx, qr.Content, 3)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var docsContents []string
	for _, doc := range docs {
		docsContents = append(docsContents, doc.PageContent)
	}

	ragQuery := fmt.Sprintf(ragTemplateStr, qr.Content, strings.Join(docsContents, "\n"))
	respText, err := llms.GenerateFromSinglePrompt(rs.ctx, rs.geminiClient, ragQuery, llms.WithModel("gemini-1.5-flash"))
	if err != nil {
		http.Error(w, "generative model error", http.StatusInternalServerError)
		return
	}

	renderJSON(w, respText)
}

const ragTemplateStr = `
I will ask you a question and will provide some additional context information.
Assume this context information is factual and correct, as part of internal
documentation.
If the question relates to the context, answer it using the context.
If the question does not relate to the context, answer it as normal.

For example, let's say the context has nothing in it about tropical flowers;
then if I ask you about tropical flowers, just answer what you know about them
without referring to the context.

For example, if the context does mention minerology and I ask you about that,
provide information from the context along with general knowledge.

Question:
%s

Context:
%s
`

```

9. ### Run the Server

```bash
go run .
```
You should see:

```nginx 
listening on localhost:9020
```

10. ### Test the Endpoints
The `tests` folders contains the scripts to `add documents` and to `query`.

Query Examples:

```bash
./query.sh "what is gen ai"
./query.sh "what is TDXIRV"
./query.sh "where can I find information about fuel capacity"
./query.sh "fuel saving port"
./query.sh "show fuel savings"
```