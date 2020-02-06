## How to deploy to prod:
- Create static IP and assign to the instance
- Create an A record for the static IP on your DNS settings
- Create an AWS Lightsail instance(can use Ubuntu) and add launch script with the following content
``` 
curl -o ./lightsail.sh https://raw.githubusercontent.com/cyber-republic/go-playground/master/webapp/lightsail-compose.sh

chmod +x ./lightsail.sh

./lightsail.sh
```