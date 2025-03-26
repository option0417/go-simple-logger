package main

func main() {
	logger := NewLogger()
	defer logger.Close()

	logger.Info("This is an information message")
	logger.Debug("This is a debug message")
	logger.Error("This is an error message")

}
