FROM node:alpine AS build

WORKDIR /app/site

COPY package.json .
COPY package-lock.json .

# some legacy dependancies are cooked
RUN npm install --legacy-peer-deps

COPY . .

RUN npm run build

FROM nginx

COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/site/dist /usr/share/nginx/html
