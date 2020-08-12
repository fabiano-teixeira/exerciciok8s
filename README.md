## Informações sobre o exercício "Utilizando K8S" (12/08/2020)

Exercício realizado conforme informações abaixo:<br /><br />

Utilizando os conhecimentos adquiridos até o momento, crie os arquivos declarativos do Kubernetes para que os seviços abaixam possam ser executados.<br /><br />

1) Servidor Web - Nginx<br /><br />

Arquivos localizados na pasta 1-webserver-nginx.<br /><br />

2) Configuração do MySQL<br /><br />

Arquivos localizados na pasta 2-mysql.<br /><br />

3) Desafio Go!<br /><br />

Arquivos localizados na pasta 3-desafio_go.<br />
Aplicativo criado conforme instruções do exercício e disponibilizado na porta 8000.<br />
Testes Criados.<br />
Processo de CI ativado no GCB. *<br />
Imagem gerada e publicada no dockerhub, no endereço: <br />
    https://hub.docker.com/repository/docker/fabianoteixeirasilva/webserver-go-greeting<br />

Serviço disponibilizado no GCP, conforme instruções.<br /><br />

Said do comando kubectl get svc: <br /><br />
```
kubectl get svc

NAME                            TYPE           CLUSTER-IP    EXTERNAL-IP     PORT(S)          AGE
kubernetes                      ClusterIP      10.4.0.1      <none>          443/TCP          23h
mysql-exerck8s-service          ClusterIP      None          <none>          3306/TCP         5h32m
nginx-exerck8s-service          LoadBalancer   10.4.13.248   34.70.134.170   80:31691/TCP     17h
php-apache-hpa                  ClusterIP      10.4.15.129   <none>          80/TCP           19h
webserver-go-greeting-service   LoadBalancer   10.4.1.236    34.71.211.54    8000:31220/TCP   116s
```
