# Компиляция и запуск golang приложения

# Имя исполняемого файла
APP_NAME = sportnotes

# Пути к исходным файлам
SRC_PATH = ./cmd
MAIN_FILE = $(SRC_PATH)/main.go

# Компиляция и запуск
build:
    go build -o $(APP_NAME) $(MAIN_FILE)

run:
    # go run $(MAIN_FILE)
	go run cmd/main.go

# Очистка собранных файлов
clean:
    rm -f $(APP_NAME)
