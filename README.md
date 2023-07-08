# EduBuddy

Students need tutors, and tutors need students; we connect them. We aim to offer a seamless and convenient learning experience. EduBuddy is developed using the Flutter framework, ensuring cross-platform compatibility. Tutors can create postings to showcase their availability and location on the UVic campus, while students can also make postings based on their tutoring requirements.

## Tables of Contents

1. [Technologies](#technologies)
2. [Architecture](#architecture)
3. [Key Features](#key-features)
4. [Bugs and Issues](#bugs-and-issues)
5. [Future Development](#future-development)

## Technologies

### Backend Technologies

- Go: The backend is implemented using Go programming language, providing efficient and scalable server-side logic.
- Docker: Docker is used for containerization, enabling easy deployment and management of backend components.
- JSON: JSON is used for data serialization and communication between the backend and frontend.
- YAML: YAML configuration files are used for defining app settings and configurations.
- Python: Python is utilized for implementing additional backend functionalities, such as data processing or external API integrations.
- SQLite: The app uses SQLite as the local database for storing essential data.

### Frontend Technologies

- Flutter: The app’s frontend is developed using the Flutter framework, enabling cross-platform compatibility for iOS, Android, and web.
- Dart: Dart programming language is used for building the frontend logic and UI components.
- Dependency Injection: The app implements dependency injection to manage and provide necessary dependencies to various components.
- Native Channel Integration: Native channel integration is used to communicate and leverage platform-specific functionalities if required.
- Local Database: The app utilizes a local database for caching data and improving performance when offline.
- Encrypted Storage: To ensure data security, sensitive information is stored in an encrypted format.

## Architecture

### Backend Architecture

- The backend architecture of the app follows a microservices-based approach. The Go programming language is used to develop individual microservices, each responsible for specific functionalities, such as user management, course management, and scheduling.
- The microservices are containerized using Docker, enabling easy deployment and scalability. Communication between microservices is established using JSON for data exchange, and YAML files are used for configuration.
- Python scripts are employed for additional backend tasks, such as data processing or integrating with external APIs. The SQLite database stores relevant app data, ensuring efficient and reliable data storage.

### Frontend Architecture

- The frontend architecture is built using the Flutter framework. The app’s UI is constructed using Dart programming language, providing a responsive and interactive user experience.
- The app adopts a modular and component-based approach, utilizing Flutter widgets for building reusable UI elements. Dependency injection is implemented to manage and provide dependencies throughout the app.
- Native channel integration allows for seamless communication with platform-specific functionalities, if required. The local database is utilized for caching data, enabling offline functionality and improved performance.

## Key Features

- Course Scraper
- Login Authentication
- Frictionless Design of the App
- Fuzzy Searching of Courses
- Tutoring Assignment
- Tutor Finding
- Local Data Persistence

## Bugs and Issues

- Deployment
- Possible SQLi

## Future Development

### Course-based Filtering

The app will provide a robust filtering mechanism based on courses. The filtering functionality ensures that users can search for relevant tutors and courses that align with their specific academic needs for that semester. Tutors can create postings for their availability and location on campus, and students can also create postings based on their needs.

### Location Finding

The location finding feature will help users locate their tutors on the UVic campus. It will utilize the device’s GPS capabilities to determine the user’s location and display the nearby tutors on a campus map within the app. This feature will provide convenience and facilitate efficient in-person tutoring sessions.

### BCrypt Encryption

BCrypt encryption will be implemented to enhance data security. It will be utilized to encrypt sensitive user information, such as passwords, before storing them in the database. BCrypt is a robust encryption algorithm that ensures the safety of user data, protecting it from unauthorized access.

### Better Fuzzy Searching

To improve the searching capability, the app will implement a better fuzzy searching mechanism. This feature will allow users to find available courses more easily by accepting approximate search terms and returning relevant results. The implementation will utilize advanced search algorithms, such as trigram or fuzzy string matching, to provide accurate and efficient search results.

### Search Caching

Search caching will be implemented using Redis. This feature aims to improve search performance by storing frequently executed search queries and their results in memory. By caching search results, subsequent searches with the same query will retrieve data from the cache, reducing the load on the backend and enhancing overall app responsiveness.
