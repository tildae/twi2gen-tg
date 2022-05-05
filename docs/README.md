# Содержание
- [**_Функция Beat_**](#функция-beat)
- - [**Метод Config**](#метод-config)
- - - [_Типы Config_](#config-types)
- - [**Метод Beat**](#метод-beat)
- - - [_Типы Beat_](#beat-types)
- - - [*Подключение WebHook*](#webhook)

# Функция `Beat`
Это функция, возвращающая структуру для дальнейшей работы с методами.
| Аргумент             | Тип                | Значение  |
|:--------------------:|:------------------:|:---------:|
| data (необязательно) | interface{}        |           |

```go

type Bot struct{
  Token string
}

func Beat(data... interface{}) *Bot {
  return &Bot{}
}
```

### Использование
```go
bot := Bot()
```

# Метод `Config`
Это метод, позволяющий менять настройки бота
| Аргумент        | Тип                |
|:---------------:|:------------------:|
| typ             | string             |
| params          | interface{}        |

```go

func (bot *Bot) Config(typ string, params interface{}) {
	bot.Token = fmt.Sprintf("%v", params)
}
```

### Использование
```go
bot.Config("token", "your_token")
```

## `Config` Types
| Название        | Описание           | Значение `params` |
|:---------------:|:------------------:|:------------------:
| token           | Токен сообщества   | string            |

```go
// Пока что не доделал
```

# Метод `Beat`
Это метод, используемый для запуска бота. Принимает одно или два аргумента typ - режим работы бота `longpoll` или `webhook` и params, строго указывающийся только при использовании WebHook и имеющий структуру twillight.WebHook{}
| Аргумент        | Тип                |
|:--------------- |:------------------:|
| typ             | string             |
| params          | interface{}        |

```go
func (bot *Bot) Beat(typ string, params interface{}) {
	event := Events[typ]
	if event == nil {
		fmt.Println("Ошибка в Beat типе «" + typ + "», сообщение › " + "Тип не найден")
		return
	}
  processed := event.Process(&bot, params)
	if processed.Type == "error" {
		fmt.Println("Ошибка в Beat типа «" + typ + "», сообщение › " + processed.Message)
	} else {
		fmt.Println("Beat типа «" + typ + "» выполнен, сообщение › " + processed.Message)
	}
}
```

### Использование
```go
bot.Beat("webhook", twillight.WebHook{Package: "fiber", URL: "https://site.com",})
```

## `Beat` Types
| Название        | Описание                                           | Значение `params`   |
|:---------------:|:--------------------------------------------------:|:--------------------:
| webhook         | Подключение сервера к вебхуку                      | [Структура WebHook](#структура-webhook) |
| longpoll        | Запросы к Telegram и получение новых событий       |                     |

```go
type Event interface{
	Process(bot *Bot, params interface{}) Processed
}

var Events = map[string]Event{
	"webhook":	WebHookEvent{},
	"webhooks":	WebHookEvent{},
	"callback":	WebHookEvent{},
	"longpoll":	LongPollEvent{},
}
```

### Структура WebHook
| Параметр        | Тип                | Описание                  | Значение                |
|:--------------- |:------------------:|:-------------------------:|:------------------------:
| Package         | string             | Изменение библиотеки      | fasthttp, fiber, http   |
| URL             | string             | Ссылка для приёма событий | _https://site.com_      |

```go
type WebHook struct{
	Package	string
	URL string
}
```
