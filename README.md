# YTSTRM
This is a simple web application that allows users to stream and download YouTube videos. The application is built with the Gin web framework in Go and uses the yt-dlp command-line tool to download and stream videos.
Usage

To use the application, simply run the main function from the command line:

go

go run main.go

This will start the application on port 8080. The following endpoints are available:
/

This endpoint serves the stream.html file, which contains a simple HTML form for entering a YouTube video ID or URL.
/stream

This endpoint takes a YouTube video ID or URL as a query parameter and streams the video to the client. The video is streamed in the MP4 format.
/download

This endpoint takes a YouTube video ID or URL as a query parameter and downloads the video to the server. The downloaded video is saved in the MP4 format with the video ID as the file name.
Dependencies

This application requires the following dependencies:

    Gin: A web framework for Go
    yt-dlp: A command-line tool for downloading and streaming YouTube videos

Limitations

This application has the following limitations:

    The application only supports YouTube videos.
    The application does not perform any error handling for failed video downloads or streams.
    The application does not perform any input validation for the video ID or URL.

Contributing

Contributions to this application are welcome. If you find a bug or have a feature request, please open an issue or submit a pull request.
License

This application is licensed under the GLP v3 License. See the LICENSE file for more information.
