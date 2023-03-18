from fabric import Connection, task
import socket
from kubernetes import client, config
from kubernetes.watch import Watch
import logging

logging.basicConfig()
logging.getLogger().setLevel(logging.DEBUG)

namespace = 'logfire-local'


@task
def logsgowebapi(c):
    with Connection(host='api.logfire.sh', user='root', connect_kwargs={'key_filename': 'id_rsa'}) as conn:
        conn.get('/root/.kube/config', local='./kubeconfig')
        config.load_kube_config('./kubeconfig')
        api = client.CoreV1Api()
        pod_list = api.list_namespaced_pod(namespace)
        for pod in pod_list.items:
            if 'gowebapi' in pod.metadata.name:
                w = Watch()
                for e in w.stream(api.read_namespaced_pod_log, name=pod.metadata.name, namespace=namespace, follow=True):
                    print(e)


@task
def logslivetail(c):
    with Connection(host='api.logfire.sh', user='root', connect_kwargs={'key_filename': 'id_rsa'}) as conn:
        conn.get('/root/.kube/config', local='./kubeconfig')
        config.load_kube_config('./kubeconfig')
        api = client.CoreV1Api()
        pod_list = api.list_namespaced_pod(namespace)
        for pod in pod_list.items:
            if 'livetail' in pod.metadata.name:
                w = Watch()
                for e in w.stream(api.read_namespaced_pod_log, name=pod.metadata.name, namespace=namespace, follow=True):
                    print(e)


@task
def logsfilter(c):
    with Connection(host='api.logfire.sh', user='root', connect_kwargs={'key_filename': 'id_rsa'}) as conn:
        conn.get('/root/.kube/config', local='./kubeconfig')
        config.load_kube_config('./kubeconfig')
        api = client.CoreV1Api()
        pod_list = api.list_namespaced_pod(namespace)
        for pod in pod_list.items:
            if 'filter' in pod.metadata.name:
                w = Watch()
                for e in w.stream(api.read_namespaced_pod_log, name=pod.metadata.name, namespace=namespace, follow=True):
                    print(e)


@task
def logsnginx(c):
    with Connection(host='api.logfire.sh', user='root', connect_kwargs={'key_filename': 'id_rsa'}) as conn:
        conn.get('/root/.kube/config', local='./kubeconfig')
        config.load_kube_config('./kubeconfig')
        api = client.CoreV1Api()
        pod_list = api.list_namespaced_pod(namespace)
        for pod in pod_list.items:
            if 'nginx' in pod.metadata.name:
                w = Watch()
                for e in w.stream(api.read_namespaced_pod_log, name=pod.metadata.name, namespace=namespace, follow=True):
                    print(e)


@task
def logsnotification(c):
    with Connection(host='api.logfire.sh', user='root', connect_kwargs={'key_filename': 'id_rsa'}) as conn:
        conn.get('/root/.kube/config', local='./kubeconfig')
        config.load_kube_config('./kubeconfig')
        api = client.CoreV1Api()
        pod_list = api.list_namespaced_pod(namespace)
        for pod in pod_list.items:
            if 'notification' in pod.metadata.name:
                w = Watch()
                for e in w.stream(api.read_namespaced_pod_log, name=pod.metadata.name, namespace=namespace, follow=True):
                    print(e)
