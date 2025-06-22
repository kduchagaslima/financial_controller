FROM node:18 as build
WORKDIR /app
COPY package.json tsconfig.json angular.json ./
COPY src ./src
RUN npm install -g @angular/cli && npm install && ng build --configuration production

FROM nginx:alpine
COPY --from=build /app/dist/finops-frontend /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
