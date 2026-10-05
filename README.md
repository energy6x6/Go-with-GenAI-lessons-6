## Completed exercises

This branch includes both the repository exercises (`todo`, `validate`) and
`Homework6_en.md` (`config`). The written report is submitted separately.

- `todo`: JSON save/load with wrapped errors; the original specification tests remain unchanged.
- `validate`: basic email syntax and 10–15 ASCII phone digits with an optional leading `+`.
- `config`: JSON tags, save/load, port/environment validation, streaming JSON,
  injectable `FileStorage`, and handwritten mocks for filesystem failures.
- Configuration validation is explicit: call `ValidateConfig` separately from save/load.
  Passwords never enter JSON; an empty database URL is omitted.

Run from this directory:

```bash
go test -v ./...
go test -coverprofile=cover.out ./...
go tool cover -func=cover.out
go tool cover -html=cover.out
go vet ./...
```

The first implementation commit uses `MarshalIndent`/`Unmarshal`; the next
refactors configuration persistence to `Encoder`/`Decoder` while retaining its tests.
CI checks both the original exercises and total coverage above 80%.

---

# Go with GenAI — Lesson 6 homework

**Lesson 6. File I/O, JSON and Testing: a triad of production-ready code**

This repository is the starter template for the Lesson 6 homework. Each
task lives in its own package/folder with a stub to implement and a
test suite that acts as the specification.

| Folder | Task | What you implement |
|---|---|---|
| [`todo/`](todo) | 1 | `SaveTodos` / `LoadTodos` — persist a todo list to JSON, ≥80% coverage |
| [`validate/`](validate) | 2 | `ValidateEmail` / `ValidatePhone` — table-driven test with ≥8 cases |
| [`ai-blind-testing/`](ai-blind-testing) | 3 | `REPORT.md` — AI blind test generation, reviewed by your mentor |

## Getting started

1. Use this repository as a template (or clone it) into your own GitHub
   account — see **"How to get this onto your own GitHub"** below.
2. Make sure you have Go 1.22+ installed: `go version`.
3. Implement `todo/todo.go` and `validate/validate.go`, and extend the
   test tables in `validate/validate_test.go`.
4. Fill in `ai-blind-testing/REPORT.md`.
5. Run everything locally before pushing:

   ```bash
   go test -v ./...          # every test, every subtest, verbose
   go test -cover ./...      # coverage per package
   ```

   or, with the included Makefile:

   ```bash
   make test-verbose
   make cover
   ```

6. Push your changes. GitHub Actions runs automatically (see below) and
   posts a full summary in the **Actions** tab of your repo.

## How the automatic check works

`.github/workflows/check-homework.yml` runs on every push. It is built
specifically so that **one failing task never hides the others**:

- `go test` itself already runs every subtest and every package by
  default (it does *not* stop at the first failure unless you pass
  `-failfast`, which this workflow never does) — and all the test cases
  in this repo use `t.Errorf` rather than `t.Fatalf` for their actual
  assertions, so a single wrong case doesn't abort its sibling cases.
- Every workflow step for Task 1 and Task 2 uses `continue-on-error:
  true`, and every step after it runs with `if: always()`. This means
  Task 2's tests still run (and their full output is still printed)
  even if every single Task 1 test failed, and the coverage check and
  job summary still run too.
- A **Job Summary** table (visible at the top of the Actions run) gives
  you a one-glance overview per task, with a pointer to the exact log
  section for the full list of failing cases.
- A final "Final result" step looks at everything that was recorded and
  fails the overall run if anything is incomplete — so the ✗/✓ on your
  commit is still accurate, but you always see the *complete* list of
  what's left to do in a single run, not just the first problem.

To see this locally, run `go test -v ./...` and read the whole output —
every `--- FAIL:` line is a separate case you still need to handle.

## How to get this onto your own GitHub

**Option A — GitHub CLI**

```bash
gh repo create <your-username>/go-with-genai-lesson6-homework --private --source=. --remote=origin --push
```

**Option B — manual**

```bash
git init
git add .
git commit -m "Lesson 6 homework starter"
git branch -M main
git remote add origin https://github.com/<your-username>/go-with-genai-lesson6-homework.git
git push -u origin main
```

Then open the **Actions** tab on GitHub once to confirm the workflow
runs (it should show three failing tasks initially — that's expected
for an unimplemented template).

## Repository layout

```
.
├── .github/workflows/check-homework.yml   # CI: runs & reports on all 3 tasks
├── todo/
│   ├── todo.go          # Task 1 — implement this
│   ├── todo_test.go     # Task 1 — specification, do not edit
│   └── README.md
├── validate/
│   ├── validate.go        # Task 2 — implement this
│   ├── validate_test.go   # Task 2 — extend the case tables, then it's the spec
│   └── README.md
├── ai-blind-testing/
│   ├── REPORT.md   # Task 3 — fill this in
│   └── README.md
├── go.mod
├── Makefile
└── README.md   # you are here
```


# Go з GenAI — Домашнє завдання до уроку 6

**Урок 6. Робота з файлами, JSON та тестування: тріада для коду, готового до промислової експлуатації**

Цей репозиторій є шаблоном для виконання домашнього завдання до уроку 6. Кожне
завдання розміщене в окремому пакеті/папці та містить шаблон коду для реалізації
і набір тестів, що слугують специфікацією.

| Папка | Завдання | Що потрібно реалізувати |
|---|---|---|
| [`todo/`](todo) | 1 | `SaveTodos` / `LoadTodos` — збереження списку справ у JSON, покриття тестами ≥80% |
| [`validate/`](validate) | 2 | `ValidateEmail` / `ValidatePhone` — табличне тестування (table-driven test) з ≥8 випадками |
| [`ai-blind-testing/`](ai-blind-testing) | 3 | `REPORT.md` — сліпе тестування генерації коду ШІ (AI blind test), перевірка ментором |

## Початок роботи

1. Використайте цей репозиторій як шаблон (або клонуйте його) у свій власний обліковий запис
GitHub — див. розділ **«Як перенести це у власний GitHub»** нижче.
2. Переконайтеся, що у вас встановлено Go версії 1.22 або новішої: `go version`.
3. Реалізуйте `todo/todo.go` та `validate/validate.go`, а також розширте
таблиці тестів у файлі `validate/validate_test.go`.
4. Заповніть файл `ai-blind-testing/REPORT.md`.
5. Запустіть усі тести локально перед відправкою змін (push):

```bash
go test -v ./...          # усі тести, усі підтести, детальний вивід
go test -cover ./...      # покриття тестами для кожного пакета
```

або скористайтеся наявним Makefile:

```bash
make test-verbose
make cover
```

6. Надішліть (push) свої зміни. GitHub Actions запуститься автоматично (див. нижче) і
опублікує повний звіт на вкладці **Actions** вашого репозиторію.

## Як працює автоматична перевірка

Файл `.github/workflows/check-homework.yml` запускається при кожному надсиланні змін (push). Систему побудовано так, щоб **невдача в одному завданні не приховувала інші**:

- Команда `go test` за замовчуванням виконує всі підтести та пакети (вона *не* зупиняється на першій помилці, якщо не вказано прапорець `-failfast`, чого в цьому робочому процесі не робиться). Крім того, усі тестові випадки в цьому репозиторії використовують `t.Errorf` замість `t.Fatalf` для перевірок, тому один невдалий тест не перериває виконання інших тестів.
- Кожен крок робочого процесу для Завдання 1 і Завдання 2 використовує параметр `continue-on-error: true`, а всі наступні кроки виконуються з умовою `if: always()`. Це означає, що тести Завдання 2 все одно виконуються (і виводяться повні результати їхньої роботи), навіть якщо всі тести Завдання 1 зазнали невдачі; так само виконуються перевірка покриття коду тестами та формування підсумкового звіту.
- Таблиця **Job Summary** (яку видно у верхній частині сторінки виконання Actions) надає короткий огляд результатів кожного завдання та посилання на відповідний розділ журналу з повним переліком невдалих тестів.
- Кінцевий крок «Final result» аналізує всі зафіксовані результати й позначає виконання всього процесу як невдале, якщо щось залишилося невиконаним. Завдяки цьому позначка ✗ або ✓ біля вашого коміту залишається точною, але ви одразу бачите *повний* перелік того, що ще потрібно виправити, а не лише першу виявлену проблему.

Щоб побачити це локально, виконайте команду `go test -v ./...` і перегляньте весь вивід: кожен рядок, що починається з `--- FAIL:`, вказує на окремий тест, який потребує вашої уваги.

## Як перенести це у ваш власний репозиторій на GitHub

**Варіант А — GitHub CLI**

```bash
gh repo create <your-username>/go-with-genai-lesson6-homework --private --source=. --remote=origin --push
```

**Варіант Б — вручну**

```bash
git init
git add .
git commit -m "Lesson 6 homework starter"
git branch -M main
git remote add origin https://github.com/<your-username>/go-with-genai-lesson6-homework.git
git push -u origin main
```

Після цього відкрийте вкладку **Actions** на GitHub, щоб переконатися, що робочий процес (workflow)
запустився (спочатку має відобразитися три невдалі завдання — це очікувано
для шаблону, в якому ще не реалізовано код).

## Структура репозиторію

```
.
├── .github/workflows/check-homework.yml   # CI: виконання та звіт щодо всіх 3 завдань
├── todo/
│   ├── todo.go          # Завдання 1 — реалізуйте це
│   ├── todo_test.go     # Завдання 1 — специфікація, не редагуйте
│   └── README.md
├── validate/
│   ├── validate.go        # Завдання 2 — реалізуйте це
│   ├── validate_test.go   # Завдання 2 — доповніть таблиці тестових випадків (це і є специфікація)
│   └── README.md
├── ai-blind-testing/
│   ├── REPORT.md   # Завдання 3 — заповніть цей файл
│   └── README.md
├── go.mod
├── Makefile
└── README.md   # ви тут
```