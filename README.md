# Support Analytics

Checkout support-analytics into your GOTPATH

## Setup

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## Build and test

```
cd $GOPATH/src/github.com/replicatedcom/support-analytics
make deps build test
```

## Generate a dependency graph
Perform these steps If you would like to generate a dependency graph from a vendors' yaml.
```
brew install graphviz
```
```
bin/support events --yaml app.yaml > app.graph
dot -Tpng app.graph > app.png
cp app.png ~/Pictures
```

## Generate multiple dependency graphs
To generate a .png for every app yaml within the yaml directory, run
```
make graphs
```
