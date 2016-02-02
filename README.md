Early boilerplate app for:

- Go
- Reactjs
- Webpack


# Spinup Instructions

Clone the repo:
```
mkdir $GOPATH/src/github.com/nii236
cd $GOPATH/src/github.com/nii236
git@github.com:nii236/go-react-webpack.git
cd go-react-webpack
```

Get dependencies:
```
go get
```

 Build the Go project:
```
go build
```

Build the ReactJS frontend:
```
cd js/react
npm install
TARGET=prod webpack
cd ../../
```

Run the API
```
./go-react-webpack
```

Visit [http://localhost:8080/ui/](http://localhost:8080/ui/).
