const path = require('path')
function resolve(dir) {
  return path.join(__dirname, dir)
}

module.exports = {
  css: {
    loaderOptions: {
      less: {
        javascriptEnabled: true
      }
    }
  },
  configureWebpack: {
    resolve: {
      alias: {
        '@': resolve('src')
      }
    }
  },
  devServer: {
    proxy: {
      ['baseApi']: {
        target: `http://localhost:8081`,
        changeOrigin: true
      },
      ['openApi']: {
        target: `http://localhost:8081`,
        changeOrigin: true
      }      
    }
  }
}