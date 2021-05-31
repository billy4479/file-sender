all:
	frontend
	backend

frontend:
	mkdir -p dist
	cd frontend; \
	yarn build --emptyOutDir

frontend-dev:
	cd frontend; \
	yarn dev --host

backend:
	mkdir -p dist
	cd backend; \
	go build -o ../dist/file-sender

backend-dev:
	cd backend; \
	CompileDaemon -build='go build -o ../dist/file-sender' -command='../dist/file-sender'