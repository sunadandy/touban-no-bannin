FROM node:16.13.0

WORKDIR /touban-no-bannin

# 00_front以下をコンテナのtouban-no-banninディレクトリにコピー
COPY . /touban-no-bannin/

RUN npm ci
# RUN npm run build

CMD ["npm", "run", "serve"]