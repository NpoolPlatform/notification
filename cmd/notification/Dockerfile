FROM jrei/systemd-centos:7

COPY Notification.viper.yaml /etc/Notification/Notification.viper.yaml
COPY notification /usr/local/bin/notification

CMD ["notification", "run"]
