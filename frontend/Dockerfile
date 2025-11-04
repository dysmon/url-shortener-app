FROM node:22-alpine AS build-stage
WORKDIR /app
ARG API_URL
ENV REACT_APP_API_BASE=${API_URL}
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
 
FROM nginx:1.21.3-alpine AS build-release-stage
COPY --from=build-stage /app/build /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
