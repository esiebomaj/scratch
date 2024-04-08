
# Scratch: RSS Feed Aggregator

## Overview
Scratch is a lightweight, efficient RSS feed aggregator written in Go. 
It is designed to fetch and consolidate blog posts from multiple RSS feeds, providing a singular, streamlined view of updates from various sources.  
  
Perfect for individuals who want to stay informed without the hassle of checking numerous websites or for developers looking for a simple aggregation tool to integrate into their projects.

## Features
- **Concurrent Feed Fetching:** Utilizes Go's concurrency model to fetch multiple feeds simultaneously, ensuring rapid updates.
- **Authentication:** Allow users to authenticate using api key.
- **Personalization:** Users are able to follow favourite blogs and get recent post from those blogs.
- **Lightweight and Fast:** Designed with minimalism in mind, ensuring quick startup times and low resource usage.
- **Extensible:** Modular design allows for easy extension, whether you want to add new data sources, output formats, or processing features.

## Configuration
Add a `.env` file to the root of the project. See example at `.env.example`

## Installation
1. Ensure you have Go installed on your system.
2. Clone the repository:
   ```
   git clone https://github.com/esiebomaj/scratch.git
   ```
3. Navigate to the cloned directory and build the project:
   ```
   cd scratch
   go build
   ```
4. Run the aggregator:
   ```
   ./scratch
   ```

## Usage
Once configured and running, scratch will fetch updates from the specified feeds at regular intervals.  
You can adjust the fetch frequency, cuncurrency value and other behaviors by modifying the configuration in `.env` to suit your needs.

## Contributing
Contributions are welcome! Whether you're fixing bugs, adding new features, or improving the documentation, please feel free to submit a pull request.

## License
Scratch is released under the MIT License. See the LICENSE file for more details.
