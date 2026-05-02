
echo "installing redis"
# Linux
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
   
    sudo apt-get install lsb-release curl gpg
    curl -fsSL https://packages.redis.io/gpg | sudo gpg --dearmor -o /usr/share/keyrings/redis-archive-keyring.gpg
    sudo chmod 644 /usr/share/keyrings/redis-archive-keyring.gpg
    echo "deb [signed-by=/usr/share/keyrings/redis-archive-keyring.gpg] https://packages.redis.io/deb $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/redis.list
    sudo apt-get update
    sudo apt-get install redis

# Mac OSX
elif [[ "$OSTYPE" == "darwin"* ]]; then
    # Mac OSX
   echo "test"

else
    # Unknown
    echo "unknown os type"
fi
