# go-tx-context

Библиотека для работы с транзакциями `sql.Tx` в бизнес-логике по принципам чистой архитектуры для Golang. Транзакция начинается внутри use-case/service слоя и "протаскивается" дальше через context.Context.

## Особенности

- Транзакция начинается внутри бизнес-логики, а не в инфраструктуре
- Передача транзакционного контекста через `context.Context`
- Минимальная зависимость от конкретных реализаций БД
- Простая интеграция с любыми слоями приложения

## Установка

```bash
go get github.com/saifutdinov/ca-ctx-tx
```

## Пример использования

```golang
func (u *Usecase) DoSomething(ctx context.Context) {

	// main context "copied" with transaction
	txCtx, err := u.BeginTx(ctx)
	if err != nil {
		panic(err)
	}
	// check err if you want!:)
	defer u.RollbackTx(txCtx)

	// inside of each method woriking our methods with transacton
	u.repository.DoSomething(txCtx)
	u.repository.DoSomethingElse(txCtx)
	u.repository.DoSomethingAgain(txCtx)
	u.repository.DoSomethingElseAgain(txCtx)
	u.repository.DoSomethingElseAgainAgain(txCtx)

	if err := u.CommitTx(txCtx); err != nil {
		panic(err)
	}
}
```

## Как это работает

- Траназакция начинается внутри метода бизнес-логики и сохраняется в новый контекст созданный из текущего.
- Дальнейшие манипуляции происходят внутри методов пакет, где проверяется контекст на наличие транзакции и выполяет запросы внутри транзакции или просто отправляет и закрывает один запрос к БД.
- Если транзакция была начата ее необходимо закоммитить или сделать `defer u.RollbackTx(txCtx)` чтобы изменения сохранились или отменились.
- Важно отметить #1. Если транзакция уже закоммичена(правки в базе), то в случаее отката по выходу из метода, и проверке нужно учесть что попытка `rollback` вернет ошибку `transacion already committed or rollbacked`.

```golang
defer func() {
	if err := u.RollbackTx(txCtx); err != nil {
		// нужно проверить, ошибка не о уже закрытой или откаченной транзакции
		panic(err)
	}
}()
```

- Важно отметить #2. Закрывать `rows.Close()` и `CommitTx(txCtx)` крайне важно так как иначе будут накапливаться открытые транзакции и подключения к базе в "мертовом" режиме. Однако транзакция живет пока живет контекст! Так что если у вас есть ограничение на время жизни контекста, или контекст умеет сам аккуратно выходить по ошибке, то транзакция последует за ним и разорвет соединение.

<!-- TODO -->
<!-- Updarade to use with ORMs -->
