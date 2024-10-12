# Decentralized Food Distribution App

This is an attempt to create a decentralized application for distributing excess food among the community.

## Prerequisites

- Go 1.22.6 or later
- IPFS daemon running locally

## Setup and Running

1. Clone the repository:
   ```
   git clone https://github.com/yourusername/decentralized_food_distribution_app.git
   cd decentralized_food_distribution_app
   ```

2. Install dependencies:
   ```
   cd backend
   go mod download
   ```

3. Ensure your IPFS daemon is running:
   ```
   ipfs daemon
   ```

4. Start the application:
   ```
   go run server.go
   ```

5. Open your web browser and navigate to `http://localhost:8080`

## Features

- Interactive map for creating location-based chat rooms
- P2P chat functionality using IPFS
- Real-time updates using HTMX

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
