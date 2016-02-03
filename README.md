[](/screenshots/screenshot.jpg)

Early boilerplate app for:

- Go
- Reactjs
- Webpack


# Spinup Instructions

Clone the repo:
```
mkdir $GOPATH/src/github.com/nii236
cd $GOPATH/src/github.com/nii236
git clone git@github.com:nii236/go-react-webpack.git
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

Install JS dependencies:
```
cd js/react
npm install
npm install -g webpack
```

## For front-end development

Run the API
```
./go-react-webpack
```

Host the ReactJS front-end on `webpack-dev-server`:
```
cd js/react
npm start
```

Visit [http://localhost:3000/](http://localhost:3000/) and you can edit the JS project and watch your changes live!.

## For production

Run the API
```
./go-react-webpack
```

Build the ReactJS front-end. This will be served by the Go `jsController` service:
```
cd js/react
TARGET=prod webpack
```

Visit [http://localhost:8080/ui/](http://localhost:8080/ui/).
