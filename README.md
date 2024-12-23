# GuestBook

This is a GuestBook for [edmateo.site](https://edmateo.site). It's only a shitty REST API

## Technologies

1. **GoLang**  
2. **PostgreSQL**

## Deployment

### 1. Create a Database
Create a PostgreSQL database for storing guestbook entries:
```bash
su -c "createuser --pwprompt guestbook" postgres
su -c "psql -c 'CREATE DATABASE guestbook OWNER guestbook;'" postgres
```

### 2. Clone the Repository
Clone this repo to your server:
```bash
git clone https://github.com/ImnotEdMateo/guestbook.git
```

### 3. Build the Project
Navigate to the project directory and build the binary:
```bash
cd guestbook/
go build ./cmd/main
```
After building, you should see a `main` binary in the project root.

### 4. Create a Systemd Service
To manage the application as a service, create a Systemd service file:
```bash
sudo vim /etc/systemd/system/guestbook.service
```
Add the following content to the file:
```
[Unit]
Description=GuestBook
After=network.target

[Service]
Environment=DB_HOST=yourhost
Environment=DB_USER=guestbook
Environment=DB_PASSWORD=yourpassword
Environment=DB_NAME=guestbook
Environment=DB_PORT=yourpostgresport
User=root
Type=simple
Restart=always
ExecStart=/absoulte/path/to/your/binary/main

[Install]
WantedBy=multi-user.target
```

### 5. Reload and Start the Service
Reload the Systemd daemon and start the GuestBook service:
```bash
sudo systemctl daemon-reload
sudo systemctl start guestbook
```

### 6. Access the API
Open your browser and navigate to your server’s address on port `3000`. You should see an empty JSON response, indicating the API is running.
