FROM grafana/grafana:6.2.4

ENV GF_INSTALL_PLUGINS grafana-piechart-panel

COPY grafana.ini /etc/grafana/grafana.ini
COPY provisioning /etc/grafana/provisioning
COPY dashboards /var/lib/grafana/dashboards