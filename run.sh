term() {
    echo "Caught SIGTERM signal!"
    kill -TERM "$child" 2>/dev/null
}

trap _term SIGTERM

if [ -f upload ]; then
    mv upload echoapp
    chmod +x echoapp
fi

./echoapp server --config=config.yaml &

child=$!
wait "$child"

