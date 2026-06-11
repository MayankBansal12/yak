# yak — a CLI Todo Manager
> Stop yak shaving. Just yak it.

---

## Commands

### Add
| Command | Description |
|---|---|
| `yak add "<title>"` | Add a new todo with just a title |
| `yak add "<title>" -p <priority>` | Add a todo with priority (high / medium / low) |
| `yak add "<title>" -d "<details>"` | Add a todo with a description |
| `yak add "<title>" -t <DD/MM>` | Add a todo with a deadline (current year assumed) |
| `yak add "<title>" -t <DD/MM/YY>` | Add a todo with a full deadline |
| `yak add "<title>" -pdt <priority> "<details>" <DD/MM>` | Add a todo with priority, details and deadline in one shot |
| `yak add -i` | Add a todo interactively (guided prompts) |

---

### View & List
| Command | Description |
|---|---|
| `yak list` | List all pending todos |
| `yak list -n <n>` | List last n todos including completed ones |
| `yak list -t <DD/MM>` | List all todos assigned to a specific date |
| `yak list -d` | List todos in detailed format (title, description, priority, deadline, created at) |
| `yak list #<ref>` | Show full details for a specific todo by ref number |
| `yak today` | List 3 todos for today sorted by priority — shows empty message if none |
| `yak today -a` | List all todos for today sorted by priority — shows empty message if none |
| `yak today -n <n>` | List n todos for today sorted by priority — shows empty message if none |
| `yak next` | Show 3 todos sorted by nearest deadline then priority — pulls from any date if today is clear |

- note: `view` is an alias for list (eg: yak view - list all pending todos)
---

### Update
| Command | Description |
|---|---|
| `yak update #<ref> "<title>"` | Update the title of a todo |
| `yak update #<ref> -p <priority>` | Update the priority of a todo |
| `yak update #<ref> -d "<details>"` | Update the description of a todo |
| `yak update #<ref> -t <DD/MM>` | Update the deadline of a todo |
| `yak update #<ref> -i` | Update a todo interactively (guided prompts) |

---

### Mark / Done
| Command | Description |
|---|---|
| `yak mark #<ref>` | Mark a todo as completed |
| `yak done #<ref>` | Alias for `yak mark` — same behaviour |

> Errors: shows a message if the ref doesn't exist or is already marked.

---

### Delete
| Command | Description |
|---|---|
| `yak delete #<ref>` | Delete a todo (prompts for confirmation) |
| `yak delete #<ref> -y` | Delete a todo and skip confirmation prompt |

> Errors: shows a message if the ref doesn't exist.

---

## Data Model

| Field | Type | Notes |
|---|---|---|
| `ref` | number | Auto-incremented, stable (never re-numbered after delete) |
| `title` | string | Required |
| `details` | string | Optional description |
| `priority` | string | `high` / `medium` / `low` — defaults to `medium` |
| `deadline` | date | Optional, format DD/MM or DD/MM/YY |
| `status` | string | `pending` / `done` |
| `created_at` | datetime | Auto-set on creation |

---

## Sort Logic

| Command | Primary Sort | Secondary Sort |
|---|---|---|
| `yak today` | Priority | — |
| `yak next` | Nearest deadline | Priority |
| `yak list` | Created at (newest first) | — |

---

## Storage
- Single JSON file at `~/.yak/todos.json`
- Human-readable, easily backed up or grepped

---

## Installation
```bash

# via curl (coming soon)
curl -fsSL https://getyak.dev/install.sh | sh

# via pnpm or bun or npm (coming soon)
pnpm -g add yak@latest 
bun -g install yak@latest
npm -g install yak@latest

# via go install
go install github.com/<user>/yak@latest
```

---

## Next Version features
- Chain of ToDos 
- Project Management
- Web view (eg: yak web)
- Cloud login for multi device support (if it gets popular)

---

## Agent Capabilities (coming soon)
ToDo Manager for your agent tasks - view, review, edit easily, your agent never forgets and if it does, you can guide it to find.
