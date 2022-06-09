import { RequestConfig } from 'umi';
import { errorHandler } from '@/util';

export const request: RequestConfig = {
  prefix: 'http://oauth.windranger.tk:8888',
  timeout: 5000,
  headers: {
    'x-md-global-ip': '76.54.32.10',
  },
  errorHandler
};
