# =============================================
# Makefile untuk monorepo 
# =============================================

# Default environment
ENV ?= dev
APP_ENV = $(ENV)
# Path root project (otomatis terdeteksi)
#PROJECT_ROOT := $(shell cd)
PROJECT_ROOT := $(CURDIR)
# Lokasi service
HR_SERVICE_DIR = services/hr-service/cmd/api
USERS_SERVICE_DIR = services/users-service/cmd/api
MST_SERVICE_DIR = services/mst-service/cmd/api
# =============================================
# Perintah utama
# =============================================
.PHONY: run-hr run-users run-mst run-all watch-hr tidy build-hr help
# Jalankan hr-service
run-hr:
@echo "Menjalankan hr-service [env=$(APP_ENV)] dari root:
$(PROJECT_ROOT)"
cd $(HR_SERVICE_DIR) && \
set PROJECT_ROOT=$(PROJECT_ROOT) && \
set APP_ENV=$(APP_ENV) && \
go run main.go
# Hot-reload hr-service (pakai air - install dulu: go install
github.com/air-verse/air@latest)
watch-hr:
@echo "Hot-reload hr-service [env=$(APP_ENV)]..."
PROJECT_ROOT=$(PROJECT_ROOT) \
APP_ENV=$(APP_ENV) \
cd services/hr-service && air -c .air.toml
# Jalankan users-service
run-users:
@echo "Menjalankan user-service [env=$(APP_ENV)] dari root:
$(PROJECT_ROOT)"
cd $(HR_SERVICE_DIR) && \
set PROJECT_ROOT=$(PROJECT_ROOT) && \
set APP_ENV=$(APP_ENV) && \
go run main.go
# Jalankan mst-service
run-mst:
@echo "Menjalankan mst-service [env=$(APP_ENV)] dari root:
$(PROJECT_ROOT)"
cd $(HR_SERVICE_DIR) && \
set PROJECT_ROOT=$(PROJECT_ROOT) && \
set APP_ENV=$(APP_ENV) && \
go run main.go

# Jalankan semua service secara paralel (background)
run-all:
@echo "Menjalankan semua service di background [env=$(APP_ENV)]..."
@$(MAKE) run-hr & 
@$(MAKE) run-users & 
@$(MAKE) run-mst & 
wait
# Update dependencies di semua module
tidy:
@echo "Menjalankan go mod tidy di semua service..."
@cd services/hr-service && go mod tidy
@cd services/users-service && go mod tidy
@cd services/mst-service && go mod tidy
@go work sync
# Build binary hr-service (contoh)
build-hr:
@echo "Build hr-service..."
cd services/hr-service/cmd && \
CGO_ENABLED=0 GOOS=linux go build -o ../../bin/hr-service main.go
# Help
help:
@echo "Makefile commands:"
@echo " make run-hr : Jalankan hr-service (default env=dev)"
@echo " make run-hr ENV=prod : Jalankan dengan env=prod"
@echo " make watch-hr : Hot-reload hr-service"
@echo " make run-all : Jalankan semua service paralel"
@echo " make tidy : go mod tidy semua service"
@echo " make build-hr : Build binary hr-service"
@echo ""
@echo "Environment variables yang bisa di-set:"
@echo " ENV=dev|test|prod (default: dev)"