var path = require('path');
var webpack = require('webpack');
var HtmlWebpackPlugin = require('html-webpack-plugin')
var TARGET = process.env.TARGET;

var config = {
  devtool: 'eval-source-map',
  entry: [
    'webpack-dev-server/client?http://localhost:3000',
    'webpack/hot/only-dev-server',
    './app/index'
  ],
  output: {
    path: path.join(__dirname, 'dist'),
    filename: 'bundle.js',
  },
  plugins: [
    new webpack.HotModuleReplacementPlugin(),
    new HtmlWebpackPlugin({
            filename: 'index.html',
            template: 'App/Views/template.html'
    })
  ],
  module: {
    loaders: [{
      test: /\.js$/,
      loaders: ['react-hot', 'babel'],
      include: path.join(__dirname, 'app')
    }]
  },
  resolve: {
    root: [
      path.resolve(__dirname)
    ]
  }
};

if (TARGET === 'prod') {
  var config = {
    devtool: 'source-map',
    entry: [
      './app/index'
    ],
    output: {
      path: path.join(__dirname, 'dist'),
      filename: 'bundle.js',
      publicPath: '/ui/'
    },
    plugins: [
      new webpack.NoErrorsPlugin(),
      new webpack.optimize.UglifyJsPlugin({
        compressor: { warnings: false },
        sourceMap: false
      }),
      new webpack.DefinePlugin({
        'process.env': { NODE_ENV: JSON.stringify('production') }
      }),
      new webpack.optimize.DedupePlugin(),
      new webpack.optimize.OccurenceOrderPlugin(),
      new HtmlWebpackPlugin({
              filename: 'index.html',
              template: 'App/Views/template.html'
      })
    ],
    module: {
      loaders: [{
        test: /\.js$/,
        loaders: ['react-hot', 'babel'],
        include: path.join(__dirname, 'app')
      }]
    },
    resolve: {
      root: [
        path.resolve(__dirname)
      ]
    }
  };
}

module.exports = config;
