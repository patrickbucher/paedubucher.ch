server {
	listen 443 http2 default_server;
	listen [::]:443 http2 default_server;

	root /var/www/paedubucher.ch;

	index index.html;

	server_name paedubucher.ch www.paedubucher.ch;

	location / {
		try_files $uri $uri/ =404;
	}

	ssl on;
	ssl_certificate /etc/letsencrypt/live/paedubucher.ch/fullchain.pem;
	ssl_certificate_key /etc/letsencrypt/live/paedubucher.ch/privkey.pem;

	gzip on;
	gzip_types text/css;
	gunzip on;
}

server {
	listen 0.0.0.0:80;
	server_name paedubucher.ch www.paedubucher.ch;
	rewrite ^ https://$host$request_uri? permanent;
}
