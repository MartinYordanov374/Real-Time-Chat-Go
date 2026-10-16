# Real-time Chat Application in Go

# Tech Stack 
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Redis](https://img.shields.io/badge/redis-%23DD0031.svg?style=for-the-badge&logo=redis&logoColor=white)
![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
![React](https://img.shields.io/badge/react-%2320232a.svg?style=for-the-badge&logo=react&logoColor=%2361DAFB)
![JavaScript](https://img.shields.io/badge/javascript-%23323330.svg?style=for-the-badge&logo=javascript&logoColor=%23F7DF1E)
![TypeScript](https://img.shields.io/badge/typescript-%23007ACC.svg?style=for-the-badge&logo=typescript&logoColor=white)
![TailwindCSS](https://img.shields.io/badge/tailwindcss-%2338B2AC.svg?style=for-the-badge&logo=tailwind-css&logoColor=white)
![Next JS](https://img.shields.io/badge/Next-%23000.svg?style=for-the-badge&logo=next.js&logoColor=white)
![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)
![GitHub Actions](https://img.shields.io/badge/github%20actions-%232671E5.svg?style=for-the-badge&logo=githubactions&logoColor=white)

## Table of Contents

1. [The motivation behind the project](#The-motivation-behind-the-project)
2. [What I strived to achieve and learn through this project](#What-I-strived-to-achieve-and-learn-through-this-project)
3. [Project Features](#Project-Features)
5. [Project Architecture](#Project-Architecture)
6. [Project Structure](#Project-Structure)
8. [Getting Started](#Getting-Started)
9. [Improvement Areas](#Improvement-Areas)

## The motivation behind the project
A couple of months ago I decided I wanted to learn Go, so I built a static load balancer. I saw that Go is very good for writing fast-executing and efficient code, so I decided to build a bigger project with Go - a real-time chat application.
## What I strived to achieve and learn through this project
The goals of this project were to learn how to write an efficient backend in Go, implement REST APIs in Go, learn Redis pub/sub, improve my system design skills, and improve my overall Go programming skills.
## Project Features
1. **Real-Time Messaging System**
    - Users can send messages to other users in real-time via WebSocket technology.
    - Users can delete sent messages.
    - If a chat does not exist between two users and a user sends a message, a message request is sent to the receiver.
    - If the receiver accepts the message, the chat remains and users can continue sending each other messages.
    - If the chat request is rejected by the receiver, the chat and request objects are deleted from the database.

2. **Message Caching**
    - The last 100 messages sent in a chat are cached in Redis.
    - This lowers message retrieval latency and reduces the number of database fetches without compromising functionality.
    - If a user deletes a message in a chat, the respective messages cache is invalidated.
      
3. **Redis Pub/Sub**
    - The application publishes and subscribes to two Redis events: "Message" and "Request".
    - A goroutine is started on startup of the Go backend that handles subscriptions to the "Message" and "Request" Redis events.
    - Every user is subscribed to both events, so when another user sends them a chat request or a message, they receive it in real time through the combined WebSocket and Redis pub/sub pipeline.
    
3. **User Searching System**
    - Users can search for other users by username and send them message requests.
    - This also includes searching users within their existing contacts.
    - Database indexing was used to increase database search performance.

4. **Server-Side Sessions**
    - The application makes use of server-side sessions to prevent client-side attacks such as cookie stealing, which could lead to account takeover.
    - Every user server session is created upon a successful login and lasts 1 hour.
    
5. **Login System**
    - Users can log in and create accounts, granted that their username is not already taken.
    - Password requirements are implemented following the latest NIST guidelines.
    - The login system features email validity checks, password complexity checks, and username availability checks.

## Project Architecture
## Project Structure
## Getting Started
## Improvement Areas
