package main

import (
	"net/http"
	"os"

	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	r := mux.NewRouter()

	var svcBook BookService
	svcBook = NewBookService(logger)

	// svcBook = loggingMiddleware{logger, svcBook}
	// svcBook = instrumentingMiddleware{requestCount, requestLatency, countResult, svcBook}

	CreateBookHandler := httptransport.NewServer(
		makeCreateBookEndpoint(svcBook),
		decodeCreateBookRequest,
		encodeResponse,
	)
	GetByBookIdHandler := httptransport.NewServer(
		makeGetBookByIdEndpoint(svcBook),
		decodeGetBookByIdRequest,
		encodeResponse,
	)
	DeleteBookHandler := httptransport.NewServer(
		makeDeleteBookEndpoint(svcBook),
		decodeDeleteBookRequest,
		encodeResponse,
	)
	UpdateBookHandler := httptransport.NewServer(
		makeUpdateBookendpoint(svcBook),
		decodeUpdateBookRequest,
		encodeResponse,
	)
	http.Handle("/", r)
	http.Handle("/book", CreateBookHandler)
	http.Handle("/book/update", UpdateBookHandler)
	r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
	r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")

	var svcAuthor AuthorService
	svcAuthor = NewAuthorService(logger)

	// svcAuthor = loggingMiddleware{logger, svcAuthor}
	// svcAuthor = instrumentingMiddleware{requestCount, requestLatency, countResult, svcAuthor}

	CreateAuthorHandler := httptransport.NewServer(
		makeCreateAuthorEndpoint(svcAuthor),
		decodeCreateAuthorRequest,
		encodeResponse,
	)
	GetByAuthorIdHandler := httptransport.NewServer(
		makeGetAuthorByIdEndpoint(svcAuthor),
		decodeGetAuthorByIdRequest,
		encodeResponse,
	)
	DeleteAuthorHandler := httptransport.NewServer(
		makeDeleteAuthorEndpoint(svcAuthor),
		decodeDeleteAuthorRequest,
		encodeResponse,
	)
	UpdateAuthorHandler := httptransport.NewServer(
		makeUpdateAuthorendpoint(svcAuthor),
		decodeUpdateAuthorRequest,
		encodeResponse,
	)
	http.Handle("/author", CreateAuthorHandler)
	http.Handle("/author/update", UpdateAuthorHandler)
	r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
	r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")

	var svcPublisher PublisherService
	svcPublisher = NewPublisherService(logger)

	// svcPublisher = loggingMiddleware{logger, svcPublisher}
	// svcPublisher = instrumentingMiddleware{requestCount, requestLatency, countResult, svcPublisher}

	CreatePublisherHandler := httptransport.NewServer(
		makeCreatePublisherEndpoint(svcPublisher),
		decodeCreatePublisherRequest,
		encodeResponse,
	)
	GetByPublisherIdHandler := httptransport.NewServer(
		makeGetPublisherByIdEndpoint(svcPublisher),
		decodeGetPublisherByIdRequest,
		encodeResponse,
	)
	DeletePublisherHandler := httptransport.NewServer(
		makeDeletePublisherEndpoint(svcPublisher),
		decodeDeletePublisherRequest,
		encodeResponse,
	)
	UpdatePublisherHandler := httptransport.NewServer(
		makeUpdatePublisherendpoint(svcPublisher),
		decodeUpdatePublisherRequest,
		encodeResponse,
	)

	http.Handle("/publisher", CreatePublisherHandler)
	http.Handle("/publisher/update", UpdatePublisherHandler)
	r.Handle("/publisher/{publisherid}", GetByPublisherIdHandler).Methods("GET")
	r.Handle("/publisher/{publisherid}", DeletePublisherHandler).Methods("DELETE")

	// http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
	logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}
