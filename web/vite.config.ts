import { defineConfig } from 'vite'
import preact from '@preact/preset-vite'

// https://vitejs.dev/config/
export default defineConfig({
  base: './',
  plugins: [preact()],
  server: {
    // Set the port to 3000 for local development to avoid conflicts with the backend
    port: 3000,
  },
  resolve: {
    alias: {
      '@': import.meta.url + '/src',
    }
  },
  optimizeDeps: {
    include: ['preact']
  }
})
