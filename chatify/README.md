# Chatify

WebSocket-based chat server for real-time communication

## Overview

Chatify provides a simple and efficient way to set up a real-time chat system using WebSocket technology. It enables seamless communication between clients, making it ideal for applications that require instant messaging features.

## Features
- **Modular Design:** Utilize modular design principles to enable easy expansion and customization of server functionalities;

- **On Connection Callbacks:** Trigger custom actions when new clients connect to the server, allowing for dynamic interactions;

- **Middleware Support:** Integrate middleware functions to extend server functionality and enhance the user experience. For instance, it enables to add authentication and access control middlewares.

- **Flexible Configuration:** Fine-tune server settings, including custom port, WebSocket path, and more, using intuitive configuration options.

- **Message Processing:** Efficiently process incoming messages to ensure secure and organized communication. Available messages handlers: Formatter, message persisting and customs.

## Running the Project

To run the project, execute the following command:

```bash
go run cmd/server/main.go
``` 

To use chatify in your application, check the examples in the `examples` folder.

## Contributing

Contributions to the project are welcome! If you'd like to contribute to Chatify, please follow these guidelines:

- Fork this repository and create a new branch for your feature or bug fix.
- Commit your changes with clear and descriptive commit messages.
- Submit a pull request detailing your changes and explaining how they enhance the project.


For any questions or issues, feel free to contact me at mikaellaferreira0@gmail.com. Happy coding!