package graphql

type Config struct {
	MaxParallelism int `json:"max_parallelism" mapstructure:"max_parallelism"`
	MaxDepth       int `json:"max_depth" mapstructure:"max_depth"`
}
