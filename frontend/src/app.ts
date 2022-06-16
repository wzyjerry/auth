import type { RequestConfig } from 'umi';
import { errorHandler } from '@/util';
import moment from 'moment';
moment.locale('zh-cn');

export const request: RequestConfig = {
  prefix: 'http://oauth.windranger.tk:8888',
  timeout: 5000,
  headers: {
    'x-md-global-ip': '76.54.32.10',
  },
  errorHandler,
};
