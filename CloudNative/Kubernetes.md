# 1 在 Pod 中的容器之间如何共享进程命名空间？
使用 Pod `.spec` 中的 `shareProcessNamespace` 字段可以启用进程命名空间共享。
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox
spec:
  shareProcessNamespace: true
  containers:
  - name: busybox
    image: busybox:1.28
    command: ["sleep", "1d"]
```

在容器中查看进程如下，可以看到 PID 为 1 的进程是 pause 容器，说明容器间共享了进程命名空间。
```bash
$ kubectl exec -it busybox -- ps ax
PID   USER     TIME  COMMAND
    1 root      0:00 /pause
    6 root      0:00 sleep 1d
   23 root      0:00 ps ax
```

参考资料：
- [在 Pod 中的容器之间共享进程命名空间](https://kubernetes.io/zh-cn/docs/tasks/configure-pod-container/share-process-namespace/)