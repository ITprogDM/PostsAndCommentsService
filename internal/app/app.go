package app

import (
	"OdinVOdin/graph"
	"OdinVOdin/internal/config"
	"OdinVOdin/internal/constants"
	"OdinVOdin/internal/graphql"
	"OdinVOdin/internal/mode"
	"OdinVOdin/internal/mode/inmemory"
	"OdinVOdin/internal/mode/postgres"
	"OdinVOdin/internal/service"
	"OdinVOdin/pkg/logger"
	db "OdinVOdin/pkg/postgres"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/lib/pq"
	"net/http"
	"os"
)

func Run() {
	logger := logger.NewLogger()
	logger.Info("Executing InitLogger.")

	envFile := ".env"
	if len(os.Args) >= 2 {
		envFile = os.Args[1]
	}

	logger.Info("Executing InitConfig.")
	logger.Info("Reading %s \n", envFile)
	if err := config.InitConfig(envFile); err != nil {
		logger.Fatalf(err.Error())
	}

	logger.Info("Connecting to Postgres.")

	configs := db.PostgresConfigs{
		Name:     os.Getenv("POSTGRES_DBNAME"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Host:     os.Getenv("POSTGRES_HOST"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	logger.Info(configs)

	postgresDb, err := db.NewPostgresDB(configs)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	var modes *mode.Modes

	logger.Info("Creating Modes")
	logger.Info("IN_MEMORY = ", os.Getenv("IN_MEMORY"))

	if os.Getenv("IN_MEMORY") == "true" {
		posts := inmemory.NewPostsInMemory(constants.PostsPullSize)
		comments := inmemory.NewCommentsInMemory(constants.CommentsPullSize)
		modes = mode.NewModes(posts, comments)
	} else {
		posts := postgres.NewPostsPostgres(postgresDb)
		comments := postgres.NewCommentsPostgres(postgresDb)
		modes = mode.NewModes(posts, comments)
	}

	logger.Info("Creating Services.")
	services := service.NewServices(modes, logger)

	logger.Info("Creating graphql server.")
	port := os.Getenv("PORT")
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graphql.Resolver{
		PostsService:      services.Posts,
		CommentsService:   services.Comments,
		CommentsObservers: graphql.NewCommentsObserver(),
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	logger.Infof("Connect to http://localhost:%s/ for GraphQL playground", port)
	logger.Fatal(http.ListenAndServe(":"+port, nil))

}
