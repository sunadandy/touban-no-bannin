FROM mysql:5.7

# [TODO]本来書くべきではない情報
ENV MYSQL_ROOT_PASSWORD=root
ENV MYSQL_DATABASE=toubanDB

# # MySQLの設定ファイル
COPY ./my.cnf /etc/mysql/conf.d/my.cnf
# DBのリストアのための情報
COPY ./init/dump.sql /docker-entrypoint-initdb.d/

RUN chmod 644 /etc/mysql/conf.d/my.cnf