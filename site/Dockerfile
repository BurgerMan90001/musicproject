FROM node:alpine AS build

WORKDIR /app/site

COPY package.json .
COPY package-lock.json .

RUN npm ci
COPY . .

RUN npm run build

#FROM nginx:stable-alpine

#COPY --from=build /app/site/nginx/nginx.conf /etc/nginx/conf.d/default.conf
#COPY --from=build /app/site/dist /usr/share/nginx/html

#RUN chmod -R 755 /usr/share/nginx/html

#EXPOSE 80

#ENTRYPOINT [ "nginx", "-g", "daemon off;" ]