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
    port: 8081,
    // proxy: 'http://172.18.52.226:3000',     //開発時有効化。eth0のipを記述
    proxy: 'http://touban_b_cont:3000',  //コンテナでnpm run serveするならこっちを有効化
  }
}