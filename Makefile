GOCMD=go
GOGET=$(GOCMD) get

deps:
	$(GOGET) -u github.com/mitchellh/go-homedir
	$(GOGET) -u github.com/spf13/cobra
	$(GOGET) -u github.com/spf13/viper
