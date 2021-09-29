import 'styles/globals.css'
import type { AppProps } from 'next/app'

declare global {
  namespace NodeJS {
    interface Global {
      proto: any
    }
  }
}

if (!process.browser) {
  global.XMLHttpRequest = require('xhr2')
  global.proto = {}
  console.log("proto")
}

function MyApp({ Component, pageProps }: AppProps) {
  console.log("aa")
  return <Component {...pageProps} />
}
export default MyApp
