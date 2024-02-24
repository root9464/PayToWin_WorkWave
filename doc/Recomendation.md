# Общие рекомендации

## Архитектура
> Используемая архитектура на backend - [микросервисная](https://www.evogeek.ru/articles/765795/). Соответственно все мк -> "микросервисы" должны быть построенны так чтобы в будущем их было удобно контейнеризировать (docker).

> Примечание, вот так должна выглядеть архитектура на golang:
```
microservices/
├── users/
│   ├── main.go
│   ├── userhandler.go
│   ├── userrepository.go
│   └── user.go
├── products/
│   ├── main.go
│   ├── producthandler.go
│   ├── productrepository.go
│   └── product.go
└── orders/
    ├── main.go
    ├── orderhandler.go
    ├── orderrepository.go
    └── order.go
```
> эта же архитектура токо на NestJs:

```
microservices/
├── users/
│   ├── src/
│   │   ├── users.module.ts
│   │   ├── users.controller.ts
│   │   ├── users.service.ts
│   │   ├── user.entity.ts
│   │   ├── user.repository.ts
│   ├── test/
│   ├── package.json
│   └── tsconfig.json
├── products/
│   ├── src/
│   │   ├── products.module.ts
│   │   ├── products.controller.ts
│   │   ├── products.service.ts
│   │   ├── product.entity.ts
│   │   ├── product.repository.ts
│   ├── test/
│   ├── package.json
│   └── tsconfig.json
└── orders/
    ├── src/
    │   ├── orders.module.ts
    │   ├── orders.controller.ts
    │   ├── orders.service.ts
    │   ├── order.entity.ts
    │   ├── order.repository.ts
    ├── test/
    ├── package.json
    └── tsconfig.json
```
### Рекомендации по использованию: 
> 1. Ясное определение границ сервисов: Определите четкие границы между микросервисами, чтобы избежать излишней сложности и проблем интеграции.
> 2. Управление конфигурацией: Используйте управление конфигурацией для централизованного управления конфигурацией микросервисов.
> 3. Мониторинг и логирование: Обеспечьте мониторинг и логирование для каждого микросервиса, чтобы обеспечить прозрачность и возможность отслеживания проблем.
> 4. Развертывание и масштабируемость: Используйте средства автоматического развертывания и масштабирования для обеспечения гибкости и высокой доступности микросервисов.
> 5. Тестирование и обеспечение качества: Проводите тестирование каждого микросервиса независимо, а также создавайте автоматизированные тесты для проверки взаимодействия между сервисами.

...
