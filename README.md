# Basic Go Book Project

## Description

This project is a basic Book Management System that allows users to create and manage books and authors. It utilizes the PostgreSQL database for data storage. The project provides functionality for signing up, logging in, creating new entries, and displaying all entries. The project utilizes the GORM library for database-related operations and Gorilla Sessions for session authentication.

## Features

- User Authentication: Users can sign up, log in, and log out of the system using session-based authentication.
- Create Authors: Users can create new author entries by providing the author's name.
- Create Books: Users can create new book entries by providing the book title and selecting an existing author.
- Display Entries: All book entries are displayed, including the book title, author name, and author ID.

## Technologies Used

- GORM: The project utilizes the GORM library, which provides an Object-Relational Mapping (ORM) for interacting with the database.
- Gorilla Sessions: Gorilla Sessions is used for session management and authentication.
- PostgreSQL: The project uses PostgreSQL as the underlying database for data storage.
- HTML/CSS: The project includes HTML templates styled with Tailwind CSS for the user interface.


## Usage

1. Sign Up: Create a new user account by providing the required details.
2. Log In: Log in to the system using your credentials.
3. Create Author: Create a new author by providing the author's name.
4. Create Book: Create a new book by providing the book title and selecting an existing author.
5. Display Entries: View all book entries with their corresponding details.
6. Log Out: Log out of the system to end the session.
request.



## Acknowledgements

- [GORM](https://gorm.io/)
- [Gorilla Sessions](https://www.gorillatoolkit.org/pkg/sessions)
- [PostgreSQL](https://www.postgresql.org/)
- [Tailwind CSS](https://tailwindcss.com/)

