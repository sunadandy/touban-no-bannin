module.exports = {
  configureWebpack: {
    module: {
      rules: [
        {
          test: /.csv$/,
          loader: 'csv-loader'
        }
      ]
    }
  },
  publicPath: './',
  // CORS(Access-Control-Allow-Origin)エラー対策。フロントとバックでOriginが異なる場合にエラー
  devServer: {
    host: "localhost",
    port: 8081, //defaultは8080
    proxy: 'http://192.168.226.46:3000',
    // proxy: 'http://touban_b_cont:3000', //コンテナでnpm run serveするならこっちを有効化
  }
}