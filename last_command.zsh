
preexec () {
    cmd=$(echo $2 | base64)
    #echo "Latest Command (base64): $cmd"
    /usr/bin/curl -s -d "${cmd}" localhost:8080/setCommand 2>&1 >/dev/null
}

echo "Loaded last-command hook. preexec function overloaded"
