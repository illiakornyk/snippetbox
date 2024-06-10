# Snippet Box

## Overview
Snippet Box is a personal project management web application that allows users to securely create, manage, and set expiration timers for their notes. This application is built with Go (Golang) and follows the practices outlined in the book "Let's Go" by Alex Edwards.

## Features
- **User Authentication:** Users can sign up and sign in to manage their personal space securely.
- **Note Creation and Management:** Users have the ability to create notes and manage them within the application.
- **Expiration Timers:** Each note can have a timer set for its duration, with options for a day, week, or month.
- **Security:** The application uses HTTPS and secure headers to protect against common web attacks. Sensitive user data is encrypted to ensure privacy.

## Technologies
- **Backend:** Written in Go, utilizing libraries for HTTP routing, MySQL database interaction, and cryptographic functions.
- **Frontend:** (Planned for future development) To be separated from the backend for a more modular approach.
- **Database:** MySQL is used for storing user data, with encryption for sensitive information.

## Current Capabilities
- Users can currently sign up, sign in, create, delete notes, and set expiration timers for their notes within the application.

## Future Plans
- The project aims to transition into a more modular architecture, potentially employing microservices for each core functionality to enhance scalability and maintainability.

## Acknowledgements
This project has been developed with guidance from the book "Let's Go" by Alex Edwards, which has been instrumental in the application's design and implementation.
