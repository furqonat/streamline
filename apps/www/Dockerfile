
FROM node:20.11.1-alpine3.18

WORKDIR /web/apps

COPY apps/www/ ./apps/www
COPY .env .
RUN true
COPY .eslintignore .
RUN true
COPY .eslintrc.json .
RUN true
COPY .editorconfig .
RUN true
COPY .prettierrc .
RUN true
COPY .prettierignore .
RUN true
COPY jest.config.ts .
RUN true
COPY jest.preset.js .
RUN true
COPY nx.json .
RUN true
COPY package.json .
RUN true
COPY tsconfig.base.json .
RUN true
COPY pnpm-lock.yaml .
RUN true
COPY prisma/ prisma
COPY package.json .


RUN npm i -g pnpm
RUN pnpm install
RUN npx prisma generate
RUN true

RUN  npx nx run www:build


CMD ["npx nx run www:start" ]