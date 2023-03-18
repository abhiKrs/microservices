#!/usr/bin/env bash
#cd /opt/microservices/helm-chart
helm upgrade webapp /opt/microservices/helm-chart/webapp-prodchart/ -n logfire-local --set=gowebapiDeployment.gowebapi.image.tag=$1 --set=livetailDeployment.livetail.image.tag=$1 --set=notificationDeployment.notification.image.tag=$1
