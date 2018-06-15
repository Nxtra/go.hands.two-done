FROM node:10-alpine
RUN npm install -g @angular/cli
WORKDIR /fif
COPY fif/ .
RUN npm install && ng build --prod

FROM nginx:alpine
COPY nginx.conf /etc/nginx/nginx.conf
WORKDIR /usr/share/nginx/html
COPY --from=0 /fif/dist/fif/ .
