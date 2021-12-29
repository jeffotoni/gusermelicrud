# Makefile
.EXPORT_ALL_VARIABLES:	

#GO111MODULE=on
#GOPROXY=direct
#GOSUMDB=off
GOPRIVATE=github.com/jeffotoni/gusermeli

API_ENV=PROD
API_AMB=BETA
ENV_AMBI=

build:
	@echo "########## Compilando nossa API ... "
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o userwrite main.go
	@echo "buid completo..."
	@echo "\033[0;33m################ Enviando para o server #####################\033[0m"

update:
	@echo "########## Compilando nossa API ... "
	@rm -f go.*
	go mod init github.com/jeffotoni/gusermeli/userwrite
	go mod tidy
	CGO_ENABLED=0 GOOS=linux go build --trimpath -ldflags="-s -w" -o userwrite main.go
	@echo "buid update completo..."
	@echo "fim"


tests:
	go test github.com/jeffotoni/gusermeli/userwrite/controller/handler -v
	
deploy:
	@make build
	@echo "########## Compilando nossa API ... "
	sh deploy.sh
	@echo "fim"
