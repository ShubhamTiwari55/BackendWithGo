# RSS Feed Aggregator

An RSS feed aggregator built with Golang and PostgreSQL. This web server allows clients to:

- Add RSS feeds to be collected
- Follow and unfollow RSS feeds that other users have added
- Fetch all of the latest posts from the RSS feeds they follow

RSS feeds are a way for websites to publish updates to their content. You can use this project to keep up with your favorite blogs, news sites, podcasts, and more!

## Features

- User authentication
- Add and manage RSS feeds
- Follow/unfollow feeds added by other users
- Fetch and store latest posts from followed feeds
- RESTful API for easy integration

## Tech Stack

- **Backend:** Golang (Go)
- **Database:** PostgreSQL
- **Libraries & Tools:**
  - `gorilla/mux` for routing
  - `gorm` for ORM (optional)
  - `go-chi/chi` for middleware handling
  - `go-feed` for RSS parsing
  - `godotenv` for environment variable management

## Getting Started

### Prerequisites

- Install [Go](https://golang.org/doc/install)
- Install [PostgreSQL](https://www.postgresql.org/download/)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/ShubhamTiwari55/rss-agg.git
   cd rss-agg
   ```

2. Create a `.env` file and configure your database credentials:
   ```env
  PORT=yourPortNumber
  DB_URL=yourDbUrl
   ```

3. Install dependencies:
   ```sh
   go mod tidy
   ```

4. Run database migrations (if applicable):
   ```sh
   goose postgres DB_URL up
   ```

5. Start the server:
   ```sh
   go build && rss-agg.exe
   ```

## API Endpoints

| Method | Endpoint                           | Description |
|--------|-----------------------------------|-------------|
| GET    | `/v1/ready`                       | Check service readiness |
| GET    | `/v1/err`                         | Trigger an error for testing |
| POST   | `/v1/users`                       | Create a new user |
| GET    | `/v1/users`                       | Get user details (requires authentication) |
| POST   | `/v1/feeds`                       | Create a new RSS feed (requires authentication) |
| GET    | `/v1/feeds`                       | Get all available RSS feeds |
| POST   | `/v1/feeds_follows`               | Follow a feed (requires authentication) |
| GET    | `/v1/feeds_follows`               | Get followed feeds (requires authentication) |
| DELETE | `/v1/feeds_follows/{feedFollowID}`| Unfollow a feed (requires authentication) |
| GET    | `/v1/posts`                       | Get the latest posts from followed feeds (requires authentication) |

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License.

## Contact

For any inquiries, open an issue on GitHub.

