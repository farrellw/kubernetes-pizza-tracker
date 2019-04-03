FROM node:8.12.0-alpine as build
RUN mkdir -p /build
COPY . /build/
WORKDIR /build
RUN npm install
RUN npm run build

FROM nginx:1.15.3-alpine
COPY --from=build /build/build/ /usr/share/nginx/html/
COPY --from=build /build/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80