### Informações sobre o exercício "Utilizando K8S" (12/08/2020)

<p>
Exercício realizado conforme informações abaixo:<br />

Utilizando os conhecimentos adquiridos até o momento, crie os arquivos declarativos do Kubernetes para que os seviços abaixam possam ser executados.

1) Servidor Web - Nginx  
    Arquivos localizados na pasta 1-webserver-nginx.

2) Configuração do MySQL  
    Arquivos localizados na pasta 2-mysql.<br />

3) Desafio Go!  
    Arquivos localizados na pasta 3-desafio_go.  
    Aplicativo criado conforme instruções do exercício e disponibilizado na porta 8000.  
    Testes Criados.  
    Processo de CI ativado no GCB. Log do build encontra-se no arquivo "build.log"  
    Imagem gerada e publicada no dockerhub, no endereço:   
        https://hub.docker.com/repository/docker/fabianoteixeirasilva/webserver-go-greeting<br />

    Serviço disponibilizado no GCP, conforme instruções.<br /> </p>

Saida do comando kubectl get pods:  
```
kubectl get pods
NAME                                     READY   STATUS    RESTARTS   AGE
mysql-server-exerck8s-85c56fbd59-ggctc   1/1     Running   0          21h
nginx-exerck8s-65866d67b8-2tsjv          1/1     Running   0          21h
nginx-exerck8s-65866d67b8-7mvds          1/1     Running   0          21h
nginx-exerck8s-65866d67b8-q8hnd          1/1     Running   0          21h
php-apache-hpa-dbcb5c7d9-pwldb           1/1     Running   0          23h
webserver-go-greeting-6b6bf9474d-z57qk   1/1     Running   0          19s
```
<br />
Saida do comando kubectl get svc:  
```
kubectl get svc
NAME                            TYPE           CLUSTER-IP    EXTERNAL-IP     PORT(S)          AGE
kubernetes                      ClusterIP      10.4.0.1      <none>          443/TCP          27h
mysql-exerck8s-service          ClusterIP      None          <none>          3306/TCP         10h
nginx-exerck8s-service          LoadBalancer   10.4.13.248   34.70.134.170   80:31691/TCP     21h
php-apache-hpa                  ClusterIP      10.4.15.129   <none>          80/TCP           23h
webserver-go-greeting-service   LoadBalancer   10.4.1.236    34.71.211.54    8000:31220/TCP   4h41m
```