## gator-agg

RSS feed aggreGator🐊 

A command-line tool for working with your PostgreSQL-backed Go project. 

Aggregate posts from followed RSS feeds at a time frequency indefinetly.

## 🛠 Requirements

Before you begin, make sure you have the following installed:

- [Go](https://golang.org/doc/install) (version 1.18 or later)
- [PostgreSQL](https://www.postgresql.org/download/)

Make sure your PostgreSQL server is running and accessible.

## 🚀 Installation

Install the `gator` CLI tool using `go install`:

```bash
go install github.com/yourusername/gator/cmd/gator@latest
```

⚙️ Configuration
Before running gator, you'll need a configuration file. Create a gator.yaml file in the root of your project directory:

```yaml
database:
  host: localhost
  port: 5432
  user: postgres
  password: your_password
  dbname: your_database

output: ./generated
```

You can also pass a custom config path using the --config flag.

🧪 Running the CLI
Available commands:
 - Login existing users
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