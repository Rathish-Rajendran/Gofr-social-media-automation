Gofr Social Media Presence Automation

To get the project up and running, you need three things:
1. Start the LLM models
2. Start the backend GoFr service
3. Start the frontend react service

# Backend (GoFr) service

## Starting
run (make sure you have the `make` utility)
```
make start-backend
```
to start the backend service that is responsible for contacting
the LLMs and generating the necessary responses needed for the
frontend

It is created by leverating the GoFr framework.

# Frontend (react) service

## Starting
run
```
make start-frontend
```
to start the react frontend service.

Open `http://localhost:5173/` in your browser to interact with the
UI.

# Features

What are available in the UI
- A common page to control Twitter/X, LinkedIn and Mail
- AI (Claude AI at the core) is used to generate the necessary responses
- Approval mechanism to post the generated posts and mails

NOTE: The code will need some tokens and secret keys to send and receive any
data from the aforementioned platforms. Please make sure to add them in
`service/.secrets.env` env file.

TODO: list the necessary secret keys...
