package main

import (
	"cmp"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores/weaviate"
)

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

func main() {
	ctx := context.Background()
	geminiClient := connectGemini(ctx)
	wvStore := connectWeaviate(ctx, geminiClient)

	server := &ragServer{
		ctx:          ctx,
		wvStore:      wvStore,
		geminiClient: geminiClient,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("POST /add/", server.addDocumentsHandler)
	mux.HandleFunc("POST /query/", server.queryHandler)

	port := cmp.Or(os.Getenv("SERVERPORT"), "9020")
	address := "localhost:" + port
	log.Println("listening on", address)
	log.Fatal(http.ListenAndServe(address, mux))
}

func connectGemini(ctx context.Context) *googleai.GoogleAI {

	apiKey := os.Getenv("GEMINI_API_KEY")

	geminiClient, err := googleai.New(ctx,
		googleai.WithAPIKey(apiKey),
		googleai.WithDefaultEmbeddingModel("text-embedding-004"))

	if err != nil {
		log.Fatal(err)
	}
	return geminiClient
}

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

func connectOpenAI(ctx context.Context) *openai.LLM {
	apiKey := os.Getenv("GEMINI_API_KEY")
	openaiClient, err := openai.New(openai.WithToken(apiKey))

	if err != nil {
		log.Fatal(err)
	}
	return openaiClient
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
