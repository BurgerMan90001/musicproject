FROM node:alpine AS build

WORKDIR /app/site

COPY package.json .

# some legacy dependancies are cooked
RUN npm install --legacy-peer-deps

COPY . .

RUN npm run build


FROM nginx:alpine AS run


COPY --from=build /app/site/dist/ ./usr/share/nginx/html
