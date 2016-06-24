
### Setup keys in vault
```bash
export STORAGENAME=cfmanifest

$> vault write secret/${STORAGENAME} \
ca-cert=@ca-cert.txt
consul-agent-cert=@consul-agent-cert.txt
consul-agent-key=@consul-agent-key.txt
consul-server-cert=@consul-server-cert.txt
consul-server-key=@consul-server-key.txt
etcd-server-cert=@etcd-server-cert.txt
etcd-server-key=@etcd-server-key.txt
etcd-client-cert=@etcd-client-cert.txt
etcd-client-key=@etcd-client-key.txt
etcd-peer-cert=@etcd-peer-cert.txt
etcd-peer-key=@etcd-peer-key.txt
bbs-server-cert=@bbs-server-cert.txt
bbs-server-key=@bbs-server-key.txt
bbs-client-cert=@bbs-client-cert.txt
bbs-client-key=@bbs-client-key.txt
rtr-ssl-cert=@rtr-ssl-cert.txt
rtr-ssl-key=@rtr-ssl-key.txt

$> omg --print-manifest keyparty /
--storage-name ${STORAGENAME} /
--manifest-file dev-manifest-cf.yml 
```
