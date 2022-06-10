package db

type configModel struct {
	mongoUri    string
	mongoDb     string
	tokenSecret string
	tokenExp    string
	serveUri    string
}

var config = configModel{
	mongoUri:    "mongodb://localhost:27017",
	mongoDb:     "neowayBD", // DB name
	tokenSecret: "secret",   // Secret to use in Tokens
	tokenExp:    "1h",       // Expiration of Token
	serveUri:    ":4444",    // Serve
}
