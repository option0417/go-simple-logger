package main

type LogConfig struct {
	Level       LogLevel `json:"level"`
	FilePrefix  string   `json:"file_prefix"`
	MaxFileSize int      `json:"max_file_size"`
}
