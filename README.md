# EduBuddy

Students need tutors, and tutors need students; we connect them. We aim to offer a seamless and convenient learning experience. EduBuddy is developed using the Flutter framework, ensuring cross-platform compatibility. Tutors can create postings to showcase their availability and location on the UVic campus, while students can also make postings based on their tutoring requirements.

- Client code [here](https://github.com/mattrltrent/eng_comp_client).
- The team: [Jaspreet Sidhu](https://github.com/jaspreetks), [Chris Huk](https://github.com/TalentedB), [Hal Nguyen](https://github.com/hn275), and [Matthew Trent](https://github.com/mattrltrent).
- From idea to tangible product in <8 hours.

## Tables of Contents

1. [Technologies](#technologies)
2. [Architecture](#architecture)
3. [Key Features](#key-features)
4. [Bugs and Issues](#bugs-and-issues)
5. [Future Development](#future-development)
6. [Results 🥇](#results-1st-place-)

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

### Course Scraper

The course scraper feature allows the app to fetch and retrieve up-to-date course information from UVic's official sources. It ensures that the app has the most accurate and recent course data for that semester, enabling users to search for tutors and courses with precision.

### Login Authentication

The login authentication feature ensures the security of user accounts and information. Users can create accounts, log in securely, and access personalized features within the app. This feature protects user data and provides a seamless and trustworthy user experience.

### Frictionless App Design

The app incorporates a frictionless design approach, focusing on creating an intuitive and user-friendly interface. The design aims to minimize complexity and streamline the user journey, making it easy for users to navigate through the app, search for tutors, and schedule tutoring sessions effortlessly.

### Fuzzy Searching of Courses

The fuzzy searching feature enhances the search functionality within the app. It provides accurate and relevant results even when users input approximate or misspelled course names or keywords. This feature improves the user experience by ensuring users can find the desired courses quickly and easily.

### Tutoring Assignment

The tutoring assignment feature enables the app to match students with suitable tutors based on their academic needs and preferences. It considers factors such as course requirements, availability, location, and user preferences to facilitate efficient and effective tutoring assignments. This simplifies the process of finding the right tutor and provides students with a comprehensive directory of available tutors, enabling them to make informed decisions when selecting a tutor.

### Local Data Persistence

The local data persistence feature allows the app to store relevant data locally on the device. This enables seamless offline access to previously fetched information, such as tutor profiles, course details, and user preferences. Local data persistence ensures a consistent user experience, even in situations where internet connectivity is limited or unavailable.

## Bugs and Issues

- Deployment
- Connection Failure
- Possible SQL injection

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

## Results: 1st place 🥇

<img src="https://raw.githubusercontent.com/mattrltrent/random_assets/main/victory.JPG" alt="img" width="300" />

<sub>^ (everyone on the team got a certificate, this is just mine)</sub>
