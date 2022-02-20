FROM mysql

EXPOSE 3306

ADD ./volume/my.cnf /etc/mysql/conf.d/my.cnf

RUN mkdir /var/log/mysql
RUN chown mysql:mysql /var/log/mysql

CMD ["mysqld"]
