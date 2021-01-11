## Load test using artillery

```shell
artillery run ping.yml
```

## Load test using Hey

```shell
hey -z 2100ms -c 20  http://localhost:8080/ping
```

