# yak — a CLI Todo Manager
> following todos are hard but managing them shouldn't be

---

## Commands

### Add
| Command | Description |
|---|---|
| `yak add "<title>"` | Add a new todo with just a title |
| `yak add "<title>" -p <priority>` | Add a todo with priority (0-2, 0 is highest) |
| `yak add "<title>" -d "<details>"` | Add a todo with a description |
| `yak add "<title>" -t <DD/MM>` | Add a todo with a deadline (current year assumed) |
| `yak add "<title>" -t <DD/MM/YY>` | Add a todo with a full deadline |
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

- note: `view`, `see`, `get`, `find` are aliases for list (eg: yak view - list all pending todos)
---

### Update
| Command | Description |
|---|---|
| `yak update #<ref> "<title>"` | Update the title of a todo |
| `yak update #<ref> -p <priority>` | Update the priority of a todo |
| `yak update #<ref> -d "<details>"` | Update the description of a todo |
| `yak update #<ref> -t <DD/MM>` | Update the deadline of a todo |

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
| `id` | number | Auto-incremented, stable (never re-numbered after delete) |
| `title` | string | Required |
| `desc` | string | Optional description |
| `priority` | number | `0` / `1` / `2` — optional (0 is highest) |
| `due_date` | string | Optional, format DD/MM or DD/MM/YY |
| `created_at` | datetime | Auto-set on creation |
| `updated_at` | datetime | Auto-set on creation and updates |
| `completed_at` | datetime | Set when marked as done (empty = pending) |

---

## Sort Logic

| Command | Primary Sort | Secondary Sort |
|---|---|---|
| `yak today` | Priority | — |
| `yak next` | Nearest deadline | Priority |
| `yak list` | Updated at (newest first) | — |

---

## Storage
- Single JSON file at `~/.local/share/yak/todos.json`
- Respects `$YAK_DATA_DIR` and `$XDG_DATA_HOME` environment variables
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
