FROM node:20-alpine AS build

WORKDIR /

COPY package.json package-lock.json ./
RUN npm ci

COPY . .
RUN npm run build