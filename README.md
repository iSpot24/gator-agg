# gator-agg

RSS feed aggreGator🐊

A command-line tool for working with your PostgreSQL-backed Go project. 

Aggregate posts from followed RSS feeds at a time frequency indefinetly.

## 🛠 Requirements

Before you begin, make sure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.22 or later)
- [PostgreSQL](https://www.postgresql.org/download/)

Make sure your PostgreSQL server is running and accessible.

## 🚀 Installation

Install the `gator` CLI tool using `go install`:

```bash
go install github.com/iSpot24/gator-agg
```

🧪 Running the CLI
Available commands:
 - Login for existing users
    ```bash
    gator login username 
    ```
 - Register existing users
    ```bash
    gator register username
    ```
- Clear databse (demo reasons; not present in real-case scenario)
    ```bash
    gator reset
    ```
- Display existing users
    ```bash
    gator users
    ```
- Create a feed
    ```bash
    gator addfeed name url
    ```
- Display existing feeds
    ```bash
    gator feeds
    ```
- Follow a feed for aggregating
    ```bash
    gator follow url
    ```
- Unfollow a feed
    ```bash
    gator unfollow url
    ```
- Display followed feeds
    ```bash
    gator following
    ```
- Start aggregating at frequency (string time durations: 1s, 30s, 1m, 1h, 1d)
    ```bash
    gator agg time_freq_string
    ```
- Browse aggregated posts with a limit if given (default: 2)
    ```bash
    gator browse limit
    ```

✨ Possible Improvements

🧹 Add sorting and filtering to the browse command

📄 Add pagination to handle large result sets

🤹 Add concurrency to the agg command to fetch posts more frequently and efficiently

🔤 Add a search command with fuzzy matching to find posts more easily

🧭 Add a TUI (Text User Interface) to browse and view posts inside your terminal — or open them in your browser with a click!

🔐 Add an HTTP API with authentication & authorization so others can interact with the service remotely

👷 Service manager to keep the agg command running in the background and auto-restart it on failure

