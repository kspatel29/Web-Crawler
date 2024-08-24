# Concurrent Web Crawler

## Overview

This Concurrent Web Crawler is a Go-based application designed to efficiently crawl websites. It leverages Go's concurrency features to crawl multiple pages simultaneously, making it fast and efficient. The crawler respects the site's domain boundaries and allows users to set limits on the number of pages crawled and the level of concurrency.

## Features

- **Concurrent Crawling**: Utilizes Go's goroutines for parallel processing of web pages.
- **Configurable Limits**: Allows setting maximum number of pages to crawl and maximum concurrency level.
- **Domain Respect**: Only crawls pages within the same domain as the starting URL.
- **URL Normalization**: Ensures unique page counting by normalizing URLs.
- **Detailed Reporting**: Provides a summary of crawled pages and their visit counts.

## Prerequisites

- Go 1.2 or higher

## Installation

1. Clone the repository:
git clone https://github.com/yourusername/web-crawler.git


2. Navigate to the project directory:
cd web-crawler



## Usage

Run the crawler using the following command:
go run .



Or build and run:
go build -o crawler
./crawler



### Parameters:

- `<starting_url>`: The URL to start crawling from.
- `<max_concurrency>`: Maximum number of concurrent crawling goroutines.
- `<max_pages>`: Maximum number of pages to crawl.

### Example:
./crawler https://example.com 10 100



This will crawl `https://example.com` using a maximum of 10 concurrent goroutines and will stop after crawling 100 pages or when there are no more pages to crawl within the domain.

## Output

The crawler will print progress to the console as it crawls, showing each URL being processed. After completion, it will display a report showing:

- Total number of pages crawled
- List of all crawled URLs with their visit counts
