start-backend:
	cd service && go mod tidy && go run .

start-frontend:
	cd AI-Alchemy && npm install && npm run dev
