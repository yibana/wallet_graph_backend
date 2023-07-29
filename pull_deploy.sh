#!/bin/bash

# Check if the "-kill" parameter is passed
if [ "$1" = "-kill" ]; then
  # Kill the existing wallet_graph process if it's running
  if pgrep wallet_graph >/dev/null 2>&1; then
    echo "Stopping existing wallet_graph process..."
    pkill -f wallet_graph
  else
    echo "wallet_graph process is not running"
  fi
  exit 0
fi

# Check if wallet_graph process is already running, and end the process
if pgrep wallet_graph >/dev/null 2>&1; then
  echo "Stopping existing wallet_graph process..."
  pkill -f wallet_graph
fi

# 更新代码
echo "Updating repository..."
git reset --hard HEAD
git pull
# 编译
echo "Building binary..."
go build -o wallet_graph main.go



# Set environment variable
export MONGO_URL="mongodb://admin:3WPIki9dXShd6ZZhGXKZ@127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+1.8.2"
export REDIS_URL="redis://localhost:6379"
chmod +x pull_deploy.sh

# Start wallet_graph and log output to wallet_graph.log
nohup ./wallet_graph > wallet_graph.log 2>&1 &
# View logs
tail -f wallet_graph.log
