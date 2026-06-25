# StreamIt Worker (Go)

![streamit-worker](assets/worker.png)

A high-performance distributed video processing worker powering StreamIt.

Built with Go for efficient concurrent media processing, the worker consumes background jobs, transcodes videos into HLS, generates thumbnails, uploads processed assets to S3, and notifies the StreamIt API when processing is complete.

## Tech Stack

- Go
- FFmpeg
- Redis
- Asynq
- AWS S3
- Docker

## Features

- Asynchronous background video processing
- HLS transcoding for adaptive streaming
- Automatic thumbnail generation
- Upload processed assets to S3
- API callback after successful processing
- Temporary workspace cleanup
- Dockerized for easy deployment
