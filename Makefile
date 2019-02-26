GOCMD=go
GOGET=$(GOCMD) get

.PHONY: deps
deps:
	$(GOGET) -u github.com/spf13/cobra
	$(GOGET) -u github.com/spf13/viper
	$(GOGET) -u github.com/olekukonko/tablewriter
