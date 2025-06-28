import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

export default {
  server: {
    port: 3000,
    allowedHosts: true,
    proxy: {
      '/api/': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
      '/swagger/': {
        target: 'http://localhost:8080',
        changeOrigin: true,
      },
    },
  },
  plugins: [react(), tailwindcss()],
}
