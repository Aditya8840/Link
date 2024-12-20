# Link - A Scalable URL Shortener

# Link

A high-performance URL shortener microservice built with scalable architecture.

## Overview

Link is a robust URL shortening service designed with scalability and performance in mind. It transforms long URLs into short, manageable links while maintaining high availability and quick response times.

## Tech Stack

- Go (Programming Language)
- Go-fiber (Web Framework)
- MongoDB (Primary Database)
- Redis (Caching & Counter Management)

## Key Features

- Base62 encoding for generating short codes
- Redis-based counter system for URL generation
- Efficient caching mechanism with LRU policy
- Cache capacity: 1GB with LRU eviction
- Containerized deployment with Docker

## API Endpoints

### 1. Create Short URL

Generates a shortened URL from a long URL.

```
POST /short

Request:
{
    "long_url": "https://long_url.com"
}

Response:
{
    "short_url": "http//base.url/shortCode"
}
```

### 2. URL Redirection

Process flow for accessing shortened URLs:

- First checks Redis cache for the short code
- If found in cache, immediately redirects to the original URL
- If not in cache, queries MongoDB for the original URL
- Saves the retrieved URL to cache for future requests
- Redirects to the original URL

## Caching Strategy

The service implements an efficient caching mechanism with the following characteristics:

- Redis-based caching system
- Maximum capacity of 1GB
- All-keys LRU (Least Recently Used) eviction policy
- Another Redis for protected keys (E.g. counter)
- Optimized for high-performance URL retrieval

## Deployment

The application is containerized using Docker for easy deployment and scaling.