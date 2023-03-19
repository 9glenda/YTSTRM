# API Documentation

This API follows the **RESTful** standard and provides a way to stream videos from YouTube using the "yt-dlp" command line tool. It exposes a single endpoint at `/stream` that accepts HTTP GET requests with a video ID or URL as a query parameter.

## Endpoint

`/stream`

## Method

`GET`

## Request Query Parameters

| Parameter | Description                                         |
| ----      | ---                                                 |
| `id`      | The video ID or URL of the YouTube video to stream. |

## Responses

| Status Code                 | Response Body | Description                                                                                                     |
| ----                        | ---           | ---                                                                                                             |
| `200 OK`                    | `video/mp4`   | If the video is successfully streamed, the response body will contain the video data in the "video/mp4" format. |
| `400 Bad Request`           |               | If the video ID or URL is missing or invalid.                                                                   |
| `500 Internal Server Error` |               | If there was an error while streaming the video.                                                                |

## Example Usage

Assuming the API is running on `localhost:8080`, the following HTTP GET request can be made to stream a video with ID `xyz123`:

bash

`http://localhost:8080/stream?id=xyz123`

**Note:** This API uses the "yt-dlp" command line tool to download videos from YouTube. Therefore, the tool must be installed and available on the system running the API for it to work correctly.
## Download Endpoint

`/download`

## Method

`GET`

## Request Query Parameters

| Parameter | Description |
| --- | --- |
| `id` | The video ID or URL of the YouTube video to download. |

## Responses

| Status Code | Response Body | Description |
| --- | --- | --- |
| `200 OK` |     | If the video is successfully downloaded, no response body will be returned. |
| `400 Bad Request` |     | If the video ID or URL is missing or invalid. |
| `500 Internal Server Error` |     | If there was an error while downloading the video. |

## Example Usage

Assuming the API is running on `localhost:8080`, the following HTTP GET request can be made to download a video with ID `xyz123`:

bash

`http://localhost:8080/download?id=xyz123`

**Note:** This API uses the "yt-dlp" command line tool to download videos from YouTube. Therefore, the tool must be installed and available on the system running the API for it to work correctly.
