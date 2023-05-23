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
  // CORS(Access-Control-Allow-Origin)エラー対策
  // フロントとバックでOrigiが異なるためエラーになる模様。リクエスト時にURLを付与しても駄目なのでconfigで指定
  devServer: {
    // proxy:'http://172.18.229.26:3000'
    proxy: 'http://touban_b_cont:3000',
  }
}