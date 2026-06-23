```
                       ░░██
 █████  ████  ██████    ░███  █████
░░███  ░███  ░░░░░███   ░███░░░███
 ░███  ░███   ███████   ░███████░
 ░███  ░███  ███░░███   ░███░  ███
 ░░███████  ░████████  ░███░  █████
  ░░░░░███  ░░░░░░░░   ░░░   ░░░░░
  ███ ░███
  ░██████
   ░░░░░                   
```

# yak — a CLI Todo Manager
> following todos are hard but managing them shouldn't be

---

## Installation

### Prerequisites
- Go 1.22 or later

### via go install
1. Install via go install
```bash
go install github.com/mayankbansal12/yak@latest
```
2. Ensure `$HOME/go/bin` is in your PATH:
```bash
export PATH="$PATH:$HOME/go/bin"
```
3. Verify using
```bash
yak --version
```

### Build From Source Directly

```bash
git clone https://github.com/mayankbansal12/yak.git
cd yak
go build -o yak .
sudo mv yak /usr/local/bin/
```


- via curl and other installation formats (coming soon)

---

## Commands

### Add
| Command | Description |
|---|---|
| `yak add "<title>"` | Add a new todo with just a title |
| `yak add "<title>" -p <priority>` | Add a todo with priority (0-2, 0 is highest) |
| `yak add "<title>" -d "<details>"` | Add a todo with a description |
| `yak add "<title>" -t <DD-MM>` | Add a todo with a deadline (current year assumed) |
| `yak add "<title>" -t <DD-MM-YY>` | Add a todo with a full deadline |
| `yak add "<title>" -j` | Add a todo and output it as JSON |
| `yak add -i` | Add a todo interactively (guided prompts) |

**Examples:**
```bash
yak add "Buy groceries"
yak add "Pay rent" -p 0 -d "Include electricity bill" -t 25-06
yak add "Read book" -t 30-06-25
yak add "Buy milk" -p 1
yak add -i
```

---

### View & List
| Command | Description |
|---|---|
| `yak list` | List all pending todos |
| `yak list -n <n>` | List last n todos including completed ones |
| `yak list -t <DD-MM>` | List all todos assigned to a specific date |
| `yak list -d` | List todos in detailed format (title, description, priority, deadline, created at) |
| `yak list <ref>` | Show full details for a specific todo by ref number |
| `yak list -c` | List todos in compact one-line format |
| `yak list -j` | Output todos as JSON |
| `yak today` | List 3 todos for today sorted by priority — shows empty message if none |
| `yak today -a` | List all todos for today sorted by priority — shows empty message if none |
| `yak today -n <n>` | List n todos for today sorted by priority — shows empty message if none |
| `yak today -c` | List today's todos in compact format |
| `yak today -j` | Output today's todos as JSON |
| `yak next` | Show 3 todos sorted by nearest deadline then priority — pulls from any date if today is clear |
| `yak next -c` | Show next todos in compact format |
| `yak next -j` | Output next todos as JSON |

- note: `view`, `see`, `get`, `find` are aliases for list (eg: yak view - list all pending todos)

**Examples:**
```bash
yak list
yak list -n 5
yak list -t 22-06
yak list -d
yak list 3
yak list -c
yak list -j
yak today
yak today -a
yak today -n 5
yak today -c
yak today -j
yak next
yak next -c
yak next -j
```

---

### Update
| Command | Description |
|---|---|
| `yak update <ref> "<title>"` | Update the title of a todo |
| `yak update <ref> -p <priority>` | Update the priority of a todo |
| `yak update <ref> -d "<details>"` | Update the description of a todo |
| `yak update <ref> -t <DD-MM>` | Update the deadline of a todo |

**Examples:**
```bash
yak update 3 "Buy organic milk"
yak update 3 -p 1
yak update 3 -d "From the local farm" -t 28-06
```

---

### Mark / Done
| Command | Description |
|---|---|
| `yak mark <ref>` | Mark a todo as completed |
| `yak done <ref>` | Alias for `yak mark` — same behaviour |

> Errors: shows a message if the ref doesn't exist or is already marked.

**Examples:**
```bash
yak mark 3
yak done 3
yak mark #5
```

---

### Delete
| Command | Description |
|---|---|
| `yak delete <ref>` | Delete a todo (prompts for confirmation) |
| `yak delete <ref> -y` | Delete a todo and skip confirmation prompt |

> Errors: shows a message if the ref doesn't exist.

**Examples:**
```bash
yak delete 3
yak delete 3 -y
yak delete #5 -y
```

---

## Data Model

| Field | Type | Notes |
|---|---|---|
| `id` | number | Auto-incremented, stable (never re-numbered after delete) |
| `title` | string | Required |
| `desc` | string | Optional description |
| `priority` | number | `0` / `1` / `2` — optional (0 is highest) |
| `due_date` | string | Optional, format DD-MM or DD-MM-YY |
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
| `yak list -c` | Due date (nearest first) | Priority |

---

## Output Formats

### Cards (default)
`yak list`, `yak today`, and `yak next` render each todo as a card. Cards show the todo's status icon, title, relative due date, priority, and status.

### Compact (`-c`)
`yak list -c`, `yak today -c`, and `yak next -c` render one todo per line, sorted by due date then priority. This is useful when you have many todos or a narrow terminal.

### JSON (`-j`)
`yak list -j`, `yak today -j`, `yak next -j`, and `yak add ... -j` output JSON for scripting and agent integration. The list commands emit an envelope:

```json
{
  "todos": [...],
  "count": 3,
  "generated_at": "2026-06-23T10:00:00Z"
}
```

`yak add -j` emits the created todo:

```json
{
  "todo": {...},
  "id": 5
}
```

### Status icons
| Icon | Meaning |
|---|---|
| `●` | Completed |
| `○` | Pending |
| `!` | Overdue (active todos only) |

### Priority bars
Priority is shown as ascending vertical bars (`▁` is the shortest, `█` is the tallest):

| Bars | Priority |
|---|---|
| `▁▄█` | High |
| `▁▄` | Medium |
| `▁` | Low |

Cards render the bars followed by the word (`▁▄█ High`); compact mode shows the bars only.

### Relative due dates
Due dates are shown relative to today when within ±7 days:

- `today`, `tomorrow`, `yesterday`
- `in N days` / `N days overdue`
- Absolute dates (`Jan 02`) for dates beyond a week, or (`Jan 02, 2027`) for a different year.

---

## Storage
- Single JSON file at `~/.local/share/yak/todos.json`
- Respects `$YAK_DATA_DIR` and `$XDG_DATA_HOME` environment variables
- Human-readable, easily backed up or grepped

---

## Next Version features
- Chain of ToDos 
- Project Management
- Web view (eg: yak web)
- Cloud login for multi device support (if it gets popular)

---

## Agent Capabilities (coming soon)
ToDo Manager for your agent tasks - view, review, edit easily, your agent never forgets and if it does, you can guide it to find.
