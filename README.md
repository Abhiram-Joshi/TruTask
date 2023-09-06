# Engineering Task

## Important Notes Before We Start

- The default username and password for the postgres database is hard-coded as *postgres* and *postgres* for ease of evaluation. The username and password can be added to an env file and then used to prevent exposing credentials to the repository.

- Users have been divided into 2 categories: admin and non-admin. Signing up as a user creates a non-admin user by default. Giving a user admin privileges will require you to change the value of the *is_admin* field of the corresponding user to 1.


## Problem Statement
Build a robust containerized task management system to handle user  authentication, authorization and access management.



## Features

- Secure user registration and authentication
- Account Deactivation and Deletion: Allow users to deactivate or delete their accounts, if applicable. Implement a mechanism to handle account deletion securely while considering data retention policies.
- Role-based and Group-based access management on resource (Tasks) with ability to create custom roles and groups (Need to make sure endpoints are secure)
- Protection against vulnerabilities like SQL injection attacks
- Support for bulk upload using CSV(Both users and tasks) making sure all the relationships are preserved accurately


## Usage

1. Clone the repository to your local system
```
git clone https://github.com/BalkanID-University/balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi.git
```

2. Move into the clone directory
```
cd balkanid-fte-hiring-task-vit-vellore-2023-Abhiram-Joshi
```

3. Rename the *.env.example* file to *.env* and run the following command
```
docker-compose up
```

4. Check the [documentation](https://documenter.getpostman.com/view/14681434/2s946mZ9eM) to run the application on your system



## Database Overview

### Table Overview
- User model stores user data
- Role model stores data related to user groups
- Task model stores data related to tasks

### Relationship Overview
- User model has a many-to-many relationship with the Role model
- Role model has a one-to-many relationship with the Task Model