# Video Processing Pipeline

![video-processing-pipeline](assets/worker.png)

A high-performance video processing pipeline built in Go.

The service consumes background jobs from Redis, downloads uploaded videos, transcodes them into multiple HLS renditions, generates thumbnails and master playlists, uploads processed assets to Amazon S3, and notifies the API when processing is complete.

## Tech Stack

- Go
- FFmpeg
- Redis
- Asynq
- Amazon S3
- Docker

## Features

- Asynchronous background job processing
- Multi-bitrate HLS transcoding (1080p, 720p, 480p)
- Adaptive bitrate streaming via master playlist generation
- Automatic thumbnail extraction
- Recursive upload of processed assets to Amazon S3
- API callback on successful processing
- Temporary workspace management
- Dockerized for deployment

## Next

- Parallel transcoding using goroutines
- Configurable concurrency with semaphores
