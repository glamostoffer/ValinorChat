# Установка базового образа
FROM golang:1.22

# Установка рабочей директории в контейнере
WORKDIR /app

# Копирование файлов проекта в контейнер
ADD . /app

# Установка зависимостей
RUN go mod download

# Компиляция проекта
RUN go build -o main ./cmd/main.go

# Создание директории config и файла config.yaml
RUN mkdir -p ./config && echo "\
env: 'local'\n\
start_timeout: 5s\n\
stop_timeout: 5s\n\
grpc:\n\
 timeout: 5s\n\
 port: 6001\n\
 host: 'localhost'\n\
postgres:\n\
 user: 'root'\n\
 port: '5432'\n\
 dbname: 'valinor'\n\
 password: 'root'\n\
 host: 'localhost'\n\
 driver: 'pgx'\n\
 sslMode: 'disable'\n\
 max_open_conns: 10\n\
 conn_max_lifetime: 1h\n\
 max_idle_conns: 10\n\
 conn_max_idle_time: 1h\n\
" > ./config/config.yaml

# Установка переменной окружения CONFIG_PATH
ENV CONFIG_PATH=./config/config.yaml

# Запуск приложения
CMD ["./main"]