# Install App
helm upgrade --install myapp \
 --namespace=default \
 charts/myapp \
 -f charts/myapp/values.yaml \
 --set image.tag=v1.0 \
 --set image.repository=myapp
Release "myapp" does not exist. Installing it now.
NAME: myapp
LAST DEPLOYED: Sun Jan 16 10:32:02 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:

1. Get the application URL by running these commands:
   export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=myapp,app.kubernetes.io/instance=myapp" -o jsonpath="{.items[0].metadata.name}")
  export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT

# Install Prometheus
helm install my-prometheus charts/kube-prometheus -f charts/kube-prometheus/values.yaml
W0116 15:28:09.264586    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.268026    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.271125    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.274469    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.278698    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.677431    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.679789    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.679850    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.679853    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
W0116 15:28:09.680144    1978 warnings.go:70] policy/v1beta1 PodSecurityPolicy is deprecated in v1.21+, unavailable in v1.25+
NAME: my-prometheus
LAST DEPLOYED: Sun Jan 16 15:28:08 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: kube-prometheus
CHART VERSION: 6.6.0
APP VERSION: 0.53.1

** Please be patient while the chart is being deployed **

Watch the Prometheus Operator Deployment status using the command:

    kubectl get deploy -w --namespace default -l app.kubernetes.io/name=kube-prometheus-operator,app.kubernetes.io/instance=my-prometheus

Watch the Prometheus StatefulSet status using the command:

    kubectl get sts -w --namespace default -l app.kubernetes.io/name=kube-prometheus-prometheus,app.kubernetes.io/instance=my-prometheus

Prometheus can be accessed via port "9090" on the following DNS name from within your cluster:

    my-prometheus-kube-prometh-prometheus.default.svc.cluster.local

To access Prometheus from outside the cluster execute the following commands:

    echo "Prometheus URL: http://127.0.0.1:9090/"
    kubectl port-forward --namespace default svc/my-prometheus-kube-prometh-prometheus 9090:9090

Watch the Alertmanager StatefulSet status using the command:

    kubectl get sts -w --namespace default -l app.kubernetes.io/name=kube-prometheus-alertmanager,app.kubernetes.io/instance=my-prometheus

Alertmanager can be accessed via port "9093" on the following DNS name from within your cluster:

    my-prometheus-kube-prometh-alertmanager.default.svc.cluster.local

To access Alertmanager from outside the cluster execute the following commands:

    echo "Alertmanager URL: http://127.0.0.1:9093/"
    kubectl port-forward --namespace default svc/my-prometheus-kube-prometh-alertmanager 9093:9093

# Install Grafana
 make gi 
helm install my-grafana bitnami/grafana -f charts/grafana/values.yaml
NAME: my-grafana
LAST DEPLOYED: Sun Jan 16 15:29:39 2022
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
CHART NAME: grafana
CHART VERSION: 7.6.0
APP VERSION: 8.3.3

** Please be patient while the chart is being deployed **

1. Get the application URL by running these commands:
    echo "Browse to http://127.0.0.1:8080"
    kubectl port-forward svc/my-grafana 8080:3000 &

2. Get the admin credentials:

    echo "User: admin"
    echo "Password: $(kubectl get secret my-grafana-admin --namespace default -o jsonpath="{.data.GF_SECURITY_ADMIN_PASSWORD}" | base64 --decode)"

