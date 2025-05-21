package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	host = "https://go-vibes-6e8qcp9.svc.aped-4627-b74a.pinecone.io"
)

// Record represents a vector record for Pinecone
type Record struct {
	ID       string            `json:"id"`
	Values   []float32         `json:"values,omitempty"` // Embedding vector
	Metadata map[string]string `json:"metadata,omitempty"`
}

// TextEmbedRequest represents a request to create text embeddings
type TextEmbedRequest struct {
	Texts []string `json:"texts"`
}

// TextEmbedResponse represents a response from a text embedding request
type TextEmbedResponse struct {
	Embeddings [][]float32 `json:"embeddings"`
}

// UpsertRequest represents a request to upsert records to Pinecone
type UpsertRequest struct {
	Vectors   []Record `json:"vectors"`
	Namespace string   `json:"namespace"`
}

// IndexStats represents statistics about an index
type IndexStats struct {
	Dimension        int                       `json:"dimension"`
	IndexFullness    float64                   `json:"index_fullness"`
	Namespaces       map[string]NamespaceStats `json:"namespaces"`
	TotalVectorCount int                       `json:"total_vector_count"`
	VectorType       string                    `json:"vector_type"`
	Metric           string                    `json:"metric"`
}

// NamespaceStats represents statistics about a namespace
type NamespaceStats struct {
	VectorCount int `json:"vector_count"`
}

// DocumentRecord represents a document to be processed
type DocumentRecord struct {
	ID       string
	Content  string
	Metadata map[string]string
}

func main() {
	// Get token from environment variable
	token := os.Getenv("PINECONE_API_KEY")
	if token == "" {
		fmt.Println("Error: PINECONE_API_KEY environment variable is not set")
		return
	}

	// Parse command-line arguments
	docsPath := flag.String("path", "../docs", "Path to the markdown documentation files")
	namespace := flag.String("namespace", "docs-namespace", "Namespace to use in Pinecone")
	chunkSize := flag.Int("chunk-size", 10, "Maximum number of records to send in a single batch")
	flag.Parse()

	// Validate path
	if _, err := os.Stat(*docsPath); os.IsNotExist(err) {
		fmt.Printf("Error: The specified path '%s' does not exist\n", *docsPath)
		return
	}

	// Process markdown files
	documents, err := processMarkdownFiles(*docsPath)
	if err != nil {
		fmt.Printf("Error processing markdown files: %v\n", err)
		return
	}

	if len(documents) == 0 {
		fmt.Println("No markdown files found to process")
		return
	}

	fmt.Printf("Found %d markdown files to process\n", len(documents))

	// Upsert the records in chunks to avoid payload size issues
	for i := 0; i < len(documents); i += *chunkSize {
		end := i + *chunkSize
		if end > len(documents) {
			end = len(documents)
		}

		chunk := documents[i:end]
		fmt.Printf("Processing chunk %d to %d of %d documents\n", i+1, end, len(documents))

		// Get content from the documents in this chunk
		var texts []string
		for _, doc := range chunk {
			texts = append(texts, doc.Content)
		}

		// Get embeddings for the texts
		embeddings, err := getEmbeddings(texts)
		if err != nil {
			fmt.Printf("Error getting embeddings: %v\n", err)
			return
		}

		// Create records with embeddings
		var records []Record
		for i, doc := range chunk {
			if i < len(embeddings) {
				record := Record{
					ID:       doc.ID,
					Values:   embeddings[i],
					Metadata: doc.Metadata,
				}
				records = append(records, record)
			}
		}

		// Upsert records to Pinecone
		err = upsertRecords(records, *namespace, token)
		if err != nil {
			fmt.Printf("Error upserting records: %v\n", err)
			return
		}

		fmt.Printf("Successfully upserted %d records\n", len(records))

		// Small delay between batches
		time.Sleep(1 * time.Second)
	}

	fmt.Println("All records upserted successfully")

	// Wait a bit for indexing to complete
	time.Sleep(5 * time.Second)

	// Check index stats
	stats, err := getIndexStats(token)
	if err != nil {
		fmt.Printf("Error getting index stats: %v\n", err)
		return
	}

	fmt.Printf("Index stats: %+v\n", stats)
}

func processMarkdownFiles(rootPath string) ([]DocumentRecord, error) {
	var documents []DocumentRecord

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only process markdown files
		if !strings.HasSuffix(strings.ToLower(info.Name()), ".md") {
			return nil
		}

		// Read file content
		content, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		// Get relative path
		relPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			relPath = path
		}

		// Extract title from frontmatter or first heading, if available
		title := extractTitle(string(content))
		if title == "" {
			title = filepath.Base(path)
		}

		// Create document record
		document := DocumentRecord{
			ID:      strings.ReplaceAll(relPath, string(filepath.Separator), "-"),
			Content: string(content),
			Metadata: map[string]string{
				"path":     relPath,
				"title":    title,
				"filename": filepath.Base(path),
			},
		}

		documents = append(documents, document)
		fmt.Printf("Processed file: %s\n", relPath)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking through directory: %w", err)
	}

	return documents, nil
}

func extractTitle(content string) string {
	// Try to extract title from frontmatter
	if strings.HasPrefix(content, "---") {
		if endIdx := strings.Index(content[3:], "---"); endIdx > 0 {
			frontmatter := content[3 : endIdx+3]
			for _, line := range strings.Split(frontmatter, "\n") {
				if strings.HasPrefix(line, "title:") {
					return strings.TrimSpace(strings.TrimPrefix(line, "title:"))
				}
			}
		}
	}

	// Try to extract from first heading
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "# ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "# "))
		}
	}

	return ""
}

// Generate 1536-dimensional embeddings for the text
func getEmbeddings(texts []string) ([][]float32, error) {
	// For this example, we'll create mock embeddings with the correct dimension (1536)
	// In a real implementation, you would call an embedding service like OpenAI
	embeddings := make([][]float32, len(texts))

	for i := range texts {
		// Create a 1536-dimensional embedding filled with small values
		embedding := make([]float32, 1536)
		for j := range embedding {
			embedding[j] = 0.01 // Small constant value for demonstration
		}
		embeddings[i] = embedding
	}

	return embeddings, nil
}

func upsertRecords(records []Record, namespace, token string) error {
	upsertData := UpsertRequest{
		Vectors:   records,
		Namespace: namespace,
	}

	jsonData, err := json.Marshal(upsertData)
	if err != nil {
		return fmt.Errorf("error marshaling data: %w", err)
	}

	req, err := http.NewRequest("POST", host+"/vectors/upsert", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Api-Key", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Read the response body to see the error details
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unexpected status code: %d, and failed to read response body: %v", resp.StatusCode, err)
		}
		return fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	return nil
}

func getIndexStats(token string) (*IndexStats, error) {
	req, err := http.NewRequest("GET", host+"/describe_index_stats", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Api-Key", token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unexpected status code: %d, and failed to read response body: %v", resp.StatusCode, err)
		}
		return nil, fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(bodyBytes))
	}

	var stats IndexStats
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&stats); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &stats, nil
}
