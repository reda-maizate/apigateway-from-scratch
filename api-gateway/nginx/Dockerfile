FROM nginx

# RUN apk update && apk add --no-cache certbot openssl certbot-nginx
RUN apt-get update \
    && apt-get install -y --no-install-recommends \
		certbot \
		python3-certbot-nginx \
    	iputils-ping \
    && rm -rf /var/lib/apt/lists/*

ENTRYPOINT ["tail", "-f", "/dev/null"]

# Generate certbot certonly
#RUN certbot certonly --standalone -d api.navyblueglove.uk --non-interactive --agree-tos --email reda.maizate@gmail.com

#COPY default.conf /etc/nginx/conf.d/default.conf

#RUN certbot --nginx -d api.navyblueglove.uk --non-interactive --agree-tos --email reda.maizate@gmail.com
# Debug mode -> Keep the container running
#ENTRYPOINT ["tail", "-f", "/dev/null"]