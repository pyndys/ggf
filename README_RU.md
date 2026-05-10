# ggf (Great Go Fetch)

[English](README.md) | **Русский**

Быстрая fetch-утилита для Linux-based систем, написанная на Go. Отображает: os, kernel, shell, uptime, term, wm, memory.

## Основные особенности

- **Высокая производительность:** Прямое чтение данных из `/proc`, `/etc/os-release` и переменных окружения.
- **Легкость:** Оптимизированный бинарный файл без внешних зависимостей.
- **Поддержка дистрибутивов:** Адаптивные ASCII-логотипы для популярных Linux-дистрибутивов.

## Установка и использование

### На NixOS (Flakes)

Вы можете запустить утилиту мгновенно без установки:
```sh
nix run github:pyndys/ggf
```

Для постоянной установки добавьте в `flake.nix` в раздел inputs:
```nix
ggf.url = "github:pyndys/ggf";
```

Затем добавьте пакет в `systemPackages`:
```nix
inputs.ggf.packages.${system}.default
```

### Сборка из исходников

Требуются **Go >= 1.18** и **git**.

```sh
git clone https://github.com/pyndys/ggf.git
cd ggf
go build -ldflags="-s -w" -o ggf .
```

*Флаги `-s -w` используются для удаления таблицы символов и отладочной информации, что значительно уменьшает размер бинарного файла.*

## Лицензия
Проект распространяется под лицензией [MIT](LICENSE).

## Благодарности
* ASCII-логотипы: адаптированы из проекта [pfetch](https://github.com/dylanaraps/pfetch).
