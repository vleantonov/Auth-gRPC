# Сервис авторизации (gRPC)
## Описание

Сервис предназначен для регистрации и авторизации пользователей. 
Сервис предоставляет следующий интерфейс gRPC:

1) Регистрация - позволяет пользователям создавать общий аккаунт в системе.
2) Авторизация - по данным логина и пароля генерирует JWT токен, подписанный секретным ключом определенного приложения.

[Protobuf](https://github.com/vleantonov/Auth-gRPC/blob/master/pkg/protos/proto/sso/sso.proto) сервиса

## Запуск

1) Настройка конфигурации [local](https://github.com/vleantonov/Auth-gRPC/blob/master/config)
2) Запуск миграции

```shell
make migrate
```
3) Запуск тестов

```shell
make test
```

4) Сборка сервиса

```shell
make build
```

5) Запуск сервиса

```shell
make run
```
