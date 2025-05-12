import axios from 'axios';

// .env 파일에 VITE_API_BASE=http://localhost:8081
const baseURL =
  import.meta.env.VITE_API_BASE || 'http://localhost:8081';

export const api = axios.create({
  baseURL,
  headers: {
    'Content-Type': 'application/json',
  },
});

export const isAxiosError = axios.isAxiosError;
