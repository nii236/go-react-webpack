Early boilerplate app for:

- Go
- Reactjs
- Webpack


# Spinup Instructions

1. Clone the repo:
```
mkdir $GOPATH/src/github.com/nii236
cd $GOPATH/src/github.com/nii236
git@github.com:nii236/go-react-webpack.git
cd go-react-webpack
```

2. Get dependencies:
```
go get
```

3. Build the Go project:
```
go build
```

4. Build the ReactJS frontend:
```
cd js/react
npm install
TARGET=prod webpack
cd ../../
```

5. Run the API
```
./go-react-webpack
```

