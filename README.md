# Lang Trainer
Word training application for different languages

This is fully my hobby project and it's developed for learning principally Swedish vocabulary. I need to pass Swedish course provided by the school, so I need help and this is what I need. I'm going to make this software larger as a programming workout.

## Specification

Users have three different access levels. A higher access level inherits the properties of all lower ones.

Admin 
* As an admin I can add new users
* As an admin I can remove existing users
* As an admin I can change user priviledges

Teacher
* As a teacher I can create a new group
* As a teacher I can add new words to group
* As a teacher I can add new translations to word
* As a teacher I can remove words from specific group
* As a teacher I can remove groups
* As a teacher I can publish groups for specific students
* As a teacher I can add new students

Student
* As a student I can practise words which belong to group where I have access
* As a student I can log in

Maybe I can get started with these. More specifications later.

## Staging

I'm deployed this software to Heroku and you can test it if you like. 

https://trainerlang.herokuapp.com/

* username: Admin
* password: salainen

## Setting up development environment

Before starting you should install [Go](https://golang.org/) and [Node](https://nodejs.org/en/). Note that frontend may not work with latest node version. I suggest you to install node version 14 or 16. You also need to setup postgres database to your local machine. Specify ```.env``` file where you store all necessary environment variables.

Next you should clone this repository.

### Installing dependencies

Applications frontend is located to ```client``` directory. To install its dependencies you should go inside it and run

    npm install

Backend depedencies can be installed using command ```go mod download``` in the repository root.

### Starting and testing

You can start frontend with command

    npm start
    
To start backend you can simply run 

    sh go start
    
Backend tests be run with command

    sh go test
